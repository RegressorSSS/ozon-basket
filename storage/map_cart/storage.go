package map_cart

import (
	"fmt"
	"strings"
)

type Cart map[int64]int64 // реализовали чтобы было удобно читать, чтобы каждый раз не писать map[int64]int64

func (c Cart) String() string {
	builder := strings.Builder{} // .Используем потому что тратиться мало памяти чем при конкантенации строк(переаллокация)
	for productID, count := range c {
		builder.WriteString(fmt.Sprintf("product_id %d count %d\n", productID, count))
	}

	return builder.String()
}

type Storage struct {
	carts map[int64]Cart
}

func New() *Storage {
	return &Storage{
		carts: make(map[int64]Cart),
	}
}

func (s *Storage) AddToCart(userID, productID, count int64) {
	if _, ok := s.carts[userID]; !ok {
		s.carts[userID] = make(Cart)
	}
	s.carts[userID][productID] += count
}

func (s *Storage) GetCartByUserID(userID int64) Cart {
	return s.carts[userID]
}
