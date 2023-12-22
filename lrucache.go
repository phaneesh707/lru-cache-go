package main 

import(
	"fmt"
)


func (cache *Cache) Get(){
	var k int32
	fmt.Print("Key : ")
	fmt.Scan(&k)

	val , ok := cache.hash[k]
	if ok{
		cache.DeleteAndAddToFront(k)
		fmt.Printf("Key : %v    Value : %v\n",k,(*val).val)
	}else{
		fmt.Printf("Key %v does not exist!\n",k)
	}
}


func (cache *Cache) Set(){
	//input Key and val
	var k,v int32
	fmt.Print("key : ")
	fmt.Scan(&k)
	fmt.Print("Value: ")
	fmt.Scan(&v)

	// update if key present
	_,ok := cache.hash[k]
	if ok {
		(*(cache.hash[k])).val = v
		cache.DeleteAndAddToFront(k)
		return
	}

	// if cache is full remove the least recently used item
	if cache.capacity == int32(len(cache.hash)) {
		//delete the last node 
		keyToDel := (*((*cache.tail).left)).key

		cache.DeleteNode(keyToDel)

		delete(cache.hash,keyToDel)
	}
	
	//Add the item to cache
	cache.AddNewNode(k,v)

	// increase size if size less than cap
	if int32(len(cache.hash)) < cache.capacity {
		cache.capacity+=1
	}

	cache.PrintState()
	
}


// prints the current state of cache
func (cache *Cache) PrintState(){
	var nodePtr = (*(cache.head)).right
	fmt.Print("Head<--->")
	for nodePtr != cache.tail {
		fmt.Printf("(%v,%v)<--->",(*nodePtr).key,(*nodePtr).val)
		nodePtr = (*nodePtr).right
	}
	fmt.Print("Tail")
}


// clear the cache
func (cache *Cache) ClearCache(){
	cache.hash = make(map[int32](*Node))
	(*cache.head).right = cache.tail
	(*cache.tail).left = cache.head
	cache.capacity = 0
}