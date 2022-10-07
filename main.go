package main

import (
	"fmt"
	"os"
	"strings"
)

// set global Vars to use in diffrent functions

func main() {
	// Declare some default VARs
	var default_stages = "dev,test,prod"
	var default_structure = "simple"

	// Ask Some Questions for the Project
	fmt.Println("/(default structure:", default_structure, ") customer project structure type to create? [simple|complex]: ")
	var structure string
	fmt.Scanln(&structure)
	if structure == "" {
		structure = default_structure
	}

	fmt.Println("HLQ:")
	var hlq string
	fmt.Scanln(&hlq)

	fmt.Println("Name of App:")
	var app string
	fmt.Scanln(&app)

	fmt.Println("(default stages:", default_stages, ") stages:")
	var stages string
	fmt.Scanln(&stages)
	if stages == "" {
		stages = default_stages
	}
	arr := strings.Split(stages, ",")
	fmt.Println(arr)

	fmt.Println(`We Create following Repo:`, "\n",
		`Structure:`, structure, "\n",
		`HLQ:`, hlq, "\n",
		`App Name:`, app, "\n",
		`Stages:`, stages)
	create_structure(app)
}

func create_structure(app string) {
	// First create all Directories
	fmt.Println("hello", app)
	os.MkdirAll(fmt.Sprintf("../deploy/base/%s", app), os.ModePerm)

}
