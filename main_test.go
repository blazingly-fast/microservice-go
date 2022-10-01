package main

import (
	"testing"

	"github.com/blazingly-fast/microservice-go/sdk/client"
	"github.com/blazingly-fast/microservice-go/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	c := client.Default
	params := products.NewListProductsParams()
	c.Products.ListProducts(params)
}
