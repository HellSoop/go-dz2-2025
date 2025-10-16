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
	// so second row may be skipped
	if len(rows[1]) == 0 {
		rows = append(rows[:1], rows[2:]...)
	}
	
	food := rows[0]
	rows = rows[1:len(rows) - 1]
	
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