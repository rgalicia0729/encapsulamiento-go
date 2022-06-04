package main

import (
	"fmt"

	"github.com/rgalicia0729/encapsulamiento-go/study"
)

func main() {
	goCourse := study.NewCourse("Curso de go desde cero")
	goCourse.SetPrice(12.0)

	classes := []string{
		"Introducción",
		"Generalidades del lenguaje",
		"Declaración de variables",
	}

	goCourse.SetClasses(classes)

	fmt.Println(goCourse.Name())
	fmt.Println(goCourse.Price())
	fmt.Println(goCourse.IsFree())
	fmt.Println(goCourse.Classes())
}
