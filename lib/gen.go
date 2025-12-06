package lib

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IdModel struct {
	Id string
}

func GenerateId(username string) (IdModel, error) {
	data := fmt.Sprintf(username, time.Now().String())
	
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return IdModel{}, err
	}

	res := IdModel{
		Id: string(hash),
	}

	return res, nil
}
