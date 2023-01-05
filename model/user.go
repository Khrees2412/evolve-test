package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Base
	Name        string `json:"name"`
	Email       string `json:"email"`
	Bio         string `json:"bio"`
	AvailableOn string `json:"available_on"`
}

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func Paginate(p Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (p.Page - 1) * p.Size
		return db.Offset(offset).Limit(p.Size)
	}
}

// GenerateISOString generates a time string equivalent to Date.now().toISOString in JavaScript
func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// Base contains common columns for all tables
type Base struct {
	ID        uint   `gorm:"primaryKey" autoIncrement:"true" json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// generate timestamps
	t := GenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}
