package game

import (
	"github.com/eleniums/mining-post/data"
)

type Storage interface {
	Open(conn string) error
	Close() error
	LoadPlayers() ([]data.Player, error)
	SavePlayers(players []data.Player) error
}
