package main

import (
	"fmt"

	"./book"
)

func main() {
	url := "http://localhost:9200"
	esclient, _ := book.NewESClient(url)
	bm := book.NewBookManager(esclient)
	// create new book
	books := createSomeBooks()
	// insert to es
	var err error
	insertedBookCount := 0
	for _, b := range books {
		err = bm.AddBook(b)
		if err != nil {
			fmt.Println("Cannot insert book: ", b, err)
			continue
		}
		insertedBookCount++
	}
	fmt.Printf("Inserted: %v books \n", insertedBookCount)

	// search
	// case 1: found
	resultSearchBooksSuccess := bm.SearchBooks("One Stop")
	fmt.Println("Found books: ", resultSearchBooksSuccess)
	// case 2: not found
	resultSearchBooksFailed := bm.SearchBooks("opencommerce group")
	fmt.Println("Found books: ", resultSearchBooksFailed)
	// delete
	bm.DeleteBook(&book.Book{ID: "1"})
	// search again, but not found
	resultSearchBooksSuccess = bm.SearchBooks("One Stop")
	fmt.Println("Found books: ", resultSearchBooksSuccess)
}

func createSomeBooks() []*book.Book {
	return []*book.Book{
		&book.Book{
			ID:            "1",
			ISBN:          "1250244498",
			Title:         "One Last Stop",
			Author:        "Casey McQuiston",
			NumberOfPage:  432,
			PublishedDate: "June 1, 2021",
			Tags:          "time travel, romance",
			Brief:         "For cynical twenty-three-year-old August, moving to New York City is supposed to prove her right: that things like magic and cinematic love stories don’t exist, and the only smart way to go through life is alone. She can’t imagine how waiting tables at a 24-hour pancake diner and moving in with too many weird roommates could possibly change that. And there’s certainly no chance of her subway commute being anything more than a daily trudge through boredom and electrical failures.",
		},
		&book.Book{
			ID:            "2",
			ISBN:          "B091J2CP5X",
			Title:         "The Bomber Mafia: A Dream, a Temptation, and the Longest Night of the Second World War",
			Author:        "Malcolm Gladwell",
			NumberOfPage:  368,
			PublishedDate: "April 27, 2021",
			Tags:          "military, australian",
			Brief:         "In The Bomber Mafia: A Dream, a Temptation, and the Longest Night of the Second World War, Malcolm Gladwell, author of New York Times best sellers including Talking to Strangers and host of the podcast Revisionist History, uses original interviews, archival footage, and his trademark insight to weave together the stories of a Dutch genius and his homemade computer, a band of brothers in central Alabama, a British psychopath, and pyromaniacal chemists at Harvard. As listeners hear these stories unfurl, Gladwell examines one of the greatest moral challenges in modern American history",
		},
		&book.Book{
			ID:            "3",
			ISBN:          "1623159911",
			Title:         "Crystals for Beginners: The Guide to Get Started with the Healing Power of Crystals",
			Author:        "Karen Frazier",
			NumberOfPage:  206,
			PublishedDate: "October 17, 2017",
			Tags:          "rock and mineral, divination",
			Brief:         "",
		},
	}
}
