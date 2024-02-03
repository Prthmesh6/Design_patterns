package main

import (
	"fmt"
)

//Open for extension but closed for modification
// Enterprice pattern specification

type Color int

const (
	red Color = iota
	green
	blue
)

// func (c Color) string() {
// 	fmt.Sprintf()
// }

type Size int

const (
	small Size = iota
	medium
	large
)

// as per single responsibillity the Product should only have methods like
// add product remove product(crud) and not the filters, so for filter we will create
// new method
type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f *Filter) FilterByColor(color Color, products []Product) []Product {
	filteredProducts := make([]Product, 0)
	for i := 0; i < len(products); i++ {
		if products[i].color == color {
			filteredProducts = append(filteredProducts, products[i])
		}
	}
	return filteredProducts
}

// -------- New Better Approach --------------

type Specification interface {
	isSatisfied(p *Product) bool
}

type colorSpecification struct {
	color Color
}

func (c *colorSpecification) isSatisfied(p *Product) bool {
	return c.color == p.color
}

type sizeSpecification struct {
	size Size
}

func (s *sizeSpecification) isSatisfied(p *Product) bool {
	return s.size == p.size
}

type BetterFilter struct{}

// this one generic function will work for all requirements
func (b *BetterFilter) filterBySpecification(s Specification, products []Product) []Product {
	filteredProducts := make([]Product, 0)
	for i := 0; i < len(products); i++ {
		if s.isSatisfied(&products[i]) {
			filteredProducts = append(filteredProducts, products[i])
		}
	}
	return filteredProducts
}

// ---------------------------------------------------------------

func main() {
	//Wrong approach
	Products := []Product{
		Product{name: "Tree", color: green, size: large},
		Product{name: "Apple", color: red, size: small},
		Product{name: "House", color: green, size: large},
	}

	oldFilter := Filter{}
	oldResponse := oldFilter.FilterByColor(green, Products)
	for i := range oldResponse {
		fmt.Printf("Name :- %s Color :- %v Size :- %v \n",
			oldResponse[i].name, oldResponse[i].color, oldResponse[i].size)
	}
	//Now the issue in this old approach is, If you have to add new filter
	//like filter by size, then you have to create new function FilterByColor
	//similarly if have to add FilterByColorAndSize then also you need to create
	//new function for it all these 3 functions will be different.

	//New Approach
	newFilter := BetterFilter{}
	colorSpecification := colorSpecification{color: green}
	newResponse := newFilter.filterBySpecification(&colorSpecification, Products)

	for i := range newResponse {
		fmt.Printf("Name :- %s Color :- %v Size :- %v \n",
			newResponse[i].name, newResponse[i].color, newResponse[i].size)
	}
}
