package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("match", func(t *testing.T) {

		got, err := dictionary.Search("test")
		want := "this is just a test"

		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})

	t.Run("no word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")
		assert.EqualError(t, err, ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		//given
		dictionary := Dictionary{}
		word := "test"
		def := "this is a just a test"

		//when
		err := dictionary.Add(word, def)
		assert.NoError(t, err)

		got, err := dictionary.Search(word)

		//then
		assert.NoError(t, err)
		assert.Equal(t, def, got)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "new test")
		assert.EqualError(t, err, ErrWordExists.Error())

		got, err := dictionary.Search(word)
		assert.NoError(t, err)
		assert.Equal(t, def, got)
	})

	t.Run("update", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{word: def}
		newDef := "new def"

		dictionary.Update(word, newDef)
		got, err := dictionary.Search(word)
		assert.NoError(t, err)
		assert.Equal(t, newDef, got)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, def)

		assert.EqualError(t, err, ErrWordDoesNotExist.Error())
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assert.EqualError(t, err, ErrNotFound.Error())
}
