package routes

type RouterGroup struct {
	TestRouter
	UserRouter
}

var RouterGroupApp = &(RouterGroup{})
