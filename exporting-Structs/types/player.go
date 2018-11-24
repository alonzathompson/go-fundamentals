/*********************
* Exporting a package
* Notice how all the fields are Capitalized to export this as a package
********************/

package athletes

import (
	"strings"
)

// Gets second level info
type Info struct {
	Country, HairColor string
}

// Gets first level info
type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

// Reciever function that takes in the player
func (p *Player) ToLowerCase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Name)
	p.Country = strings.ToLower(p.Country)
	p.HairColor = strings.ToLower(p.HairColor)

	return p
}
