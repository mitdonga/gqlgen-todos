package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/mitdonga/gqlgen-todos/graph/model"
	"gorm.io/gorm"
)

type Resolver struct {
	todos []*model.Todo
	DB    *gorm.DB
}


