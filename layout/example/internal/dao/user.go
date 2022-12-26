package dao

import (
	"github.com/cro4k/gms/layout/example/internal/model"
	"github.com/cro4k/gms/layout/public/db"
)

var User = new(userService)

type userService struct{}

func (s *userService) Create(u *model.User) error {
	return db.DB().Create(u).Error
}

func (s *userService) Update(id int, name string) error {
	return db.DB().Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
}

func (s *userService) Find(id int) (*model.User, error) {
	var u = new(model.User)
	err := db.DB().First(u, "id = ?", id).Error
	return u, err
}

func (s *userService) Delete(id int) error {
	return db.DB().Delete(&model.User{}, "id = ?", id).Error
}
