package service

type Service struct {
	CasbinService
	UserService
	ApiService
	RoleService
}

var ServiceApp = new(Service)
