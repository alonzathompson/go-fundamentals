/************
* Embedded Structs
************/

package main

import "fmt"

// struct that will be embedded
type generalInfo struct {
	hairColor, country string
}

// struct that is embedding the generalInfo struct
type player struct {
	name, sport string
	age         int
	info        generalInfo
}

func main() {
	var player1 player

	//Assignment to struct
	player1.name = "Wayne Gretksy"
	player1.sport = "hockey"
	player1.age = 58

	//embedded structs inside other structs are extended with dot notation
	player1.info.country = "USA"
	player1.info.hairColor = "brown"

	fmt.Println(player1)
	fmt.Printf("player name=%s, player sport=%s, player age=%d, player info-country=%s, player info-hairColor=%s\n", player1.name, player1.sport, player1.age, player1.info.country, player1.info.hairColor)

	player2 := player{
		name:  "Just Blaze",
		age:   43,
		sport: "Music",
		info: generalInfo{
			country:   "USA",
			hairColor: "Black",
		},
	}

	fmt.Println(player2)
	fmt.Printf("player name=%s, player sport=%s, player age=%d, player info-country=%s, player info-hairColor=%s\n", player2.name, player2.sport, player2.age, player2.info.country, player2.info.hairColor)

}
