package main

import (
	"OnlineOfficeServer/Databases/Mysql"
	"OnlineOfficeServer/Router"
)

func main() {
	defer Mysql.DB.Close()
	Router.InitRouter()
}
