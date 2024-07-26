package main

import "github.com/Gntrstwn19/echo-framework-register-network.git/route"

func main() {
	e := route.InitRoute()

	e.Logger.Fatal(e.Start(":1234"))
}
