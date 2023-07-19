package main

import (
	"github.com/Alperen10/Image-Recognition/router"
)

func main() {
	router.CreateRouter()
	router.InitializeRoute()
	router.ServerStarter()
}
