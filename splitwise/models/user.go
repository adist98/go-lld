// models/user.go
package models

type User struct {
    ID     string
    Name   string
    Email  string
    Mobile string
}

type Balance struct {
    From   string
    To     string
    Amount float64
}
