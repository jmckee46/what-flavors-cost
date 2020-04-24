package main

import (
	"fmt"
	"testing"
)

func TestWhatFlavorsCost1(t *testing.T) {
	cost := []int32{1, 4, 5, 3, 2}
	money := int32(4)

	whatFlavors(cost, money)
	fmt.Println("Answer should be: 1 4")
}

func TestWhatFlavorsCost2(t *testing.T) {
	cost := []int32{2, 2, 4, 3}
	money := int32(4)

	whatFlavors(cost, money)
	fmt.Println("Answer should be: 1 2")
}
