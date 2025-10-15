package task1

import "unicode"

type Animal struct {
	type_ string
	name string
	ate []rune
}

func (a Animal) CanEat(c rune) bool {
	switch a.type_ {
		case "cat":
			 if unicode.IsDigit(c) {
				return true
			 }
		
		case "dog":
			if unicode.IsLetter(c) {
				return true
			}

		case "bird":
			if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
				return true
			}
	}
	
	return false
}

func (a *Animal) Eat(c rune) {
	if a.CanEat(c) {
		a.ate = append(a.ate, c)
	}
}

func (a Animal) WhatDidYouEat() string {
	return string(a.ate)
}