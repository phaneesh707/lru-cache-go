package main

import(
)

// Deletes a item from back and adds it to front of cache
func (cache *Cache) DeleteAndAddToFront(k int32){
	//delete from back 
	ptr := cache.hash[k]
	leftNode := (*ptr).left
	rightNode := (*ptr).right
	
	(*leftNode).right = rightNode
	(*rightNode).left = leftNode

	// add it to front
	(*ptr).left = cache.head
	(*ptr).right = (*cache.head).right
	(*cache.head).right = ptr
	(*((*ptr).right)).left = ptr
}

func (cache *Cache) AddNewNode(k int32,v int32){
	var newNode  = Node{
		key: k,
		val: v,
		left: nil,
		right: nil,
	}
	var nodePtr = &newNode

	(*nodePtr).left = cache.head
	(*nodePtr).right = (*cache.head).right
	(*cache.head).right = nodePtr
	(*((*nodePtr).right)).left = nodePtr
	cache.hash[k] = nodePtr
}

func (cache *Cache) DeleteNode(k int32){
	ptr := cache.hash[k]
	leftNode := (*ptr).left
	rightNode := (*ptr).right
		
	(*leftNode).right = rightNode
	(*rightNode).left = leftNode
}