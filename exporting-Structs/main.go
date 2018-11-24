/*********************
* Exporting a package
********************/

package main

//when importing use full path when using other files
import (
	"fmt"
	"strings"

	"github.com/alonzathompson/go-fundamentals/exporting-Structs/types"
)

func changePlayerName(p athletes.Player) {
	p.Name = "Led Zepplin"
}

func changePlayerName2(p *athletes.Player) {
	p.Name = "Led Zepplin"
	p.Country = strings.ToUpper(p.Country)
}

func main() {
	player1 := athletes.Player{Name: "Anderson Cooper", Sport: "MMA", Age: 43, Info: athletes.Info{Country: "Brazil", HairColor: "Black"}}

	fmt.Println("(1) player1: ", player1)

	changePlayerName(player1)
	fmt.Println("(1A) player1: ", player1)

	changePlayerName2(&player1)
	fmt.Println("(2) player1 with pointer: ", player1)

	// Notice even with the methods staying true to every thing exported must
	// be capitalized even the reciever functions
	fmt.Println("(3) using method form player struct: ", player1.ToLowerCase())
}
