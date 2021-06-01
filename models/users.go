package models

import (
	"golang-example/database"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name" validate:"required,min=1,max=128"`
	Username  string     `json:"username" validate:"required,min=1,max=32"`
	Email     string     `json:"email" validate:"required,min=1,max=128,email"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func UserGet(id int) (user *User, _ error) {
	db := database.Connect()

	result := db.Where("deleted_at is null").First(&user, id)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return user, nil
}

func UserList() (users []User, _ error) {
	db := database.Connect()
	results := db.Where("deleted_at is null").Find(&users)

	if results.Error != nil {
		return users, results.Error
	}

	return users, nil
}
func UserCreate(newUser User) (User, error) {
	db := database.Connect()
	results := db.Create(&newUser)

	if results.Error != nil {
		return newUser, results.Error
	}

	return newUser, nil
}

func UserUpdate(updatedUser User) (User, error) {
	db := database.Connect()
	result := db.Model(&updatedUser).
		Updates(User{Name: updatedUser.Name, Username: updatedUser.Username, Email: updatedUser.Email})

	if result.Error != nil {
		return updatedUser, result.Error
	}

	return updatedUser, nil
}

func UserDelete(id int) error {
	user, err := UserGet(id)

	if err != nil {
		return err
	}

	db := database.Connect()

	db.Model(&user).Update("deleted_at", time.Now().UTC())

	return nil
}

func UserValidation(user User) error {
	validate := validator.New()

	err := validate.Struct(user)

	if err == nil {
		return nil
	}

	return err
}
