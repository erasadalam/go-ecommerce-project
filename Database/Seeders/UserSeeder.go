package Seeders

import (
	Cfg "GoEcommerceProject/Config"
	Mod "GoEcommerceProject/Models"
	"golang.org/x/crypto/bcrypt"
)

var users []Mod.User

func UserSeeder() {
	user1()
	db := Cfg.DBConnect()
	for i, _ := range users {
		db.Where(&Mod.User{Email:users[i].Email}).FirstOrCreate(&users[i])
	}

}

func user1() {
	hash, _:= bcrypt.GenerateFromPassword([]byte("*123456#"), 10)
	user := Mod.User{
		FullName: "Mr. Admin",
		Email: "admin@xyz.com",
		Password: string(hash),
		ActiveStatus: 1,
		RoleID: 1,
	}
	users = append(users, user)
}

