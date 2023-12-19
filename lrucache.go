package main 

import(
	"fmt"
)


func Get(){
	var k int32
	fmt.Print("Key : ")
	fmt.Scan(&k)

	val , ok := myMap[k]
	if ok{
		DeleteAndAddToFront(k)
		fmt.Printf("Key : %v    Value : %v\n",k,(*val).val)
	}else{
		fmt.Printf("Key %v does not exist!\n",k)
	}

}

func Set(){
	//input Key and val
	var k,v int32
	fmt.Print("key : ")
	fmt.Scan(&k)
	fmt.Print("Value: ")
	fmt.Scan(&v)

	// update if key present
	_,ok := myMap[k]
	if ok {
		(*(myMap[k])).val = v
		DeleteAndAddToFront(k)
		return
	}

	// if cache is full remove the least recently used item
	if cacheSize == capacity {
		//delete the last node 
		keyToDel := (*((*tail).left)).key

		ptr := myMap[keyToDel]
		leftNode := (*ptr).left
		rightNode := (*ptr).right
		
		(*leftNode).right = rightNode
		(*rightNode).left = leftNode

		delete(myMap,keyToDel)
		
	}
	
	var newNode  = node{
		key: k,
		val: v,
		left: nil,
		right: nil,
	}
	var nodePtr = &newNode

	(*nodePtr).left = head
	(*nodePtr).right = (*head).right
	(*head).right = nodePtr
	(*((*nodePtr).right)).left = nodePtr
	myMap[k] = nodePtr

	// increase size if size less than cap
	if cacheSize<capacity {
		cacheSize+=1
	}
	
}


// prints the current state of cache
func PrintState(){
	var nodePtr = (*head).right
	fmt.Print("Head<--->")
	for nodePtr != tail {
		fmt.Printf("(%v,%v)<--->",(*nodePtr).key,(*nodePtr).val)
		nodePtr = (*nodePtr).right
	}
	fmt.Print("Tail")
}


// clear the cache
func ClearCache(){
	myMap = make(map[int32](*node))
	(*head).right = tail
	(*tail).left = head
	cacheSize = 0
}