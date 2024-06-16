package models

type Position struct {
    Start int
    End   int
}

type Player struct {
    Name     string
    Position int
}

type Game struct {
    BoardSize int
    Snakes    map[int]int
    Ladders   map[int]int
    Players   []*Player
}
