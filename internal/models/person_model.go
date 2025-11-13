package models

import "github.com/uptrace/bun"

type Person struct {
	bun.BaseModel `bun:"table:persons"`

	ID        int    `json:"id" bun:"id,pk,autoincrement"`
	Email     string `json:"email" bun:"email,notnull"`
	Phone     string `json:"phone" bun:"phone,notnull"`
	FirstName string `json:"firstName" bun:"first_name,notnull"`
	LastName  string `json:"lastName" bun:"last_name,notnull"`
}

type CreatePersonRequest struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
