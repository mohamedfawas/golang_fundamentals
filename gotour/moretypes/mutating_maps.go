package main

import "fmt"

func main() {
	fmt.Println("mutate maps")
	players := make(map[int]string)
	players[7] = "Vinicius Jr"
	players[11] = "Rodrygo Goes"
	players[5] = "Jude Bellingham"
	players[21] = "Brahim Diaz"
	delete(players, 21) // Delete the key-value pair of Brahm Diaz
	fmt.Println(players)

	v, ok := players[5]
	fmt.Println("The value:", v, "Present?", ok)
}
