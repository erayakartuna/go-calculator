package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var total = 0

func main() {

	fmt.Printf("Write a number\n")
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	text = strings.Replace(text,"\n","",-1)
	number,err := strconv.Atoi(text)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("what action do you want to do (+,-,/,*) \n")
	operation, _ := reader.ReadString('\n')
	operation = strings.Replace(operation,"\n","",-1)

	proccess(operation,number)

	fmt.Printf("Result = " + strconv.Itoa(total) + "\n")
	fmt.Printf("do you want to continue ? Y/N \n")

	operation, _ = reader.ReadString('\n')
	operation = strings.Replace(operation,"\n","",-1)

	//loop
	if operation == "Y"{
		main()
		return;
	}

	fmt.Printf("Finished \n")
}


func proccess(operation string,number int) {
	switch operation {
		case "+":
			sum(number)
		case "-":
			subtraction(number)
		case "*":
			multiplication(number)
		case "/":
			division(number)
	}
}

func sum(input int){
	total += input
}

func subtraction(input int){
	total -= input
}

func multiplication(input int){
	total *= input
}

func division(input int){
	total /= input
}