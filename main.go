package main

import (
	Cfg "GoEcommerceProject/Config"
	Mig "GoEcommerceProject/Database/Migrations"
	Seed "GoEcommerceProject/Database/Seeders"
	"GoEcommerceProject/Routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func main() {
	err := os.MkdirAll("./Storage/Images", 0777)
	if err!=nil {
		log.Println(err.Error())
	}
	Cfg.AppConfig()
	Mig.Migrate()
	Seed.Seed()
	Routes.Routes()
}
