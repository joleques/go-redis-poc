package integration

import (
	"github.com/joleques/go-redis-poc/src/useCase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOrderSuccessfully(t *testing.T) {
	identifier := "12_test"
	products := []string{"test_1", "test_2", "test_3", "test_4", "test_5"}
	err := useCase.CreateOrder(identifier, products)
	assert.Nil(t, err)
}

func TestCreateOrderWithErrorWhenIdentifierIsEmpty(t *testing.T) {
	identifier := ""
	products := []string{"test_1", "test_2", "test_3", "test_4", "test_5"}
	err := useCase.CreateOrder(identifier, products)
	assert.Equal(t, "error: identifier empty", err.Error())
}

func TestCreateOrderWithErrorWhenProductsIsEmpty(t *testing.T) {
	identifier := "12_test"
	products := []string{}
	err := useCase.CreateOrder(identifier, products)
	assert.Equal(t, "error: products empty", err.Error())
}

func TestCreateOrderWithErrorWhenProductsBiggerThen5(t *testing.T) {
	identifier := "12_test"
	products := []string{"test_1", "test_2", "test_3", "test_4", "test_5", "test_6"}
	err := useCase.CreateOrder(identifier, products)
	assert.Equal(t, "error: Product list cannot be longer than 5", err.Error())
}
