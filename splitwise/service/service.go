// service/service.go
package service

import (
    "splitwise/models"
    "splitwise/repository"
    "fmt"
    "math"
)

type ExpenseService struct {
    repo repository.Repository
}

func NewExpenseService(repo repository.Repository) *ExpenseService {
    return &ExpenseService{repo: repo}
}

func (s *ExpenseService) AddExpense(payerID string, amount float64, userIDs []string, splitType string, values []float64) error {
    var balances []models.Balance

    switch splitType {
    case "EQUAL":
        share := math.Round((amount/float64(len(userIDs)))*100) / 100
        for _, userID := range userIDs {
            if userID != payerID {
                balances = append(balances, models.Balance{From: userID, To: payerID, Amount: share})
            }
        }
    case "EXACT":
        if len(values) != len(userIDs) {
            return fmt.Errorf("values count does not match users count")
        }
        total := 0.0
        for _, value := range values {
            total += value
        }
        if total != amount {
            return fmt.Errorf("exact values do not sum up to total amount")
        }
        for i, userID := range userIDs {
            if userID != payerID {
                balances = append(balances, models.Balance{From: userID, To: payerID, Amount: values[i]})
            }
        }
    case "PERCENT":
        if len(values) != len(userIDs) {
            return fmt.Errorf("values count does not match users count")
        }
        total := 0.0
        for _, value := range values {
            total += value
        }
        if total != 100 {
            return fmt.Errorf("percent values do not sum up to 100")
        }
        for i, userID := range userIDs {
            if userID != payerID {
                amountOwed := math.Round((amount*values[i]/100)*100) / 100
                balances = append(balances, models.Balance{From: userID, To: payerID, Amount: amountOwed})
            }
        }
    default:
        return fmt.Errorf("invalid split type")
    }

    return s.repo.AddExpense(balances)
}

func (s *ExpenseService) ShowBalances() {
    balances := s.repo.GetBalances()
    if len(balances) == 0 {
        fmt.Println("No balances")
        return
    }
    for _, balance := range balances {
        fmt.Printf("%s owes %s: %.2f\n", balance.From, balance.To, balance.Amount)
    }
}

func (s *ExpenseService) ShowUserBalances(userID string) {
    balances := s.repo.GetUserBalances(userID)
    if len(balances) == 0 {
        fmt.Println("No balances")
        return
    }
    for _, balance := range balances {
        fmt.Printf("%s owes %s: %.2f\n", balance.From, balance.To, balance.Amount)
    }
}
