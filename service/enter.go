package service

type Service struct {
	CasbinService
	UserService
}

var ServiceApp = new(Service)
