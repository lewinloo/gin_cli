package routes

type RouterGroup struct {
	TestRouter
	UserRouter
	RoleRouter
	ApiRouter
	PublicRouter
}

var RouterGroupApp = &(RouterGroup{})
