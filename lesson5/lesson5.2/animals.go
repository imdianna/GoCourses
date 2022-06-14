package main
import(

	"bufio"
	"fmt"
	"strings"
	"os"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter animal: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	if strings.Contains(text,"dog") {
		fmt.Println("dog has 4 paws")
	}else if strings.Contains(text,"chicken"){
		fmt.Println("chicken has 2 legs")
	}else if strings.Contains(text, "spider"){
		fmt.Println("spider has 6 legs")
	}else {
		fmt.Println("I don't know this animal")
	}



}