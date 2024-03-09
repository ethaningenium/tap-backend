package main

import "fmt"

type Request struct {
	UserID string `json:"user_id"`
	PageID string `json:"page_id"`
}

type Page struct {
	Title string `json:"title"`
	Address string `json:"address"`
	UserID string `json:"user_id"`
	PageID string `json:"page_id"`
}

func main() {
	fmt.Println("Hello, World! from scripts/main.go")
}