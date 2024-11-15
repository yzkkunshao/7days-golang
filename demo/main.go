package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	map1[1] = "Monday"
	map1[2] = "Tuesday"
	map1[3] = "Wendesday"
	map1[4] = "Tursday"
	map1[5] = "Friday"
	map1[6] = "Saturday"
	map1[0] = "Sunday"
	for key := range map1 {
		if value, ok := map1[key]; ok {
			if value == "Tuesday" {
				fmt.Println("has Tuesday")
			}
			fmt.Println("day", value)
		}
	}
}
