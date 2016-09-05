# Go-DAS

[![Build Status](https://travis-ci.org/Wen777/go-DSA.svg?branch=master)](https://travis-ci.org/Wen777/go-DSA)

Data Structure and Algorithm

# Data Structure

- [List](#lists)
    - [Singly Link List](#singlylinkedlist)
    - [Doubly Link List](#doublylinklist)
    - [Array Link List](#arraylinklist)
- [Queue]
- [Stack]

## Lists

### SinglyLinkedList

```go
type List struct {
    head, last *Node
    len        int
}
```

### DoublyLinkList

```golang
type List struct {
    head, last *element
    len        int
}
```

### ArrayLinkList

```go
type List struct {
    elements []interface{}
    size     int
}
```

## Queue

```go
type Queue struct {
    head, tail *node
    len        int
    lock       *sync.Mutex
}
```

###### Cocurrent Safty

## Stack

```go
type Stack struct {
    head *node
    lock *sync.Mutex
    len  int
}
```

###### Cocurrent Safty

# TODO

* [ ] [Set](https://en.wikipedia.org/wiki/Set_(abstract_data_type))
* [ ] Map
* [ ] [Red Black Tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)
* [ ] [Binary Heap](https://en.wikipedia.org/wiki/Binary_heap)
* [ ] [B+ Tree](https://en.wikipedia.org/wiki/B%2B_tree)
* [ ] [Tri Tree](https://en.wikipedia.org/wiki/Trie)
* [ ] [Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree)
