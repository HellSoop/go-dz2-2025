package task1

import (
	"io"
	"strings"
	"bufio"
)


func AnimalFeeding(stream io.Reader) ([]string) {
	food, animals := parseFoodAndAnimalsFromStream(stream)
	result := []string{}

	if len(animals) == 0{
		return result
	}

	i := 0 // in a for loop, the increment of i may be different from 1 due to unicode symbols
	for _, c := range food {
		animals[i % len(animals)].Eat(c)
		i++
	}
	
	for _, animal := range animals {
		result = append(result, animal.name + " " + animal.WhatDidYouEat())
	}

	return result
}


func parseFoodAndAnimalsFromStream(stream io.Reader) (string, []Animal) {
	food := ""
	animals := []Animal{}
	
	reader := bufio.NewReader(stream)
	isFoodRow := true

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSuffix(row, "\n")
		
		if err != nil {  // current row is "end" 
			break
		}

		if isFoodRow {
			food = row
			isFoodRow = false
		
		} else if len(row) > 0 {  // so rows can be skipped
			animals = append(animals, createAnimal(row))
		}
	}

	return food, animals
}


func createAnimal(animal_string string) Animal {
	data := strings.Fields(animal_string)
	return Animal{strings.ToLower(data[0]), data[1], []rune{}}
}
