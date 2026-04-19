package models

import (
	"fmt"
	"strings"
)

type ICart interface {
	GetCountByProductID(productID int64) int64
	fmt.Stringer
}

type CartMap map[int64]int64 // реализовали чтобы было удобно читать, чтобы каждый раз не писать map[int64]int64

func (c CartMap) String() string {
	builder := strings.Builder{} // .Используем потому что тратиться мало памяти чем при конкантенации строк(переаллокация)
	for productID, count := range c {
		builder.WriteString(fmt.Sprintf("product_id %d count %d\n", productID, count))
	}

	return builder.String()
}

func (c CartMap) GetCountByProductID(productID int64) int64 {
	return c[productID]
}
