package service

import (
    "fmt"
    "snake-and-ladder/internal/models"
    "snake-and-ladder/internal/repository"
    "snake-and-ladder/internal/utils"
)

type GameService struct {
    game *models.Game
    dice *utils.Dice
}

func NewGameService(repo repository.GameRepository) (*GameService, error) {
    game, err := repo.LoadGame()
    if err != nil {
        return nil, err
    }

    return &GameService{
        game: game,
        dice: utils.NewDice(6),
    }, nil
}

func (s *GameService) Play() {
    for {
        for _, player := range s.game.Players {
            initialPosition := player.Position
            diceRoll := s.dice.Roll()
            newPosition := initialPosition + diceRoll

            if newPosition > s.game.BoardSize {
                newPosition = initialPosition
            } else {
                newPosition = s.checkPosition(newPosition)
            }

            player.Position = newPosition
            fmt.Printf("%s rolled a %d and moved from %d to %d\n", player.Name, diceRoll, initialPosition, newPosition)

            if newPosition == s.game.BoardSize {
                fmt.Printf("%s wins the game\n", player.Name)
                return
            }
        }
    }
}

func (s *GameService) checkPosition(position int) int {
    for {
        if tail, ok := s.game.Snakes[position]; ok {
            position = tail
        } else if end, ok := s.game.Ladders[position]; ok {
            position = end
        } else {
            break
        }
    }
    return position
}
