package config

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

func NewDsn(dialect string, user string, password string, host string, port string, db string) (dsn string) {
	encodedPassword := url.QueryEscape(password)
	return fmt.Sprintf(`%s://%s:%s@%s:%s/%s`, dialect, user, encodedPassword, host, port, db)
}

// Initialize DB connection and returns the connection
func InitDB(env *Env) *gorm.DB {
	// create dsn
	dsn := NewDsn("postgres", env.DBUser, env.DBPassword, env.DBHostName, env.DBPort, env.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: NewGormLogger(gormlog.Warn),
	})

	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
