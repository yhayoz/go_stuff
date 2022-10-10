package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Declare some default VARs

	// Some cli arguments
	var structure = flag.String("s", "simple", "Structure simple or complex")
	var hlq = flag.String("h", "", "HLQ")
	var app = flag.String("n", "", "App Name")
	var stages = flag.String("e", "", "Stages to create")
	flag.Parse()

	var default_stages = "dev,test,prod"
	var default_structure = "simple"

	if *structure == "" {
		structure = &default_structure
	}

	if *stages == "" {
		stages = &default_stages
	}
	var arr = strings.Split(*stages, ",")
	fmt.Println(`We Create following Repo:`, "\n",
		`Structure:`, *structure, "\n",
		`HLQ:`, *hlq, "\n",
		`App Name:`, *app, "\n",
		`Stages:`, arr)

	create_structure(arr, hlq, app, structure)
}

func create_structure(arr []string, hlq *string, app *string, structure *string) {
	if *structure == "simple" {
		// First create base Directorie
		os.MkdirAll(fmt.Sprintf("../deploy/stages/base/%s", *app), os.ModePerm)

		for i := 0; i < len(arr); i++ {
			os.MkdirAll(fmt.Sprintf("../deploy/stages/%s/%s", arr[i], *app), os.ModePerm)
		}
	}

	if *structure == "complex" {
		// First create all Directories
		os.MkdirAll(fmt.Sprintf("../deploy/base/%s", *app), os.ModePerm)
	}

}
