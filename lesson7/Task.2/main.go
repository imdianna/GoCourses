package main

import "fmt"
import "strings"


func main() {

	input:= "1 1"
	//повернути найбільше та найменше число
	var result string

    newstring := strings.Split(input, " ")

	if len(newstring) == 1 {
        result = newstring [0]
	} else if len(newstring) > 1 {
		  var min int
		  var max int
		  var curr int

		fmt.Sscan(newstring[0], &min) //передаємо значення "1" -> 1
		max = min

//forrange_юзати ключі та значенняб ключ _, значення і
		for _, val := range newstring {
			fmt.Sscan(val, &curr)
			if min > curr {
				min = curr
			} else if max < curr {
				max = curr
			}
		}

		result = fmt.Sprintf("%d %d", max, min)
	}

	fmt.Println("результат", result)
}
	