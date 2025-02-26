package main

import (
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/cgin"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
	"github.com/Cai-ki/go-caiki-blog/pkg/validate"
	"github.com/Cai-ki/go-caiki-blog/routes"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	models.SetupModels(db)
	storage.SetupStorage(db)

	v := validator.New()
	validate.SetupValidate(v)

	r := cgin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
