package models

// Item ...
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ItemServiceInterface ...
type ItemServiceInterface interface {
	GetItem(id string) (*Item, error)
	GetItems() ([]Item, error)
	CreateItem(i *Item) error
	DeleteItem(id string) error
}
