package repository

import (
	"nusatech/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Status   string
	Token    string  `gorm:"-:migration;<-:false"`
	Balance  Balance `gorm:"foreignKey:IDUser"`
	Mailer   Mailer
}

type Balance struct {
	gorm.Model
	IDCurrency uint
	IDUser     uint
	Amount     int
}

type Currency struct {
	gorm.Model
	Currency string
}

type Mailer struct {
	gorm.Model
	Email  string
	Pin    uint
	Status string
}

func FromCore(uc users.UserCore) User {
	return User{
		Model:    gorm.Model{ID: uc.ID, CreatedAt: uc.CreatedAt, UpdatedAt: uc.UpdatedAt},
		Name:     uc.Name,
		Email:    uc.Email,
		Password: uc.Password,
		Status:   uc.Status,
		Token:    uc.Token,
		Balance:  Balance{IDCurrency: uc.Balance.IDCurrency, IDUser: uc.Balance.IDUser, Amount: uc.Balance.Amount},
		Mailer:   Mailer{Email: uc.Email, Pin: uc.Mailer.Pin, Status: uc.Status},
	}
}

func ToCore(u User) users.UserCore {
	return users.UserCore{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Status:   u.Status,
		Token:    u.Token,
		Balance:  users.BalanceCore{IDCurrency: u.Balance.IDCurrency, IDUser: u.Balance.IDUser, Amount: u.Balance.Amount},
		Mailer:   users.MailerCore{Email: u.Email, Pin: u.Mailer.Pin, Status: u.Mailer.Status},
	}
}

func ToCoreArray(au []User) []users.UserCore {
	var res []users.UserCore
	for _, val := range au {
		res = append(res, users.UserCore{
			ID:       val.ID,
			Name:     val.Name,
			Email:    val.Email,
			Password: val.Password,
			Status:   val.Status,
			Token:    val.Token,
			Balance:  users.BalanceCore{IDCurrency: val.Balance.IDCurrency, IDUser: val.Balance.IDUser, Amount: val.Balance.Amount},
			Mailer:   users.MailerCore{Email: val.Email, Pin: val.Mailer.Pin, Status: val.Mailer.Status},
		})
	}

	return res
}
