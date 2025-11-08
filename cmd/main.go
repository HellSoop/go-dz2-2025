package main

import (
	"fmt"
	"strings"
	"dz2/internal/task_1"
)

func main() {
	// default
	checkAnimalFeeding(`q--3wer4301-4+r_e**
dog d1
cat c1
cat c2
bird b1
end`)

	// just one more
	checkAnimalFeeding(`wwa?tt123123i++
cat meow
cat pur
dog bark
bird чирик
end`)

	// case test 
	checkAnimalFeeding(`q--3wer4301-4+r_e**
Dog d1
cAt c1
caT c2
BIRD b1
end`)

	// food is empty
	checkAnimalFeeding(`
cat Мазик
dog Виталя
bird Жора
end`)

	// unicode test
	checkAnimalFeeding(`a1😎👀✌✌✔🎉
dog first
cat second
bird Биба
bird Боба
end`)

	// no animals 
	checkAnimalFeeding(`asd124po23089sdgjkl2-09sdfg12123123fg
end`)
	
	// skip second row test
	checkAnimalFeeding(`

dog Бобик
cat Жорик
bird Попка-дурак
end`)
}


func checkAnimalFeeding(input string) {
	fmt.Printf("Данные:\n%s\n\n", input)
	stream := strings.NewReader(input)
	results := task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Printf("\n\n\n")
}
