package main

import (
	"BlogDemo/routes"
)

func main() {
	r := routes.NewRouter()
	r.Run(":5555")
}
