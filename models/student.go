package models

import (
	"time"

	"gorm.io/gorm"
)

// Student represents the student table in the database.
type Student struct {
	// gorm.Model includes fields like ID, CreatedAt, UpdatedAt, DeletedAt.
	// ID will be SERIAL PRIMARY KEY by default if not specified otherwise.
	gorm.Model

	// Name VARCHAR(100) NOT NULL
	Name string `json:"name" binding:"required" gorm:"type:varchar(100);not null"`

	// Email VARCHAR(100) UNIQUE NOT NULL
	Email string `json:"email" binding:"required,email" gorm:"type:varchar(100);unique;not null"`

	// Age INT CHECK (age >= 16)
	// 'gte=16' is a validator tag for "greater than or equal to 16".
	Age int `json:"age" binding:"required,gte=16" gorm:"type:integer;check:age >= 16"`

	// Department VARCHAR(50)
	Department string `json:"department" gorm:"type:varchar(50)"`

	// EnrolledAt DATE DEFAULT CURRENT_DATE
	// GORM will handle DEFAULT CURRENT_DATE automatically for time.Time if not set.
	EnrolledAt time.Time `json:"enrolled_at" gorm:"type:date"`
}

func (Student) TableName() string {
	return "student"
}
