package task1

import (
	"io"
	"strings"
)

func parseStream(stream io.Reader) (string, []Animal) {
	buffer := make([]byte, 32)
	full_data := []byte{}
	n := 0
	var err error = nil
	
	for err == nil {
		n, err = stream.Read(buffer)
		full_data = append(full_data, buffer[:n]...)
	}

	rows := strings.Split(string(full_data), "\n")
	food := ""
	
	// so first row may be skipped (test "Пустая строка корма")
	if len(rows[0]) == 0 && len(strings.Fields(rows[1])) != 3 {
		food = rows[1]
		rows = rows[2:len(rows) - 1]
	} else {
		food = rows[0]
		rows = rows[1:len(rows) - 1]
	}
	
	var current_data []string
	animals := []Animal{}

	for _, animal_string := range rows {
		current_data = strings.Fields(animal_string)
		animals = append(animals, Animal{strings.ToLower(current_data[0]), current_data[1], []rune{}})
	}

	return food, animals
}

func AnimalFeeding(stream io.Reader) ([]string) {
	food, animals := parseStream(stream)
	i := 0 // in a for loop, the increment of i may be different from 1 due to unicode symbols
	
	for _, c := range food {
		animals[i % len(animals)].Eat(c)
		i++
	}

	result := []string{}
	for _, animal := range animals {
		result = append(result, animal.name + " " + animal.WhatDidYouEat())
	}

	return result
}