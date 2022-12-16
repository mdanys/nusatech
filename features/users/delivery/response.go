package delivery

import "nusatech/features/users"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type UserResponse struct {
	ID     uint   `json:"id_user" form:"id_user"`
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Status string `json:"status" form:"status"`
}

type LoginResponse struct {
	ID     uint   `json:"id_user" form:"id_user"`
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Status string `json:"status" form:"status"`
	Token  string `json:"token" form:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "user":
		cnv := core.(users.UserCore)
		res = UserResponse{ID: cnv.ID, Name: cnv.Name, Email: cnv.Email, Status: cnv.Status}
	case "login":
		cnv := core.(users.UserCore)
		res = LoginResponse{ID: cnv.ID, Name: cnv.Name, Email: cnv.Email, Status: cnv.Status, Token: cnv.Token}
	case "getall":
		var arr []UserResponse
		cnv := core.([]users.UserCore)
		for _, val := range cnv {
			arr = append(arr, UserResponse{ID: val.ID, Name: val.Name, Email: val.Email, Status: val.Status})
		}
		res = arr
	}

	return res
}
