package main

import (
)

func main() {
	//pokemon := Pokemon{Id: 1, Name: "チコリータ", H: 45, A: 49, B: 65, C: 49, D: 65, S: 45, Sum: 318}
	//insert(pokemon);
	//printAll();

	list := scraping()
	printList(list)
}
