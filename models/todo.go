package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `gorm:"size:255;not null;unique" json:"title"`
	Deskripsi string `gorm:"size:255;not null;" json:"deskripsi"`
}

func (u *Todo) SaveTodo() (*Todo, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Todo{}, err
	}
	return u, nil
}

func GetTodo() ([]Todo, error) {
	var u []Todo
	if err := DB.Find(&u).Error; err != nil {
		return u, errors.New("User not found!")
	}

	return u, nil
}

func GetTodobyId(uid uint) (Todo, error) {
	var u Todo
	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	return u, nil
}

func (u *Todo) UpdateTodobyId(uid uint) (*Todo, error) {
	if err := DB.Model(&u).Where("id = ?", uid).Updates(&u).Error; err != nil {
		return u, errors.New("User not found!")
	}

	return u, nil
}
