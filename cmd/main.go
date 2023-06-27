package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	"github.com/punkestu/open_theunderground/internal/user/handler/api"
	"github.com/punkestu/open_theunderground/internal/user/repo/db"
)

func main() {
	app := fiber.New()
	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/theunderground")
	defer conn.Close()
	if err != nil {
		println(err.Error())
		return
	}
	userRepo := db.NewUserDB(conn)
	midUser := auth.CreateMiddleware(userRepo)
	api.InitUser(app, userRepo, midUser)
	app.Listen(":8080")
}
