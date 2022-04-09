package model

import (
	"time"
	uuid "github.com/satori/go.uuid"	
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault( value: true )
}

type Base struct {
	ID 		string 		`json:"id" valid:"uuid"` 
	CreatedAt time.Time	`json:"created_at" valid:"-"`
	UpdatedAt time.Time	`json:"updated_at" valid:"-"`
}