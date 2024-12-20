package datastructures

type BinaryTreeNode[T any] struct {
	Val   T
	Left  *BinaryTreeNode[T]
	Rigth *BinaryTreeNode[T]
}
