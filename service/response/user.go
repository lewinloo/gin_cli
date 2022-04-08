package response

import "gin_cli/model"

type LoginResponse struct {
	UserInfo model.User `json:"userInfo"`
	Token    string     `json:"token"`
}
