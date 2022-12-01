package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)


//strings are immutable - probably because they are actually arrays of chars!
//arrays need to be initialized with how big they need to be, but slices can be initialized as 'empty', can go from there
//can make dynamically sized slices with make
//need to use all imported/declared packages and variables
//_ is used as a placeholder


func main() {
    fmt.Println("Hello, welcome to day-1!")
	content, err:= os.Open("input.txt")

	if err != nil{
		log.Fatal(err)
	}

	defer content.Close()
	scanner := bufio.NewScanner(content)
	var calorieCounter, sum, elf int
	calorieCounter = 0
	index := 1



	for scanner.Scan(){

		var inputString string = scanner.Text()
		i, _ := strconv.Atoi(inputString)
		sum = sum + i

		if i == 0 {
			index = index + 1
			if sum >= calorieCounter{
				calorieCounter = sum
				elf = index
			}
			sum = 0
			fmt.Println(calorieCounter,elf)
		}

	}
	


	


}