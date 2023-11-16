package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/eleniums/mining-post/mem"
	"github.com/eleniums/mining-post/models"
	"github.com/go-chi/chi"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Server_InsertItem_Success(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	req := &models.Item{
		Category: 123,
		Data:     "test-data",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	payload, err := json.Marshal(req)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(payload))

	// act
	server.InsertItem(w, rq)

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)

	var resp models.Item
	err = unmarshalBody(r, &resp)
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.ID)
	assert.Equal(t, req.Category, resp.Category)
	assert.Equal(t, req.Data, resp.Data)
	assert.ElementsMatch(t, req.Tags, resp.Tags)
	assertKeyValuesMatch(t, req.Metadata, resp.Metadata)

	item, err := cache.Get(resp.ID)
	assert.NoError(t, err)
	assert.Equal(t, req.Category, item.Category)
	assert.Equal(t, req.Data, item.Data)
	assert.ElementsMatch(t, req.Tags, item.Tags)
	assertKeyValuesMatch(t, req.Metadata, item.Metadata)
}

func Test_Unit_Server_InsertItem_PopulatedID(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	req := &models.Item{
		ID:       "d24ac3e4-6323-4099-a8d6-15c01cdfe805",
		Category: 123,
		Data:     "test-data",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	payload, err := json.Marshal(req)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(payload))

	// act
	server.InsertItem(w, rq)

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)

	var resp models.Item
	err = unmarshalBody(r, &resp)
	assert.NoError(t, err)

	assert.NotEqual(t, req.ID, resp.ID)
	assert.Equal(t, req.Category, resp.Category)
	assert.Equal(t, req.Data, resp.Data)
	assert.ElementsMatch(t, req.Tags, resp.Tags)
	assertKeyValuesMatch(t, req.Metadata, resp.Metadata)

	item, err := cache.Get(resp.ID)
	assert.NoError(t, err)
	assert.Equal(t, req.Category, item.Category)
	assert.Equal(t, req.Data, item.Data)
	assert.ElementsMatch(t, req.Tags, item.Tags)
	assertKeyValuesMatch(t, req.Metadata, item.Metadata)
}

func Test_Unit_Server_GetItemByID_Success(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	expected := &mem.Item{
		Category: 123,
		Data:     "test-data",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	expectedID, _ := cache.Add(expected)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodGet, path.Join("/items/", expectedID), nil)
	ctx := createContextWithItemID(expectedID)

	// act
	server.GetItemByID(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)

	var resp models.Item
	err := unmarshalBody(r, &resp)
	assert.NoError(t, err)

	assert.Equal(t, expectedID, resp.ID)
	assert.Equal(t, expected.Category, resp.Category)
	assert.Equal(t, expected.Data, resp.Data)
	assert.ElementsMatch(t, expected.Tags, resp.Tags)
	assertKeyValuesMatch(t, expected.Metadata, resp.Metadata)
}

func Test_Unit_Server_GetItemByID_EmptyID(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodGet, "/items/", nil)
	ctx := createContextWithItemID("")

	// act
	server.GetItemByID(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
}

func Test_Unit_Server_GetItemByID_InvalidID(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodGet, "/items/123456", nil)
	ctx := createContextWithItemID("123456")

	// act
	server.GetItemByID(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusNotFound, r.StatusCode)
}

func Test_Unit_Server_GetItems_Success(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	expected1 := &mem.Item{
		Data: "test-data1",
	}
	expected2 := &mem.Item{
		Data: "test-data2",
	}
	expected3 := &mem.Item{
		Data: "test-data3",
	}

	expectedID1, _ := cache.Add(expected1)
	expectedID2, _ := cache.Add(expected2)
	expectedID3, _ := cache.Add(expected3)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodGet, "/items", nil)

	// act
	server.GetItems(w, rq)

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)

	var resp models.Items
	err := unmarshalBody(r, &resp)
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.Items)
	assert.Equal(t, 3, len(resp.Items))

	item1 := containsItem(t, resp.Items, expectedID1)
	assert.Equal(t, expected1.Data, item1.Data)

	item2 := containsItem(t, resp.Items, expectedID2)
	assert.Equal(t, expected2.Data, item2.Data)

	item3 := containsItem(t, resp.Items, expectedID3)
	assert.Equal(t, expected3.Data, item3.Data)
}

