package main

import (
	"restful-cake-store/config"
	"restful-cake-store/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)
func main() {

	db := config.ConnectDB()
	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "db"); err != nil {
		panic(err)
	}
	router := gin.Default()

	router.GET("/cakes/:id", controller.DetailOfCake)
	router.GET("/cakes", controller.ListOfCakes)
	router.POST("/cakes", controller.AddCake)
	router.PATCH("/cakes/:id", controller.UpdateCake)
	router.DELETE("/cakes/:id", controller.DeleteCake)

	router.Run(":3000")
}

