package helper

import "fmt"

const (
	ONE_GB = 1073741824
	ONE_TB = 1099511627776
)

func calculate(num int, choice int) int {
	switch choice {
	case 1:
		return ONE_GB * num
	case 2:
		return ONE_TB * num
	default:
		return 0
	}
}

func Start() {
	fmt.Println("Select the volume unit for below options: ")
	fmt.Println("1. GB \n2. TB")
	fmt.Print("Enter your choice: ")
	var choice, volume int
	fmt.Scan(&choice)
	fmt.Print("Enter volume: ")
	fmt.Scan(&volume)
	fmt.Println(calculate(volume, choice))
}
