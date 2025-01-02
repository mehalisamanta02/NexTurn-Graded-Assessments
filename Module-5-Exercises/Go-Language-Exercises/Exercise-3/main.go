package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

// Add product to inventory
func addProduct(id int, name string, price float64, stock int) error {
	for _, prod := range inventory {
		if prod.ID == id {
			return errors.New("product ID must be unique")
		}
	}
	if price < 0 || stock < 0 {
		return errors.New("price and stock must be non-negative")
	}
	inventory = append(inventory, Product{ID: id, Name: name, Price: price, Stock: stock})
	return nil
}

// Update stock of a product
func updateStock(id int, stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, prod := range inventory {
		if prod.ID == id {
			inventory[i].Stock = stock
			return nil
		}
	}
	return errors.New("product not found")
}

// Search product by ID or name
func searchProduct(query string) (*Product, error) {
	for _, prod := range inventory {
		if fmt.Sprint(prod.ID) == query || prod.Name == query {
			return &prod, nil
		}
	}
	return nil, errors.New("product not found")
}

// Display inventory
func displayInventory() {
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, prod := range inventory {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", prod.ID, prod.Name, prod.Price, prod.Stock)
	}
}

// Sort inventory by price or stock
func sortInventory(by string) {
	if by == "Price" {
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	} else if by == "Stock" {
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	} else {
		fmt.Println("Invalid sort criteria. Use 'Price' or 'Stock'.")
	}
}

func main() {
	// Adding products to inventory
	if err := addProduct(1001, "Laptop", 57000.50, 50); err != nil {
		fmt.Println("Error adding product:", err)
	}
	if err := addProduct(2001, "Desktop", 125000.00, 20); err != nil {
		fmt.Println("Error adding product:", err)
	}
	if err := addProduct(3001, "Keyboard", 450.00, 70); err != nil {
		fmt.Println("Error adding product:", err)
	}
	if err := addProduct(4001, "Tab", 40050.00, 25); err != nil {
		fmt.Println("Error adding product:", err)
	}

	// Display inventory
	fmt.Println("Initial Inventory:")
	displayInventory()

	// Update stock
	if err := updateStock(1001, 8); err != nil {
		fmt.Println("Error updating stock!", err)
	}

	// Search for products
	for {
		var searchID int
		fmt.Println("\nEnter the product ID to search:")
		fmt.Scanln(&searchID)
		prod, err := searchProduct(fmt.Sprint(searchID))
		if err != nil {
			fmt.Println("Error searching product:", err)
		} else {
			fmt.Printf("Found product: %+v\n", *prod)
		}

		var searchAgain string
		fmt.Println("Do you want to search for another product? (yes/no):")
		fmt.Scanln(&searchAgain)
		if strings.ToLower(searchAgain) != "yes" {
			break
		}
	}

	// Ask user if they want to sort the inventory
	var sortChoice string
	fmt.Println("\nDo you want to sort the inventory? (yes/no):")
	fmt.Scanln(&sortChoice)
	if strings.ToLower(sortChoice) == "yes" {
		var sortBy string
		fmt.Println("How would you like to sort the inventory? (Price/Stock):")
		fmt.Scanln(&sortBy)
		fmt.Printf("\nInventory sorted by %s:\n", sortBy)
		sortInventory(sortBy)
		displayInventory()
	} else {
		fmt.Println("Skipping sorting.")
	}
}
