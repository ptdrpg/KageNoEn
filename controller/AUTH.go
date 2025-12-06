package controller

import (
	"KageNoEn/lib"
	"KageNoEn/model"
	"encoding/json"
	"net/http"
	"time"
)

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

	pssw, err := lib.HashPass(input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Id:        res.Id,
		Username:  input.Username,
		Password:  pssw,
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

	token, err := lib.GenerateToken(res.Id, role.Label, input.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &model.RegisterResponse{
		Data: *user,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}


