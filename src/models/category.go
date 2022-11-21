package models

import (
	"BudgBackend/src/database"
	"fmt"
)

type ICategory interface {
	GetCategories(userId int) []Transaction
	GetCategory(id string) Transaction
	CreateCategory(userId int, name string) Transaction
	DeleteCategory(id int) Transaction
}

type Category struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
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
		if category.UserId == 1 {
			category.Default = true
		} else {
			category.Default = false
		}
		categories = append(categories, category)
	}

	if i == 0 {
		return categories, fmt.Errorf("No categories for user with id: %d", userId)
	}
	return categories, nil
}

func (c *Category) CreateCategory() (int64, error) {
	InfoLogger.Println("Creating category")
	// check if category with same name exists
	query := fmt.Sprintf("INSERT INTO Category (user_id, name) VALUES (%d, '%s')", c.UserId, c.Name)
	id, err := database.InsertDB(query)
	if err != nil {
		ErrorLogger.Println("Error creating category: ", err)
		return id, err
	}
	return id, nil
}
