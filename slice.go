package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main(){
	intSlice:=make([]int,3)
	
	for {
		
		fmt.Println("Please enter an int. Enter X to close program")
		var input string
		fmt.Scan(&input)
		
		if input =="X" || input=="x"{
			break
		}
		

	num, err := strconv.Atoi(input)

	if err!=nil{
		fmt.Println("Incorrect Entry:", err)
		continue
	}

	intSlice=append(intSlice,num)
	
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	}


	fmt.Println("Exited Program")
	fmt.Println("Final Output ", intSlice)


}