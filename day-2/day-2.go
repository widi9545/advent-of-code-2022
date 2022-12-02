package main

import (
	"fmt"
	"log"
	"bufio"
	"os"

)

type ticTac struct{
	moveOne string
	moveTwo string
}

///if i have a string "hello", i can access it as individual characters by doing text[0:1] or by casting as string

func main() {
    fmt.Println("Hello, welcome to day-2!")
	content, err:= os.Open("input.txt")

	if err != nil{
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	var fileLines []string
	var moveSet []ticTac 


	for scanner.Scan(){
		var text string = scanner.Text()
		fileLines = append(fileLines, text)
	}

	for _, line := range fileLines{
		var moveStruct ticTac
		var firstMove = line[0:1]
		var secondMove = line[2:]
		moveStruct.moveOne = firstMove
		moveStruct.moveTwo = secondMove
		moveSet = append(moveSet, moveStruct)
	}

	gameBoard('C','Z')


}



//empty function valid for compiler? 

func gameBoard(moveOne rune, moveTwo rune) int {
	var result int
	var move int
	var sum int 
	switch moveOne{
	case 'A':
		if moveTwo == 'X'{
			move = 1
			result = 3
			sum = move + result
		}
		if moveTwo == 'Y'{
			move = 2
			result = 6
			sum = move + result
		}
		if moveTwo == 'Z'{
			move = 3
			result = 0
			sum = move + result
		}
	
	case 'B':
		if moveTwo == 'X'{
			move = 1
			result = 0
			sum = move + result
		}
		if moveTwo == 'Y'{
			move = 2
			result = 3
			sum = move + result
		}
		if moveTwo == 'Z'{
			move = 3
			result = 6
			sum = move + result
		}


	case 'C':
		if moveTwo == 'X'{
			move = 1
			result = 6
			sum = move + result
		}
		if moveTwo == 'Y'{
			move = 2
			result = 0
			sum = move + result
		}
		if moveTwo == 'Z'{
			move = 3
			result = 3
			sum = move + result
		}
	}
	fmt.Println(sum)
	return sum
}
