package datastructures

type DoubleLinkedListNode[T any] struct {
    Val T;
    Prev *DoubleLinkedListNode[T];
    Next *DoubleLinkedListNode[T];
}
