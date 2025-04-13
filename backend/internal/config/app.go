package config

import "gorm.io/gorm"

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = InitDB(app.Env)
	return *app
}
