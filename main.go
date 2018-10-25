package main

import "fmt"
import "github.com/shalk/trello-json/trello"

func main() {
	// 1.json export from trello
	info := trello.Read("C:\\Users\\me\\Desktop\\1.json")
	for key := range info {
		fmt.Printf("List name = %s \n", info[key][0])
		for _, v := range info[key][1:] {
			fmt.Println(" " + v)
		}
	}
}
