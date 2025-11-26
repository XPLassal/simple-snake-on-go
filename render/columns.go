package render

import "fmt"

func GetNumberOfColumns(numbersOfColumns *int) {
	for {
		fmt.Print(Bold + "Write number of columns: ")
		fmt.Scan(numbersOfColumns)
		if *numbersOfColumns%2 == 0 || *numbersOfColumns < 2 {
			fmt.Println(Red + "Error: number of columns must be not even and greater than 2" + Reset + Bold)
		} else {
			break
		}
	}
	fmt.Print(Reset)
}
