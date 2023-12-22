package main

import(
	"fmt"
)

type Node struct {
	key int32
	val int32
	left *Node
	right *Node
}

type Hash map[int32]*Node

type Cache struct{
	head *Node
	tail *Node
	hash Hash
	capacity int32
}

func buildCache(size int32) Cache{
	newCache := Cache{
		head: &Node{0,0,nil,nil},
		tail: &Node{0,0,nil,nil},
		hash: Hash{},
		capacity: size,
	}
	(*(newCache.head)).right = newCache.tail
	(*(newCache.tail)).left = newCache.head
	return newCache
}


func main(){
	fmt.Println("LRU-cache")
	fmt.Println("---------------------------------")
	fmt.Print("Enter the capacity of the LRU cache : ")
	var capacity int32
	fmt.Scan(&capacity)	

	cache := buildCache(capacity)

	// while loop
	for v:=0;v<1;{

		var choice uint8 = 0
		fmt.Print(`Enter 
		1. Set a value 
		2. Get a value
		3. Print current state of cache
		4. Clear cache
		5. Quit - 0
		choice : `)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			cache.Set()
		case 2:
			cache.Get()
		case 3:
			cache.PrintState()
		case 4:
			cache.ClearCache()
		default:
			break
		}
		fmt.Printf("\n-----------------------------------------\n\n")	
	}
}





