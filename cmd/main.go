package main

import (
	"fmt"
	"strings"
	"dz2/internal/task_1"
)

func main() {
	// default
	input := `q--3wer4301-4+r_e**
dog d1
cat c1
cat c2
bird b1
end`

	fmt.Printf("Данные:\n%s\n\n", input)
	stream := strings.NewReader(input)
	results := task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Printf("\n\n\n")

	// just one more
	input = `wwa?tt123123i++
cat meow
cat pur
dog bark
bird чирик
end`

	fmt.Printf("Данные:\n%s\n\n", input)
	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Printf("\n\n\n")

	// case test 
	input = `q--3wer4301-4+r_e**
Dog d1
cAt c1
caT c2
BIRD b1
end`

	fmt.Printf("Данные:\n%s\n\n", input)
	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Printf("\n\n\n")

	// food is empty
	input = `
cat Мазик
dog Виталя
bird Жора
end`

	fmt.Printf("Данные:\n%s\n\n", input)
	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Printf("\n\n\n")

	// unicode test
	input = `a1😎👀✌✌✔🎉
dog first
cat second
bird Биба
bird Боба
end`

	fmt.Printf("Данные:\n%s\n\n", input)
	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Printf("\n\n\n")

	
	// skip first row test
	input = `

dog Бобик
cat Жорик
bird Попка-дурак
end`

	fmt.Printf("Данные:\n%s\n\n", input)
	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}
}