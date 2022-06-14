package main

import ("fmt")

const (
 	apple_price float32 = 5.99
 	pear_price float32 = 7
	budget float32 = 23
)

func main () {
	var apple_quantity float32 = 9
	var pear_quantity float32 = 8

	fmt.Println("----------------------------------")
	fmt.Println("1. Скільки грошей треба витратити, щоб купити", apple_quantity, "яблук та", pear_quantity, "груш?")
	result := (apple_price * apple_quantity) + (pear_price * pear_quantity)
	fmt.Printf("Відповідь: %.2f грн\n", result)

	fmt.Println("----------------------------------")
	fmt.Println("2.Скільки груш ми можемо купити?")
	result = (budget / pear_price)
	fmt.Printf("Відповідь: %.g груші\n", result)

	fmt.Println("----------------------------------")
	fmt.Println("3.Скільки яблук ми можемо купити?")
	result = (budget / apple_price)
	fmt.Printf("Відповідь: %.g яблука\n", result)

	fmt.Println("----------------------------------")
	apple_quantity = 2
	pear_quantity = 2
	fmt.Println("Чи ми можемо купити", pear_quantity, "груші та", apple_quantity, "яблука?")
	result = (pear_price *pear_quantity) + (apple_price * apple_quantity)
	if result > budget {
    	fmt.Println("Відповідь: ні")
	}else{
		fmt.Println("Відповідь: так")
	}
	
	
}