package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	NAME        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sku         string  `json:"sku"`
	CreatedOn   string  `json:"createdOn"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (thisProduct *Product) FromJson(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(thisProduct) // Decode returns Error.
}

func (productList *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(productList)
}

// START : Service Methods for Adding Product to DataStore.

func AddSingleProduct(thisProd *Product) {
	thisProd.ID = getNextIdFromDB()

	productList = append(productList, thisProd)
}

func getNextIdFromDB() int {
	currentListOfProductsInDB := productList[len(productList)-1]
	return (currentListOfProductsInDB.ID + 1)
}

// END : Service Methods for Adding Product to DataStore.

// START : Service Methods for Updating Single Product to DataStore.

var ErrorProdNotFound = fmt.Errorf("Product Not found")

func UpdateSingleProduct(id int, p *Product) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[position] = p
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProdNotFound
}

// END : Service Methods for Updating Single Product to DataStore.

// START : Service Methods for GET ALL Product From DataStore.

func GetProducts() Products {
	return productList
}

var productList = []*Product{

	&Product{
		ID:          1,
		NAME:        "LATTE",
		Description: "Frothy Milky Coffee",
		Price:       2.45,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		NAME:        "ESSPRESSO",
		Description: "Strong coffe without Milk",
		Price:       1.99,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// END : Service Methods for GET ALL Product From DataStore.
