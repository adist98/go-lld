package utils

import (
    "math/rand"
    "time"
)

type Dice struct {
    Sides int
}

func NewDice(sides int) *Dice {
    return &Dice{Sides: sides}
}

func (d *Dice) Roll() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(d.Sides) + 1
}
