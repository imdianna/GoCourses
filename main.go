package main
import(
	"fmt"
)
//спільний інтерфейс для всіх тварин
type animal interface { //static type
	getFoodAmount() float32
	getType() string
	getWeight() float32
	getName() string
}

type dogge struct {
	weight float32
	name string
}

func (d dogge) getFoodAmount() float32 {
	return d.weight * 2.0
}

func (d dogge) getType() string {
	return "dog"
}

func (d dogge) getWeight() float32 {
	return d.weight
}

func (d dogge) getName() string {
	return d.name
}

type cat struct {
	weight float32
	name string
}

func (c cat) getFoodAmount() float32 {
	return c.weight * 7.0
}

func (c cat) getType() string {
	return "cat"
}
func (c cat) getName() string {
	return c.name
}
func (c cat) getWeight() float32 {
	return c.weight
}
type cow struct {
	weight float32
	name string
}

func (c cow) getFoodAmount() float32 {
	return c.weight * 25.0
}

func (c cow) getType() string {
	return "cow"
}
func (c cow) getName() string {
	return c.name
}
func(c cow) getWeight() float32 {
 	return c.weight
}

func getFoodAmountForAnimal(a animal) {
	fmt.Printf("%v is a %v. Weights %v kg. Consumes %v kg of food per month\n", 
		a.getName(), a.getType(), a.getWeight(), a.getFoodAmount())
	
}
func totalAmountKg(animals []animal) {
	var total_amount float32
	for _,a := range animals {
		getFoodAmountForAnimal(a)


	
		
		total_amount += a.getFoodAmount()
	}
	fmt.Printf("total_amount of food for all farm animals - %v kg\n", total_amount )
}


func main() {

	littleFarm := [] animal {dogge{weight:23, name:"Uki"}, dogge{weight:30, name:"Bobosu"}, cat{weight:10, name:"Murka"},
	cow{weight:125, name:"Moo"} }
	totalAmountKg(littleFarm)
}
