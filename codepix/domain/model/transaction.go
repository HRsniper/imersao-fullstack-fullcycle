package model

import (
  "errors"
  "github.com/asaskevich/govalidator"
  uuid "github.com/satori/go.uuid"
  "time"
)

// estado da transação
const (
  TransactionPending string = "pending"
  TransactionCompleted string = "completed"
  TransactionError string = "error"
  TransactionConfirmed string = "confirmed"
)

// interface para implementação do db
type TransactionRepositoryInterface interface {
  // função() (retorno)
  Register(transaction *Transaction) error
  Save(transaction *Transaction) error
  Find(id string) (*Transaction, error)
}

// lista de transação
type Transactions struct {
  Transaction []Transaction
}

type Transaction struct {
  Base `valid:"required"`
  AccountFrom *Account `valid:"-"`
  Amount float64 `json:"amount" valid:"notnull"`
  PixKeyTo *PixKey `valid:"-"`
  Status string `json:"status" valid:"notnull"`
  Description string `json:"description" valid:"notnull"`
  CancelDescription string `json:"cancel_description" valid:"-"`
}

// método attached para a struct
func (t *Transaction) isValid() error {
  validator, err := govalidator.ValidateStruct(t)

  if t.Amount <= 0 {
    return errors.New("the amount must be greater than 0")
  }

  if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
    return errors.New("invalid status for the transaction")
  }

  if t.PixKeyTo.AccountID == t.AccountFrom.ID {
    return errors.New("the source and destination account cannot be the same")
  }

  if err != nil {
    return err
  }
  
  return nil
}

// NewTransaction  sempre vai retornar 2 valores account e error
func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
  transaction := Transaction{
    AccountFrom: accountFrom,
    Amount: amount,
    PixKeyTo: pixKeyTo,
    Status: TransactionPending,
    Description: description,
  }

  transaction.ID = uuid.NewV4().String()
  transaction.CreatedAt = time.Now()

  err := transaction.isValid()
  if err != nil {
    return nil, err
  }

  return &transaction, nil
}

// método attached para a struct
func (t *Transaction) Complete() error {
  t.Status = TransactionCompleted
  t.UpdatedAt = time.Now()
  err := t.isValid()
  return err
}

// método attached para a struct
func (t *Transaction) Confirm() error {
  t.Status = TransactionConfirmed
  t.UpdatedAt = time.Now()
  err := t.isValid()
  return err
}

// método attached para a struct
func (t *Transaction) Cancel(description string) error {
  t.Status = TransactionError
  t.UpdatedAt = time.Now()
  t.Description = description
  err := t.isValid()
  return err
}