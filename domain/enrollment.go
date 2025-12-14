package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string         `json:"id" gorm:"type:char(36);not null; primaryKey;unique_index"`
	UserID    string         `json:"user_id,omitempty" gorm:"type:char(36)"`
	User      *User          `json:"user,omitempty"`
	CourseID  string         `json:"course_id" gorm:"type:char(36)"`
	Course    *Course        `json:"course,omitempty"`
	Status    string         `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}

func (e *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {

	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return
}
