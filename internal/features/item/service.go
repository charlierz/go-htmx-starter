package item

import (
	"database/sql"

	"github.com/charlierz/go-htmx-starter/internal/model"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) List() ([]model.Item, error) {
	rows, err := s.db.Query(
		`SELECT id, name, created_at FROM items`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []model.Item{}

	for rows.Next() {
		item := model.Item{}
		err = rows.Scan(&item.Id, &item.Name, &item.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) Create(newItem model.NewItem) (model.Item, error) {
	row := s.db.QueryRow(
		`INSERT INTO items (name) VALUES (?) RETURNING id, name, created_at`,
		newItem.Name,
	)
	item := model.Item{}
	err := row.Scan(&item.Id, &item.Name, &item.CreatedAt)
	return item, err
}

func (s *Service) Delete(itemId string) error {
	_, err := s.db.Exec(`DELETE FROM items WHERE id = ?`, itemId)
	return err
}
