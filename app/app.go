package app

import (
	"fmt"
	"gin_cli/initialize"
)

func Run(port string) {
	port = fmt.Sprintf(":%s", port)
	router := initialize.Routes()

	if err := router.Run(port); err != nil {
		panic(err)
	}
	fmt.Printf(`Swagger文档地址:http://127.0.0.1%s/swagger/index.html`, port)
}
