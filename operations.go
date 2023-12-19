package main

import(
)

// Deletes a node from back and adds it to front of cache
func DeleteAndAddToFront(k int32){
	//delete from back 
	ptr := myMap[k]
	leftNode := (*ptr).left
	rightNode := (*ptr).right
	
	(*leftNode).right = rightNode
	(*rightNode).left = leftNode

	// add it to front
	(*ptr).left = head
	(*ptr).right = (*head).right
	(*head).right = ptr
	(*((*ptr).right)).left = ptr
}
