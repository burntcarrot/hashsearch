package repository

import (
	"fmt"
	"strconv"

	"github.com/burntcarrot/hashsearch/entity"
	"github.com/tidwall/buntdb"
)

// BuntDB is the repository using buntdb.
type BuntDB struct {
	db *buntdb.DB
}

// NewBuntDB creates a new repository using buntdb.
func NewBuntDB(db *buntdb.DB) *BuntDB {
	return &BuntDB{
		db: db,
	}
}

// Create creates an image entry.
func (r *BuntDB) Create(e *entity.Image) error {
	err := r.db.Update(func(tx *buntdb.Tx) error {
		hash := fmt.Sprintf("%064b", e.Hash)
		_, _, err := tx.Set(e.Path, hash, nil)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

// List returns a list of image entries.
func (r *BuntDB) List() ([]*entity.Image, error) {
	images := []*entity.Image{}

	err := r.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			hash, err := strconv.ParseUint(value, 2, 64)
			if err != nil {
				return false
			}

			img := entity.Image{Path: key, Hash: hash}
			images = append(images, &img)

			return true
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	return images, nil
}
