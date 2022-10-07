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
		`Stages:`, arr)
	create_structure(app, arr, hlq, structure)
}

func contains(arr []string, str string) bool {

	for _, v := range arr {
		if v == str {

			fmt.Println("Exists")
			return true
		}
	}

	fmt.Println("does not exist")
	return true
}

func create_structure(app string, arr []string, hlq string, structure string) {
	contains(arr, "test")
	if structure == "simple" {
		// First create base Directorie
		os.MkdirAll(fmt.Sprintf("../deploy/stages/base/%s", app), os.ModePerm)

		for i := 0; i < len(arr); i++ {
			os.MkdirAll(fmt.Sprintf("../deploy/stages/%s", arr[i]), os.ModePerm)
		}
	}

	if structure == "complex" {
		// First create all Directories
		fmt.Println("hello", app)
		os.MkdirAll(fmt.Sprintf("../deploy/base/%s", app), os.ModePerm)
	}

}
