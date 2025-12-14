package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID        string         `json:"id" gorm:"type:char(36);not null;primaryKey;uniqueIndex"`
	Name      string         `json:"name" gorm:"type:char(50);not null"`
	StartDate time.Time      `json:"start_date"`
	EndDate   time.Time      `json:"end_date"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}

func (u *Course) BeforeCreate(tx *gorm.DB) (err error) {

	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
