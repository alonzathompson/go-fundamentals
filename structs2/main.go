/********
* Structs Two
********/
package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {
	player1 := player{"Lebron James", "Basketball", 36}

	fmt.Println("Player 1=", player1)
	fmt.Printf("(1) name=%s age=%d", player1.name, player1.age)

	player3 := new(player)
	player3.name = "Tiger Woods"
	player3.sport = "Golf"
	player3.age = 40

	fmt.Printf("(3) player3 name=%s player3 age=%d\n", player3.name, player3.age)

	//anonymous struct tied to player4 immediately initialized
	player4 := struct {
		name, job string
		age       int
	}{"John", "Developer", 30}

	fmt.Printf("player4 name=%s player4 job=%s player4 age=%d\n", player4.name, player4.job, player4.age)
}
