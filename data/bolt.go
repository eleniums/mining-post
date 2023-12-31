package data

import (
	"encoding/json"

	bolt "go.etcd.io/bbolt"
)

const (
	playersBucketName = "players"
)

type BoltDB struct {
	DB *bolt.DB
}

func NewBoltDB() *BoltDB {
	return &BoltDB{}
}

func (b *BoltDB) Open(conn string) error {
	// open connection to database
	db, err := bolt.Open(conn, 0600, nil)
	if err != nil {
		return err
	}

	// make sure database is set up and ready to be used
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(playersBucketName))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	b.DB = db

	return nil
}

func (b *BoltDB) Close() error {
	return b.DB.Close()
}

func (b *BoltDB) LoadPlayers() ([]Player, error) {
	players := []Player{}
	err := b.DB.View(func(tx *bolt.Tx) error {
		playersBucket := tx.Bucket([]byte(playersBucketName))

		playersBucket.ForEach(func(k, v []byte) error {
			var player Player
			err := json.Unmarshal(v, &player)
			if err != nil {
				return err
			}

			players = append(players, player)

			return nil
		})

		return nil
	})
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (b *BoltDB) SavePlayers(players []Player) error {
	err := b.DB.Update(func(tx *bolt.Tx) error {
		playersBucket := tx.Bucket([]byte(playersBucketName))

		for _, player := range players {
			serialized, err := json.Marshal(player)
			if err != nil {
				return err
			}

			err = playersBucket.Put([]byte(player.Name), serialized)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
