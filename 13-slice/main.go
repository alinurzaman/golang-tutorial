package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Changed"
	// fmt.Println(slice1)

	// slice1[0] = "Changed"
	// fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	// EXPLANATION ABOUT APPEND:
	// if the capacity is full then make new array and doesn't change the origin array
	// but if the capacity isn't full then replace the next element
	var slice3 = append(slice2, "New Month")
	fmt.Println(slice3)

	slice3[1] = "Newer Month"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// 'make' for create new slice make([]{type}, {length}, {capacity})
	newSlice := make([]string, 2, 5)

	newSlice[0] = "One"
	newSlice[1] = "Two"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice, make sure the length and capacity are sam with the origin slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// the difference is at [...] => array ; [] => slice
	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
