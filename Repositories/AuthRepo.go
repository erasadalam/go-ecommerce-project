package Repositories

import (
	Cfg "GoEcommerceProject/Config"
	G "GoEcommerceProject/Globals"
	M "GoEcommerceProject/Models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)


type DB_ENV struct {
	Host, Port, Dialect, Username, Password, DBname string
}

var (
	dbEnv G.DB_ENV
)

func init() {
	godotenv.Load()
	G.DBEnv = G.DB_ENV{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Dialect:  os.Getenv("DB_DIALECT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
	}
}

func DBConnect() (*sql.DB, error) {
	dbEnv = G.DBEnv
	db, _ := sql.Open(dbEnv.Dialect, dbEnv.Username+":"+dbEnv.Password+"@tcp("+dbEnv.Host+":"+dbEnv.Port+")/"+dbEnv.DBname+"?parseTime=true")
	return db, nil
}

/*func ReadWithEmail(user M.User) (M.User, bool) {
	db, _ := DBConnect()
	var results *sql.Rows
	var err error
	results, err = db.Query("SELECT * FROM users WHERE email=?;", user.Email)

	if results.Next() {
		err = results.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.FullName, &user.Email, &user.Phone, &user.PhoneVerification, &user.Password, &user.ActiveStatus, &user.RoleID, &user.EmailVerification, &user.RememberToken)
		if err != nil {
			log.Println("AuthRepo.go Log1", err.Error())
		}
		return user, true
	} else {
		return user, false
	}

	defer db.Close()
	defer results.Close()
	return user, true
}*/

func ReadWithEmail(user M.User) (M.User, bool) {
	db := Cfg.DBConnect()
	notFound := db.First(&user, "email=?", user.Email).RecordNotFound()

	if notFound {
		defer db.Close()
		return user, false
	} else {
		defer db.Close()
		return user, true
	}
}


/*func ReadWithPhone(user M.User) (M.User, bool) {
	db, _ := DBConnect()
	var results *sql.Rows
	var err error
	results, err = db.Query("SELECT * FROM users WHERE phone=?;", user.Phone)

	if results.Next() {
		err = results.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.FullName, &user.Email, &user.Phone, &user.PhoneVerification, &user.Password, &user.ActiveStatus, &user.RoleID, &user.EmailVerification, &user.RememberToken)
		if err != nil {
			log.Println("AuthRepo.go Log2", err.Error())
		}
		return user, true
	} else {
		return user, false
	}

	defer db.Close()
	defer results.Close()
	return user, true
}*/

func ReadWithPhone(user M.User) (M.User, bool) {
	db := Cfg.DBConnect()
	notFound := db.First(&user, "phone=?", user.Phone).RecordNotFound()

	if notFound {
		defer db.Close()
		return user, false
	} else {
		defer db.Close()
		return user, true
	}
}


/*func Register(user M.User) (M.User, bool) {
	db, _ := DBConnect()
	var results *sql.Rows
	var err error
	var success bool

	_, err = db.Query("INSERT INTO users(created_at, updated_at, full_name, email, phone, password, role_id, active_status, remember_token) VALUES(now(),now(),?, ?, ?, ?, ?, ?, ?);", user.FullName, user.Email, user.Phone, user.Password, user.RoleID, user.ActiveStatus, user.RememberToken)
	if err != nil {
		log.Println("AuthRepo.go Log3", err.Error())
		return user, false
	}
	user, success = ReadWithEmail(user)
	if success {
		return user, true
	} else {
		return user, false
	}

	log.Println("AuthRepo.go Log4 Data Inserterd Successfully.\n")
	defer db.Close()
	defer results.Close()
	return user, true
}*/

func Register(user M.User) (M.User, bool) {
	db := Cfg.DBConnect()
	err := db.Create(&user).Error
	if err != nil {
		defer db.Close()
		return user, false
	}
	defer db.Close()
	return user, true
}

func Login(user M.User) (M.User, bool) {
	var success bool
	user, success = ReadWithEmail(user)
	if success {
		return user, true
	} else {
		return user, false
	}
}


func SetRememberToken(user M.User) bool {
	db, _ := DBConnect()

	results, err := db.Query("UPDATE users set remember_token=?, updated_at=now() where email=?;", user.RememberToken, user.Email)
	if err != nil {
		log.Println("AuthRepo.go Log5", err.Error())
		return false
	}

	defer db.Close()
	defer results.Close()
	return true
}


func Logout(user M.User) {

	db, _ := DBConnect()

	results, err := db.Query("UPDATE users set remember_token=NULL, updated_at=now() where email=?;", user.Email)
	if err != nil {
		log.Println("AuthRepo.go Log11", err.Error())
		return
	}

	defer db.Close()
	defer results.Close()
	return
}
