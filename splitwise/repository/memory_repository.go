// repository/memory_repository.go
package repository

import (
    "errors"
    "splitwise/models"
)

type MemoryRepository struct {
    users    map[string]models.User
    balances []models.Balance
}

func NewMemoryRepository() *MemoryRepository {
    return &MemoryRepository{
        users:    make(map[string]models.User),
        balances: []models.Balance{},
    }
}

func (r *MemoryRepository) AddUser(user models.User) error {
    if _, exists := r.users[user.ID]; exists {
        return errors.New("user already exists")
    }
    r.users[user.ID] = user
    return nil
}

func (r *MemoryRepository) GetUser(id string) (models.User, error) {
    user, exists := r.users[id]
    if !exists {
        return models.User{}, errors.New("user not found")
    }
    return user, nil
}

func (r *MemoryRepository) AddExpense(balances []models.Balance) error {
    r.balances = append(r.balances, balances...)
    return nil
}

func (r *MemoryRepository) GetBalances() []models.Balance {
    return r.balances
}

func (r *MemoryRepository) GetUserBalances(userID string) []models.Balance {
    var result []models.Balance
    for _, balance := range r.balances {
        if balance.From == userID || balance.To == userID {
            result = append(result, balance)
        }
    }
    return result
}
