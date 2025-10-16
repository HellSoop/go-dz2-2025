package main

import (
	"fmt"
	"strings"
	"dz2/internal/task_1"
)

func main() {
	input := `q--3wer4301-4+r_e**
dog d1
cat c1
cat c2
bird b1
end`

	stream := strings.NewReader(input)
	results := task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}
}