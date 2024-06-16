package repository

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "snake-and-ladder/internal/models"
)

type GameRepository interface {
    LoadGame() (*models.Game, error)
}

type CommandLineGameRepository struct{}

func NewCommandLineGameRepository() *CommandLineGameRepository {
    return &CommandLineGameRepository{}
}

func (repo *CommandLineGameRepository) LoadGame() (*models.Game, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter number of snakes: ")
    sStr, _ := reader.ReadString('\n')
    s, _ := strconv.Atoi(strings.TrimSpace(sStr))
    snakes := make([]models.Position, s)
    for i := 0; i < s; i++ {
        fmt.Printf("Enter head and tail of snake %d: ", i+1)
        snakeStr, _ := reader.ReadString('\n')
        positions := strings.Split(strings.TrimSpace(snakeStr), " ")
        head, _ := strconv.Atoi(positions[0])
        tail, _ := strconv.Atoi(positions[1])
        snakes[i] = models.Position{Start: head, End: tail}
    }

    fmt.Print("Enter number of ladders: ")
    lStr, _ := reader.ReadString('\n')
    l, _ := strconv.Atoi(strings.TrimSpace(lStr))
    ladders := make([]models.Position, l)
    for i := 0; i < l; i++ {
        fmt.Printf("Enter start and end of ladder %d: ", i+1)
        ladderStr, _ := reader.ReadString('\n')
        positions := strings.Split(strings.TrimSpace(ladderStr), " ")
        start, _ := strconv.Atoi(positions[0])
        end, _ := strconv.Atoi(positions[1])
        ladders[i] = models.Position{Start: start, End: end}
    }

    fmt.Print("Enter number of players: ")
    pStr, _ := reader.ReadString('\n')
    p, _ := strconv.Atoi(strings.TrimSpace(pStr))
    players := make([]*models.Player, p)
    for i := 0; i < p; i++ {
        fmt.Printf("Enter name of player %d: ", i+1)
        name, _ := reader.ReadString('\n')
        players[i] = &models.Player{Name: strings.TrimSpace(name), Position: 0}
    }

    snakeMap := make(map[int]int)
    ladderMap := make(map[int]int)
    for _, snake := range snakes {
        snakeMap[snake.Start] = snake.End
    }
    for _, ladder := range ladders {
        ladderMap[ladder.Start] = ladder.End
    }

    return &models.Game{
        BoardSize: 100,
        Snakes:    snakeMap,
        Ladders:   ladderMap,
        Players:   players,
    }, nil
}
