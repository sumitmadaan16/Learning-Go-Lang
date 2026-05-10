package decorator

import "fmt"

type Coffee interface {
	Cost() int
	Description() string
}

// concrete type
type Espresso struct{}

func (e Espresso) Cost() int {
	return 250
}
func (e Espresso) Description() string {
	return "Espresso is a concentrated form of coffee produced by forcing hot water under high pressure through finely ground coffee beans. " +
		"Originating in Italy, espresso has become one of the most popular coffee-brewing methods worldwide."
}

// decorators

type Milk struct{ espreso Espresso }

func (m Milk) Cost() int {
	return m.espreso.Cost() + 70
}
func (m Milk) Description() string {
	return "Adding milk to Espresso becomes Latte. " + "A latte is a smooth coffee drink made by combining a shot of espresso with steamed milk, resulting in a creamy texture and mellow flavor. Originating in Italy, it has become a staple in cafés worldwide, often enjoyed plain or with flavored syrups."
}

func DecoratorPattern() {
	espresso := Espresso{}
	fmt.Println(espresso.Description(), " : ", espresso.Cost())

	latte := Milk{espresso}
	fmt.Println(latte.Description(), " : ", latte.Cost())

}
