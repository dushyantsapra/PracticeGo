package main

//In GO errors are value which is returned from the method.

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error %v, Error Code: %v", c.Message, c.Code)
}

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return float64(0), errors.New("Divide by zero error")
	} else {
		return x / y, nil
	}
}

func DivideWithCustomError(x, y float64) (float64, error) {
	if y == 0 {
		return float64(0), CustomError{10, "Can't Divide By Zero"}
	} else {
		return x / y, nil
	}
}

func main() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Error Reading File ", err)
	} else {
		fmt.Println("File Data is ", string(data))
	}

	value, divErr := Divide(10, 0)
	if divErr != nil {
		fmt.Println("Error Dividing ", divErr)
	} else {
		fmt.Println("Value is ", value)
	}

	value1, divError := DivideWithCustomError(1, 0)
	if divErr != nil {
		fmt.Println(divError)
		log.Fatal(divErr)
	} else {
		fmt.Println("value1 is ", value1)
	}

	fmt.Println("Done")
}
