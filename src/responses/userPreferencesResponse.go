package responses

import "BudgBackend/src/models"

// UserPreferencesResponse is a response.
type UserPreferencesResponse struct {
	User     models.User `json:"user"`
	BudgetId int         `json:"budgetId"`
}
