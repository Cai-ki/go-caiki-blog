package main

import (
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
	"github.com/Cai-ki/go-caiki-blog/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	storage.SetupStorage(db)

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
