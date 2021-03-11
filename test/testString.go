package main

import (
	"fmt"
)

func main() {
	//var stringVar = "application2_database.application_game"
	//
	//s := strings.Split(stringVar, ".")
	//
	//fmt.Println("s: ", s)

	portLabels := map[string]string{"tag_binding": "application_web", "application2": "database"}

	//fmt.Println(portLabels)
	a := fmt.Sprintf("%q", portLabels)
	fmt.Println(a)
}
