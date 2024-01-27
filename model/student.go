package model

import (
	"go-prompt/model/enum"

	"github.com/google/uuid"
)

type Student struct {
	Id                  uuid.UUID
	FirstName           string
	LastName            string
	Gender              enum.Gender
	Email               string
	TumId               string
	MatriculationNumber string
}
