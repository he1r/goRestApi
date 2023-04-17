package migrate

import (
	"restApi/initializers"
	model "restApi/models"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.User{})
}
