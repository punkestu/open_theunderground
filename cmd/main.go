package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	"github.com/punkestu/open_theunderground/internal/middleware/repo/jwt"
	api2 "github.com/punkestu/open_theunderground/internal/post/handler/api"
	db2 "github.com/punkestu/open_theunderground/internal/post/repo/db"
	"github.com/punkestu/open_theunderground/internal/user/handler/api"
	"github.com/punkestu/open_theunderground/internal/user/repo/db"
)

func main() {
	app := fiber.New()
	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/theunderground?parseTime=true")
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			print(err.Error())
		}
	}(conn)
	if err != nil {
		println(err.Error())
		return
	}

	userRepo := db.NewUserDB(conn)
	postRepo := db2.NewPostDB(conn)
	mJwtValidator := jwtValidator.NewValidator(userRepo)
	midAuth := auth.CreateMiddleware(mJwtValidator)
	api.InitUser(app, userRepo, midAuth)
	api2.InitPost(app, postRepo, midAuth)

	err = app.Listen(":8080")
	if err != nil {
		println(err.Error())
		return
	}
}
