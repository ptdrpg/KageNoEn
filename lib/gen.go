package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type IdModel struct {
	Id string
}

func GenerateId(username string) (IdModel, error) {
	data := fmt.Sprintf(username, time.Now().String())

	hash := sha256.Sum256([]byte(data))
	res := IdModel{
		Id: hex.EncodeToString(hash[:]),
	}

	return res, nil
}
