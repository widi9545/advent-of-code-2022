package test
import "fmt"

func Hello(name string) int{

	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}


func test() {
    fmt.Println("Hello, welcome to the program!")
	returnMessage := Hello("William")
	var name string = returnMessage
	fmt.Println(name)

}