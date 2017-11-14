package main

import (
	"os"
	_ "github.com/heroku/testapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	port := os.Getenv("PORT")
    beego.Run(":"+port)
}
