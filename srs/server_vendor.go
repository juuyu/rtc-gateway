// Package srs
// @title
// @description
// @author njy
// @since 2023/6/1 11:24
package srs

type Vendor int8

const (
	Tx Vendor = iota + 1 // EnumIndex = 1
	AL                   // EnumIndex = 2
)

func (v Vendor) String() string {
	return [...]string{"TX", "AL"}[v-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (v Vendor) EnumIndex() int {
	return int(v)
}
