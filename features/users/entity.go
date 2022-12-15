package users

type UserCore struct {
	ID       uint
	Email    string
	Password string
	Status   string
	Token    string
}

type BalanceCore struct {
	ID         uint
	IDCurrency uint
	IDUser     uint
	Amount     int
}

type CurrencyCore struct {
	ID uint
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
