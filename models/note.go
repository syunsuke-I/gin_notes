package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NotesAll() *[]Note {
	var notes []Note
	DB.Where("deleted at is NULL").Order("updated at desc").Find(&notes)
	return &notes
}
