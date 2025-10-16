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
