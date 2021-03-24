package repository

import (
	badger "github.com/dgraph-io/badger/v3"
)

//ListNodeBadgerDB
type ListNodeBadgerDB struct {
	db *badger.DB
}

//NewListNodeBadgerDB create new repository
func NewListNodeBadgerDB(db *badger.DB) *ListNodeBadgerDB {
	return &ListNodeBadgerDB{
		db: db,
	}
}
