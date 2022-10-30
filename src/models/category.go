package models

import (
	"BudgBackend/src/database"
	"fmt"
)

type ICategory interface {
	GetCategories(userId int) []Expense
	GetCategory(id string) Expense
	CreateCategory(userId int, name string) Expense
	DeleteCategory(id int) Expense
}

type Category struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Name   string `json:"name"`
}

var Categories []Category

func (c *Category) GetCategories(userId int) ([]Category, error) {
	var categories []Category
	InfoLogger.Println("Getting categories")
	query := fmt.Sprintf("SELECT id as id, user_id as userId, name as name FROM Category WHERE user_id = %d or user_id = 1", userId)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting categories: ", err)
		return categories, err
	}

	i := 0
	for rows.Next() {
		i++
		var category Category
		err = rows.Scan(&category.Id, &category.UserId, &category.Name)
		if err != nil {
			ErrorLogger.Println("Error scanning categories: ", err)
		}
		categories = append(categories, category)
	}

	if i == 0 {
		return categories, fmt.Errorf("No categories for user with id: %d", userId)
	}
	return categories, nil
}
