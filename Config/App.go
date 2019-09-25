package Config

import(
	G "GoEcommerceProject/Globals"
	"github.com/joho/godotenv"
	"os"
)

func AppConfig() {
	godotenv.Load()
	G.AppEnv.Name = os.Getenv("APP_NAME")
	G.AppEnv.Url = os.Getenv("APP_URL")
}
