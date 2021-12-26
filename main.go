package main

import (
	"OnlineOfficeServer/Databases/Mysql"
	"OnlineOfficeServer/Router"
)

func main() {
	defer Mysql.SqlDB.Close()
	Router.InitRouter()
}
