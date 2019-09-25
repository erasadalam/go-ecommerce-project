package Seeders

import (
	Cfg "GoEcommerceProject/Config"
)

func Seed() {
	db := Cfg.DBConnect()
	RoleSeeder()
	UserSeeder()
	PayMethodSeeder()
	defer db.Close()
}

