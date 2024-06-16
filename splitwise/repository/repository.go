// repository/repository.go
package repository

import "splitwise/models"

type Repository interface {
    AddUser(user models.User) error
    GetUser(id string) (models.User, error)
    AddExpense(balances []models.Balance) error
    GetBalances() []models.Balance
    GetUserBalances(userID string) []models.Balance
}

