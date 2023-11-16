package mem

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
)

var (
	ErrNotFound = errors.New("error - not found")
)

type Item struct {
	Category int32
	Data     string
	Tags     []string
	Metadata map[string]string
}

// Cache temporarily stores items in memory.
type Cache struct {
	cache *sync.Map
}

// NewCache creates a new cache.
func NewCache() *Cache {
	return &Cache{
		cache: &sync.Map{},
	}
}

// Add a new item with a unique id.
func (c *Cache) Add(item *Item) (string, error) {
	id := newUUID()

	c.cache.Store(id, clone(item))

	return id, nil
}

// Update an existing item.
func (c *Cache) Update(id string, item *Item) error {
	_, ok := c.cache.Load(id)
	if !ok {
		return ErrNotFound
	}

	c.cache.Store(id, clone(item))

	return nil
}

// Get retrieves an item.
func (c *Cache) Get(id string) (*Item, error) {
	item, ok := c.cache.Load(id)
	if !ok {
		return nil, ErrNotFound
	}

	return clone(item.(*Item)), nil
}

// GetAll retrieves all items.
func (c *Cache) GetAll() (map[string]*Item, error) {
	items := map[string]*Item{}
	c.cache.Range(func(k, v interface{}) bool {
		items[k.(string)] = clone(v.(*Item))
		return true
	})

	return items, nil
}

// Remove an item.
func (c *Cache) Remove(id string) error {
	c.cache.Delete(id)

	return nil
}

// newUUID will generate a new v4 uuid.
func newUUID() string {
	var err error
	for i := 0; i < 3; i++ {
		var u uuid.UUID
		u, err = uuid.NewV4()
		if err == nil {
			return u.String()
		}
	}
	panic(fmt.Sprintf("Error generating uuid: %v", err))
}

// clone will return a deep copy of the given item.
func clone(item *Item) *Item {
	new := &Item{
		Category: item.Category,
		Data:     item.Data,
		Tags:     []string{},
		Metadata: map[string]string{},
	}

	new.Tags = append(new.Tags, item.Tags...)

	for k, v := range item.Metadata {
		new.Metadata[k] = v
	}

	return new
}
