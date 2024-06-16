// main.go
package main

import (
    "splitwise/models"
    "splitwise/repository"
    "splitwise/service"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    repo := repository.NewMemoryRepository()
    expenseService := service.NewExpenseService(repo)

    users := []models.User{
        {ID: "u1", Name: "User1", Email: "user1@example.com", Mobile: "1234567890"},
        {ID: "u2", Name: "User2", Email: "user2@example.com", Mobile: "1234567891"},
        {ID: "u3", Name: "User3", Email: "user3@example.com", Mobile: "1234567892"},
        {ID: "u4", Name: "User4", Email: "user4@example.com", Mobile: "1234567893"},
    }

    for _, user := range users {
        if err := repo.AddUser(user); err != nil {
            fmt.Println(err)
        }
    }

    commands := []string{
        "SHOW",
        "SHOW u1",
        "EXPENSE u1 1000 4 u1 u2 u3 u4 EQUAL",
        "SHOW u4",
        "SHOW u1",
        "EXPENSE u1 1250 2 u2 u3 EXACT 370 880",
        "SHOW",
        "EXPENSE u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20",
        "SHOW u1",
        "SHOW",
    }

    for _, command := range commands {
        parts := strings.Fields(command)
        switch parts[0] {
        case "SHOW":
            if len(parts) == 1 {
                expenseService.ShowBalances()
            } else {
                expenseService.ShowUserBalances(parts[1])
            }
        case "EXPENSE":
            if len(parts) < 6 {
                fmt.Println("Invalid EXPENSE command format")
                continue
            }
            payerID := parts[1]
            amount, err := strconv.ParseFloat(parts[2], 64)
            if err != nil {
                fmt.Println("Invalid amount format")
                continue
            }
            numUsers, err := strconv.Atoi(parts[3])
            if err != nil {
                fmt.Println("Invalid number of users format")
                continue
            }
            userIDs := parts[4 : 4+numUsers]
            splitType := parts[4+numUsers]
            var values []float64
            if len(parts) > 5+numUsers {
                for _, v := range parts[5+numUsers:] {
                    value, err := strconv.ParseFloat(v, 64)
                    if err != nil {
                        fmt.Println("Invalid value format")
                        continue
                    }
                    values = append(values, value)
                }
            }
            err = expenseService.AddExpense(payerID, amount, userIDs, splitType, values)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}
