package responses

import "BudgBackend/src/models"

// UserPreferencesResponse is a response.
type UserPreferencesResponse struct {
	ExpenseCategories []models.Category `json:"expenseCategories"`
	User              models.User       `json:"user"`
}
