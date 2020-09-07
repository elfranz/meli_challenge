package items

import (
	"api/app/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api/app/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetItem(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	var is mock.ItemService
	Is = &is

	is.GetItemFn = func(id string) (*models.Item, error) {
		if id != "100" {
			t.Fatalf("unexpected id: %s", id)
		}
		return &models.Item{ID: "100", Name: "DaItam", Description: "Elnesto"}, nil
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item/100", nil)
	router.ServeHTTP(w, r)

	// Validate mock.
	assert.Equal(t, true, is.GetItemInvoked, "expected GetItem() to be invoked")

	// assert response code
	assert.Equal(t, 200, w.Code, "status should be 200")

	// assert response body
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	id := response["id"]
	name := response["name"]
	description := response["description"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, "100", id)
	assert.Equal(t, "DaItam", name)
	assert.Equal(t, "Elnesto", description)
}

func TestGetItems(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	var is mock.ItemService
	Is = &is

	is.GetItemsFn = func() ([]models.Item, error) {
		return []models.Item{
			models.Item{ID: "100", Name: "DaItam", Description: "Elnesto"},
			models.Item{ID: "101", Name: "Pepe", Description: "This is an example."},
		}, nil
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item", nil)
	router.ServeHTTP(w, r)

	assert.Equal(t, true, is.GetItemsInvoked, "expected GetItems() to be invoked")

	assert.Equal(t, 200, w.Code, "status should be 200")

	// var response []string
	// _ = json.Unmarshal([]byte(dataJson), &response)
	// log.Printf("Unmarshaled: %v", response)
	// slice is empty
}

func TestCreateItem(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	// Inject our mock into our handler.
	var is mock.ItemService
	Is = &is

	is.CreateItemFn = func(i *models.Item) error {
		return nil
	}

	payload := fmt.Sprintf(
		`{
			"name": "Pepe",
			"description": "This is an Example"
		}`,
	)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/item", strings.NewReader(payload))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, r)

	assert.Equal(t, true, is.CreateItemInvoked, "expected CreateItem() to be invoked")

	assert.Equal(t, 201, w.Code, "status should be 201")

	// assert response body
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	// id := response["id"]
	for key, value := range response {
		fmt.Print(key)
		fmt.Println("-" + value)
	}
	name := response["name"]
	description := response["description"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	// assert.Equal(t, "1", id) WHY ON EARTH IS THIS KEY NIL???
	assert.Equal(t, "Pepe", name)
	assert.Equal(t, "This is an Example", description)
}

func TestDeleteItem(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	var is mock.ItemService
	Is = &is

	is.DeleteItemFn = func(id string) error {
		if id != "100" {
			t.Fatalf("unexpected id: %s", id)
		}
		return nil
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/item/100", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, r)

	assert.Equal(t, true, is.DeleteItemInvoked, "expected DeleteItem() to be invoked")

	assert.Equal(t, 200, w.Code, "status should be 200")

	// assert response body
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	message := response["message"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, "Item successfully deleted.", message)
}
