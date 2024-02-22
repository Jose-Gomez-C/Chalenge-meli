package repositories

import (
	"database/sql"

	"github.com/Jose-Gomez-c/challenge/api/model"
)

type ItemRepository interface {
	Save(item model.Items) (int, error)
	SaveInBatch(items []model.Items) ([]int, error)
}

type itemRepositoryLayer struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepositoryLayer{db: db}
}

func (layer itemRepositoryLayer) Save(item model.Items) (int, error) {
	query := "INSERT INTO items (id, siteId, price, name_category, description, nickname) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := layer.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(item.Id, item.SiteId, item.Price, item.NameCategory, item.Description, item.Nickname)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (layer itemRepositoryLayer) SaveInBatch(items []model.Items) ([]int, error) {
	query := "INSERT INTO items (id, site_id, price, name_category, description, nickname) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := layer.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	tx, err := layer.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var ids []int
	for _, item := range items {
		res, err := stmt.Exec(item.Id, item.SiteId, item.Price, item.NameCategory, item.Description, item.Nickname)
		if err != nil {
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		ids = append(ids, int(id))
	}
	return ids, nil
}
