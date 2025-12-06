package controller

import (
	"KageNoEn/lib"
	"KageNoEn/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func (c *Controller) GetAllUser(w http.ResponseWriter, r *http.Request) {
	allUsers, err := c.R.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &model.UserList{
		Data: allUsers,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := c.R.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &model.UserResponse{
		Data: user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var input model.UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, genErr := lib.GenerateId(input.Username)
	if genErr != nil {
		http.Error(w, genErr.Error(), http.StatusInternalServerError)
		return
	}

	eloRank, _ := c.R.GetbyElo(0)
	role, _ := c.R.GetRoleByLabel("player")
	status, _ := c.R.GetUserStatusByLabel("active")

	user := &model.User{
		Id:        res.Id,
		Username:  input.Username,
		Password:  input.Password,
		Email:     input.Email,
		RankId:    eloRank.Id,
		RoleId:    role.Id,
		StatusId:  status.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.R.CreateUser(*user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := &model.UserResponse{
		Data: *user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
