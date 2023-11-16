package mem

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Cache_Add_Success(t *testing.T) {
	// arrange
	cache := NewCache()

	item := &Item{
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

	// act
	id, err := cache.Add(item)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func Test_Unit_Cache_Get_Success(t *testing.T) {
	// arrange
	cache := NewCache()

	expected := &Item{
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

	id, err := cache.Add(expected)
	assert.NoError(t, err)

	// act
	item, err := cache.Get(id)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, expected, item)
}

func Test_Unit_Cache_Get_NotFound(t *testing.T) {
	// arrange
	cache := NewCache()

	// act
	item, err := cache.Get(newUUID())

	// assert
	assert.Error(t, err)
	assert.Nil(t, item)
	assert.Equal(t, ErrNotFound, err)
}

func Test_Unit_Cache_GetAll_Success(t *testing.T) {
	// arrange
	cache := NewCache()

	expectedItem1 := &Item{
		Category: 123,
		Data:     "test-data1",
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
	expectedItem1ID, err := cache.Add(expectedItem1)
	assert.NoError(t, err)

	expectedItem2 := &Item{
		Category: 456,
		Data:     "test-data2",
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
	expectedItem2ID, err := cache.Add(expectedItem2)
	assert.NoError(t, err)

	// act
	items, err := cache.GetAll()

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, items)

	item1, ok := items[expectedItem1ID]
	assert.True(t, ok)
	assert.Equal(t, expectedItem1, item1)

	item2, ok := items[expectedItem2ID]
	assert.True(t, ok)
	assert.Equal(t, expectedItem2, item2)
}

func Test_Unit_Cache_Update_Success(t *testing.T) {
	// arrange
	cache := NewCache()

	original := &Item{
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

	id, err := cache.Add(original)
	assert.NoError(t, err)

	updated := &Item{
		Category: 456,
		Data:     "updated-data",
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

	// act
	err = cache.Update(id, updated)

	// assert
	assert.NoError(t, err)

	item, err := cache.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, updated, item)
}

func Test_Unit_Cache_Update_NotFound(t *testing.T) {
	// arrange
	cache := NewCache()

	updated := &Item{
		Category: 456,
		Data:     "updated-data",
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

	// act
	err := cache.Update(newUUID(), updated)

	// assert
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)
}

func Test_Unit_Cache_Remove_Success(t *testing.T) {
	// arrange
	cache := NewCache()

	item := &Item{
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

	id, err := cache.Add(item)
	assert.NoError(t, err)

	// act
	err = cache.Remove(id)

	// assert
	assert.NoError(t, err)

	// make sure item was actually removed
	retrieved, err := cache.Get(id)
	assert.Nil(t, retrieved)
	assert.Equal(t, ErrNotFound, err)
}

func Test_Unit_Cache_Remove_NotFound(t *testing.T) {
	// arrange
	cache := NewCache()

	// act
	err := cache.Remove(newUUID())

	// assert
	assert.NoError(t, err) // even though there was nothing to remove, do not expect an error
}
