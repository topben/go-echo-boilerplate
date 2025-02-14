package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/topben/go-echo-boilerplate/common/models"
)

type Blog struct {
	models.Base
	Title   string
	Content string
	UserID  uuid.UUID
}
