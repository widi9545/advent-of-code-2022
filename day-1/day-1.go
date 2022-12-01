package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
	"sort"
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
	var sum int
	sum = 0
	caloriesCarriedByEachElf := []int{}



	for scanner.Scan(){

		var inputString string = scanner.Text()
		i, _ := strconv.Atoi(inputString)
		sum = sum + i
		
		if i == 0 {
			caloriesCarriedByEachElf = append(caloriesCarriedByEachElf, sum)
			sum = 0
		}
		

	}
	sort.Ints(caloriesCarriedByEachElf)
	firstElf := caloriesCarriedByEachElf[len(caloriesCarriedByEachElf)-1]
	secondElf := caloriesCarriedByEachElf[len(caloriesCarriedByEachElf)-2]
	thirdElf := caloriesCarriedByEachElf[len(caloriesCarriedByEachElf)-3]
	fmt.Println(firstElf+secondElf+thirdElf)




	


	


}