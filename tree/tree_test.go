package tree_test

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var data = tree{
		Val: 1,
		Left: &tree{
			Val: 2,
			Left: &tree{
				Val: 4,
			},
			Right: &tree{
				Val: 5,
			},
		},
		Right: &tree{
			Val: 3,
			Left: &tree{
				Val: 6,
			},
			Right: &tree{
				Val: 7,
			},
		},
	}

	printlnData(&data)
	fmt.Println(1111)
	printlnData(changeTree(&data))
}

type tree struct {
	Val   int
	Left  *tree
	Right *tree
}

func printlnData(data *tree) int {
	fmt.Println(data.Val)
	if data == nil || (data.Left == nil && data.Right == nil) {
		return data.Val
	}
	printlnData(data.Left)
	printlnData(data.Right)
	return data.Val
}

func changeTree(data *tree) *tree {
	if data == nil || (data.Left == nil && data.Right == nil) {
		return data
	}
	left := changeTree(data.Left)
	right := changeTree(data.Right)
	data.Left = right
	data.Right = left
	return data
}
