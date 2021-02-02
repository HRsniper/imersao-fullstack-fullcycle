package model

import (
  "errors"
  "github.com/asaskevich/govalidator"
  uuid "github.com/satori/go.uuid"
  "time"
)

// serilizar para json `json:""`
type User struct {
  Base `valid:"required"`
  Name string `json:"name" valid:"required"`
  Email string `json:"email" valid:"required"`
}

// método attached para a struct
func (user *User) isValid() error {
  _, err := govalidator.ValidateStruct(user)

  if user.ID == "" || user.Name == "" || user.Email == ""{
    return errors.New("these fields are mandatory")
  }

	if err != nil {
		return err
	}

	return nil
}


// NewUser  sempre vai retornar 2 valores User e error
func NewUser(name string, email string) (*User, error) {
  user := User{
    Name: name,
    Email: email,
  }

  user.ID = uuid.NewV4().String()
  user.CreatedAt = time.Now()

  err := user.isValid()
  if err != nil {
    return nil, err
  }

  return &user, nil
}