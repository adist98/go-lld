package main

import (
    "fmt"
    "snake-and-ladder/internal/repository"
    "snake-and-ladder/internal/service"
)

func main() {
    repo := repository.NewCommandLineGameRepository()
    gameService, err := service.NewGameService(repo)
    if err != nil {
        fmt.Printf("Error initializing game: %v\n", err)
        return
    }

    gameService.Play()
}
