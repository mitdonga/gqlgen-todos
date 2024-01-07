package db

import (
    "log"

    "github.com/mitdonga/gqlgen-todos/graph/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:password@localhost:5432/gqlgen_todos"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
			log.Fatalln(err)
	}

	db.AutoMigrate(&model.Todo{})

	return db
}
