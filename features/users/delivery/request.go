package delivery

import "nusatech/features/users"

type RegisterFormat struct {
	Name     string `json:"name" form:"name" validate:"required,min=4,max=30"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,alpha,numeric"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateFormat struct {
	Name     string `json:"name" form:"name" validate:"min=4,max=30"`
	OldEmail string `json:"old_email" form:"old_email" validate:"email"`
	NewEmail string `json:"new_email" form:"new_email" validate:"email"`
	Password string `json:"password" form:"password" validate:"alpha,numeric"`
}

func ToCore(i interface{}) users.UserCore {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return users.UserCore{Name: cnv.Name, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return users.UserCore{Email: cnv.Email, Password: cnv.Password}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return users.UserCore{Name: cnv.Name, Email: cnv.NewEmail, Password: cnv.Password}
	}

	return users.UserCore{}
}
