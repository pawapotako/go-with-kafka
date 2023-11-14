package main

import (
	"database/sql"
	"fmt"

	"log"
	"strconv"
	"time"

	"go-with-kafka/internal/handler"
	"go-with-kafka/internal/model"
	"go-with-kafka/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	initTimeZone()
	config := util.LoadConfig()
	db := initDatabase(config)
	initConsumer(config)
	initRouter(db, config)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal("cannot set timezone", err)
	}

	time.Local = ict
}

func initDatabase(config util.Config) *gorm.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Asia%%2FBangkok",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("cannot connect to db ", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot open db ", err)
	}

	migration(db)

	return db
}

func migration(db *gorm.DB) {

	if err := db.AutoMigrate(&model.UserModel{}); err != nil {
		log.Fatal("cannot auto migrate db", err)
	}
}

func initRouter(db *gorm.DB, config util.Config) {

	gin := gin.Default()
	validator := validator.New()

	handler.InitDefaultHandler(gin)
	handler.InitUserHandler(db, gin, validator)

	gin.Run(":" + strconv.Itoa(config.App.Port))
}

func initConsumer(config util.Config) {

	conn := util.KafkaConn(config, "user")

	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(message.Value))
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
