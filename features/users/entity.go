package users

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Status    string
	Token     string
	Balance   BalanceCore
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BalanceCore struct {
	ID         uint
	IDCurrency uint
	IDUser     uint
	Amount     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CurrencyCore struct {
	ID       uint
	Currency string
}

type MailerCore struct {
	ID     uint
	Email  string
	Pin    uint
	Status string
}

type Repository interface {
	Insert(data UserCore) (UserCore, error)
	GetLogin(input UserCore) (UserCore, error)
	GetByEmail(email string) (UserCore, error)
	Edit(data UserCore, id uint) (UserCore, error)
}

type Service interface {
	Create(data UserCore) (UserCore, error)
	Login(input UserCore) (UserCore, error)
	ShowByEmail(email string) (UserCore, error)
	Update(data UserCore, id uint) (UserCore, error)
}
