//Прибрати всі дублікати з слайсу int. отже, потрібно створити слайс інт1
//Приклад даних на вхід: [3, 4, 4, 3, 6, 3]
//видаляємо 3 по індексу 0
//видаляємо 4 по індексу 1
//видаляємо 3 по індексу 3
//Правильний результат: [3, 4, 6]
//Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.

package main
import "fmt"
func main() {

	arr := []int{3, 4, 4, 3, 6, 3}
	var result [] int 
	for i := 0; i < (len(arr)); i++ {
		f := true;
		for j:= 0; j < (len(result)); j++{
			if arr [i] == result [j] {
				f = false
				break;
			}
		}

		if f == true{
			result = append(result, arr[i])
		}

	}
	
	fmt.Println(result)
}
