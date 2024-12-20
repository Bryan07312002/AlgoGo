package datastructures

// The root element will be at arr[0].
// The below table shows indices of other nodes for the ith node, i.e., arr[i]:
// arr[(i-1)/2]	Returns the parent node
// arr[(2*i)+1]	Returns the left child node
// arr[(2*i)+2]	Returns the right child node
// The traversal method use to achieve Array representation is Level Order Traversal. Please refer to Array Representation Of Binary Heap for details.
//
// Binary Heap Tree
//
// Operations on Heap:
// Below are some standard operations on min heap:
//
// getMin(): It returns the root element of Min Heap. The time Complexity of
// this operation is O(1). In case of a maxheap it would be getMax().
//
// extractMin(): Removes the minimum element from MinHeap. The time Complexity
// of this Operation is O(log N) as this operation needs to maintain the heap
// property (by calling heapify()) after removing the root.
//
// decreaseKey(): Decreases the value of the key. The time complexity of this
// operation is O(log N). If the decreased key value of a node is greater than
// the parent of the node, then we don’t need to do anything. Otherwise, we need
// to traverse up to fix the violated heap property.
//
// insert(): Inserting a new key takes O(log N) time. We add a new key at the
// end of the tree. If the new key is greater than its parent, then we don’t
// need to do anything. Otherwise, we need to traverse up to fix the violated
// heap property.
//
// delete(): Deleting a key also takes O(log N) time. We replace the key to be
// deleted with the minimum infinite by calling decreaseKey(). After
// decreaseKey(), the minus infinite value must reach root, so we call
// extractMin() to remove the key
