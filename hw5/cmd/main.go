package main

import (
	"fmt"
	"go-search/hw5/pkg/geom"
	"go-search/hw5/pkg/point"
)

func main() {
	p1, err := point.New(10., 20.)
	if err != nil {
		fmt.Println(err)
		return
	}

	p2, err := point.New(40., 100.)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(geom.Distance(p1, p2))
}
