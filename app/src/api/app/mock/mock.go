package mock

import "api/app/models"

// ItemService ...
type ItemService struct {
	GetItemFn      func(id string) (*models.Item, error)
	GetItemInvoked bool

	GetItemsFn      func() ([]models.Item, error)
	GetItemsInvoked bool

	CreateItemFn      func(i *models.Item) error
	CreateItemInvoked bool

	DeleteItemFn      func(id string) error
	DeleteItemInvoked bool
}

// Item ...
func (is *ItemService) GetItem(id string) (*models.Item, error) {
	is.GetItemInvoked = true
	return is.GetItemFn(id)
}

// GetItems ...
func (is *ItemService) GetItems() ([]models.Item, error) {
	is.GetItemsInvoked = true
	return is.GetItemsFn()
}

// CreateItem ...
func (is *ItemService) CreateItem(i *models.Item) error {
	is.CreateItemInvoked = true
	return is.CreateItemFn(i)
}

// DeleteItem ...
func (is *ItemService) DeleteItem(id string) error {
	is.DeleteItemInvoked = true
	return is.DeleteItemFn(id)
}
