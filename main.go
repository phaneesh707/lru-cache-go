package main

import(
	"fmt"
)

type node struct {
	key int32
	val int32
	left *node
	right *node
}

//capacity of the cache
var capacity uint32
// declaring head & tail for LL
var head *node = &node{
	key: 0,
	val: 0,
	left: nil,
	right: nil,
}

var tail *node = &node{
	key: 0,
	val: 0,
	left: nil,
	right:nil,
}

//define map key -> ptr to node
var myMap = map[int32](*node){}

// define the count
var cacheSize uint32 = 0

func main(){
	fmt.Println("LRU-cache")
	fmt.Println("---------------------------------")
	fmt.Print("Enter the capacity of the LRU cache : ")
	fmt.Scan(&capacity)	

	

	// init the LL
	(*head).right = tail
	(*tail).left = head

	// while loop
	for v:=0;v<1;{

		var choice uint8 = 0
		fmt.Print(`Enter 
		1. Set a value 
		2. Get a value
		3. Print current state of cache
		4. Clear cache
		any number to exit 
		choice : `)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Set()
		case 2:
			Get()
		case 3:
			PrintState()
		case 4:
			ClearCache()
		default:
			break
		}
		fmt.Printf("\n-----------------------------------------\n\n")	
	}
}





