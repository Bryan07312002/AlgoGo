package datastructures

type LikedListNode[T any] struct {
    Val T;
    Next *LikedListNode[T];
}
