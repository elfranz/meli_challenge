package items

import (
	"api/app/mock"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api/app/models"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	// Inject our mock into our handler.
	var is mock.ItemService
	Is = &is

	// Define mocks.
	is.GetItemFn = func(id string) (*models.Item, error) {
		if id != "100" {
			t.Fatalf("unexpected id: %s", id)
		}
		return &models.Item{ID: "100", Name: "DaItam", Description: "Elnesto"}, nil
	}

	is.GetItemsFn = func() ([]models.Item, error) {
		return []models.Item{
			models.Item{ID: "100", Name: "DaItam", Description: "Elnesto"},
			models.Item{ID: "101", Name: "Pepe", Description: "This is an example."},
		}, nil
	}

	is.CreateItemFn = func(i *models.Item) error {
		return nil
	}

	is.DeleteItemFn = func(id string) error {
		if id != "100" {
			t.Fatalf("unexpected id: %s", id)
		}
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item/100", nil)
	router.ServeHTTP(w, r)

	// Validate mock.
	if !is.GetItemInvoked {
		t.Fatal("expected Item() to be invoked")
	}

	if w.Code != http.StatusOK {
		t.Fatal("expected ok status")
	}

	// get items

	s, _ := http.NewRequest("GET", "/item", nil)
	router.ServeHTTP(w, s)

	// Validate mock.
	if !is.GetItemsInvoked {
		t.Fatal("expected Item() to be invoked")
	}

	if w.Code != http.StatusOK {
		t.Fatal("expected ok status")
	}

	// create items
	payload := fmt.Sprintf(
		`{
			"name": "Pepe",
			"description": "This is an Example"
		}`,
	)

	u, _ := http.NewRequest("POST", "/item", strings.NewReader(payload))
	router.ServeHTTP(w, u)

	// Validate mock.
	if !is.CreateItemInvoked {
		t.Fatal("expectedItem() to be invoked")
	}

	if w.Code != http.StatusCreated {
		t.Fatal("expected created status")
	}

	// delete items

	v, _ := http.NewRequest("DELETE", "/item/100", nil)
	router.ServeHTTP(w, v)

	// Validate mock.
	if !is.DeleteItemInvoked {
		t.Fatal("expected Item() to be invoked")
	}

	if w.Code != http.StatusOK {
		t.Fatal("expected ok status")
	}
}
