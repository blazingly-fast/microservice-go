package handlers

import (
	"net/http"

	"github.com/blazingly-fast/microservice-go/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := getProductID(r)
	p.l.Print("log")
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	prod.ID = id
	p.l.Println("[DEBUG] updating record id", prod.ID)
	err := p.db.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, w)
		return
	}

	// write the no content success header
	w.WriteHeader(http.StatusNoContent)
}