func Test_Unit_Server_UpdateItem_Success(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	original := &mem.Item{
		Category: 123,
		Data:     "test-data",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	originalID, _ := cache.Add(original)

	req := &models.Item{
		ID:       originalID,
		Category: 456,
		Data:     "updated",
		Tags: []string{
			"tag4",
			"tag5",
			"tag6",
		},
		Metadata: map[string]string{
			"key4": "value4",
			"key5": "value5",
			"key6": "value6",
		},
	}

	payload, err := json.Marshal(req)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodPut, path.Join("/items/", originalID), bytes.NewBuffer(payload))
	ctx := createContextWithItemID(originalID)

	// act
	server.UpdateItem(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)

	item, err := cache.Get(originalID)
	assert.NoError(t, err)
	assert.NotNil(t, item)

	assert.Equal(t, req.Category, item.Category)
	assert.Equal(t, req.Data, item.Data)
	assert.ElementsMatch(t, req.Tags, item.Tags)
	assertKeyValuesMatch(t, req.Metadata, item.Metadata)
}

func Test_Unit_Server_UpdateItem_EmptyID(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	req := &models.Item{
		ID:       "",
		Category: 456,
		Data:     "updated",
		Tags: []string{
			"tag4",
			"tag5",
			"tag6",
		},
		Metadata: map[string]string{
			"key4": "value4",
			"key5": "value5",
			"key6": "value6",
		},
	}

	payload, err := json.Marshal(req)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodPut, "/items/", bytes.NewBuffer(payload))
	ctx := createContextWithItemID("")

	// act
	server.UpdateItem(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
}

func Test_Unit_Server_UpdateItem_InvalidID(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	req := &models.Item{
		ID:       "123456",
		Category: 456,
		Data:     "updated",
		Tags: []string{
			"tag4",
			"tag5",
			"tag6",
		},
		Metadata: map[string]string{
			"key4": "value4",
			"key5": "value5",
			"key6": "value6",
		},
	}

	payload, err := json.Marshal(req)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodPut, "/items/123456", bytes.NewBuffer(payload))
	ctx := createContextWithItemID("123456")

	// act
	server.UpdateItem(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusNotFound, r.StatusCode)
}

func Test_Unit_Server_DeleteItem_Success(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	expected := &mem.Item{
		Category: 123,
		Data:     "test-data",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
		Metadata: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}

	expectedID, _ := cache.Add(expected)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodDelete, path.Join("/items/", expectedID), nil)
	ctx := createContextWithItemID(expectedID)

	// act
	server.DeleteItem(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)

	item, err := cache.Get(expectedID)
	assert.Error(t, err)
	assert.Nil(t, item)
	assert.Equal(t, mem.ErrNotFound, err)
}

func Test_Unit_Server_DeleteItem_EmptyID(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodDelete, "/items/", nil)
	ctx := createContextWithItemID("")

	// act
	server.DeleteItem(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
}

func Test_Unit_Server_DeleteItem_InvalidID(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodDelete, "/items/123456", nil)
	ctx := createContextWithItemID("123456")

	// act
	server.DeleteItem(w, rq.WithContext(ctx))

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)
}

// containsItem will return an item from the array with a matching id or fail an assertion if it doesn't exist.
func containsItem(t assert.TestingT, items []*models.Item, id string) *models.Item {
	for _, v := range items {
		if v.ID == id {
			return v
		}
	}
	assert.Failf(t, "item was not found in array with id: %s", id)
	return nil
}

// assertKeyValuesMatch asserts that the keys and values of two maps are equal.
func assertKeyValuesMatch(t assert.TestingT, mapA map[string]string, mapB map[string]string) {
	assert.Equal(t, len(mapA), len(mapB), "map lengths are not equal")
	for k, v := range mapA {
		val, ok := mapB[k]
		assert.True(t, ok, "key does not exist in map: %v", k)
		assert.Equal(t, v, val, "map values do not match for key: %v, %v != %v", k, v, val)
	}
}

// createContextWithItemID creates a new context with the item id added to the url params.
func createContextWithItemID(id string) context.Context {
	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("itemID", id)
	return context.WithValue(context.Background(), chi.RouteCtxKey, routeCtx)
}
