package data

import (
	bolt "go.etcd.io/bbolt"
)

type BoltDB struct {
	DB *bolt.DB
}

func NewBoltDB() *BoltDB {
	return &BoltDB{}
}

func (b *BoltDB) Open(conn string) error {
	db, err := bolt.Open(conn, 0600, nil)
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
	// TODO
	return nil, nil
}

func (b *BoltDB) SavePlayer(player Player) error {
	// TODO
	return nil
}
