package model

import (
  "github.com/asaskevich/govalidator"
  uuid "github.com/satori/go.uuid"
  "time"
)

// serilizar para json `json:""`
type Bank struct {
  Base `valid:"required"`
  Code string `json:"code" valid:"notnull"`
  Name string `json:"name" valid:"notnull"`
  Accounts []*Account `valid:"-"`
}

// método attached para a struct
func (bank *Bank) isValid() error {
  validator, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}


// NewBank  sempre vai retornar 2 valores bank e error
func NewBank(code string, name string) (*Bank, error) {
  bank := Bank{
    Code: code,
    Name: name,
  }

  bank.ID = uuid.NewV4().String()
  bank.CreatedAt = time.Now()

  err := bank.isValid()
  if err != nil {
    return nil, err
  }

  return &bank, nil
}