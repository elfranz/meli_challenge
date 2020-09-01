package items

import (
	"api/app/models"
	"database/sql"
	"log"
	"strconv"
)

// ItemService ...
type ItemService struct {
	DB *sql.DB
}

// GetItem ...
func (s *ItemService) GetItem(id string) (*models.Item, error) {
	var i models.Item
	row := s.DB.QueryRow(`SELECT id, name, description FROM items WHERE id = ?`, id)
	if err := row.Scan(&i.ID, &i.Name, &i.Description); err != nil {
		return nil, err
	}
	return &i, nil
}

// GetItems ...
func (s *ItemService) GetItems() ([]models.Item, error) {
	var i models.Item
	rows, err := s.DB.Query(`SELECT id, name, description FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		err := rows.Scan(&i.ID, &i.Name, &i.Description)
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return items, nil
}

// CreateItem ...
func (s *ItemService) CreateItem(i *models.Item) error {
	stmt, err := s.DB.Prepare(`INSERT INTO items(name,description) values(?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(i.Name, i.Description)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	i.ID = strconv.Itoa(int(id))
	return nil
}

// DeleteItem ...
func (s *ItemService) DeleteItem(id string) error {
	stmt, err := s.DB.Prepare(`DELETE FROM items WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	v, err := stmt.Exec(id)
	_ = v
	if err != nil {
		return err
	}
	return nil
}
