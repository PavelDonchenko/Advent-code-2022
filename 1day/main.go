package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ParseElves(fileName string) ([]int, error) {
	inputSlice := make([]int, 0)
	file, err := os.Open(fileName)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File is empty %v\n", err)
	}

	scanner := bufio.NewScanner(file)
	calorieSum := 0
	for scanner.Scan() {
		thisNumber, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("Number in loop: %d\n", thisNumber)
		if thisNumber != 0 {
			//fmt.Printf("not a 0!")
			calorieSum += thisNumber
			//fmt.Printf("calorieSum %d\n", calorieSum)
		} else {
			inputSlice = append(inputSlice, calorieSum)
			calorieSum = 0
		}
	}
	err = file.Close()
	if err != nil {
		return inputSlice, err
	}
	return inputSlice, nil
}

func main() {
	elfCalories, err := ParseElves("input.txt")
	if err != nil {
		fmt.Printf("Error%v\n", err)
	}
	//fmt.Printf("Elf Calories: %v\n", elfCalories)

	sort.Ints(elfCalories)
	//var maxSlice []int
	//temp := 0
	//
	//for _, v := range elfCalories {
	//	if v > temp {
	//		temp = v
	//		maxSlice = append(maxSlice, v)
	//	}
	//}
	maxCal := elfCalories[len(elfCalories)-1:]

	var max3sum int

	for _, v := range elfCalories[len(elfCalories)-3:] {
		max3sum += v
	}

	fmt.Printf("Max Calories: %d\n", maxCal)
	fmt.Printf("Max 3 elves Calories: %d\n", max3sum)
}
