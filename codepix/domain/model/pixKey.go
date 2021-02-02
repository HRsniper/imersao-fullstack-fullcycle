 package model

import (
  "errors"
  "github.com/asaskevich/govalidator"
  uuid "github.com/satori/go.uuid"
  "time"
)

// interface para implementação do db
type PixKeyRepositoryInterface interface {
  // função() (retorno)
  RegisterKey(pixKey *PixKey) (*PixKey, error)
  FindKeyByKind(key string, kind string) (*PixKey, error)
  AddBank(bank *Bank) error
  AddAccount(account *Account) error
  FindAccount(id string) (*Account, error)
}

type PixKey struct {
  Base `valid:"required"`
  Kind string `json:"kind" valid:"notnull"`
  Key string `json:"key" valid:"notnull"`
  AccountID string `gorm:"column:account_id;type:uuid;not null" valid:"-"`
  Account *Account `valid:"-"`
  Status string `json:"status" valid:"notnull"`
}

// método attached para a struct
func (pixKey *PixKey) isValid() error {
  _, err := govalidator.ValidateStruct(pixKey)

  if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
    return errors.New("invalid type of key")
  }

  if pixKey.Status != "active" && pixKey.Status != "inactive" {
    return errors.New("invalid status")
  }

  if err != nil {
    return err
  }

  return nil
}

// NewPixKey  sempre vai retornar 2 valores account e error
func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {

  pixKey := PixKey{
    Kind: kind,
    Key: key,
    Account: account,
    Status: "active",
  }

  pixKey.ID = uuid.NewV4().String()
  pixKey.CreatedAt = time.Now()

  err := pixKey.isValid()
  if err != nil {
    return nil, err
  }

  return &pixKey, nil
}