package main

import (
	"flag"
	"fmt"
)

const (
	countOfArgs = 3
)

func main() {
	var (
		userID    = flag.Int64("user", 0, "user id")
		productID = flag.Int64("product", 0, "product id")
		count     = flag.Int64("count", 0, "count")
	)

	flag.Parse()

	if nFlag := flag.NFlag(); nFlag != countOfArgs {
		fmt.Printf("oh no. have: %d, need: %d", nFlag, countOfArgs)
		return
	}

	AddToCart(*userID, *productID, *count)

}

func AddToCart(userID, productID, count int64) {
	fmt.Printf("user %d add product wuth id=%d (count %d)\n", userID, productID, count)
}
