package main

import (
	"fmt"
	"os"

	"github.com/phaus/platinum/converters"
)

func main() {

	var inputFile string
	var outputFolder string
	var outputKind string
	var imagePath string
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 6 || len(argsWithoutProg)%2 != 0 {
		fmt.Println("start with: platinum --kind [png|txt|svg] --input /path/content.md --output build/ (--imagepath path-in-markdown)")
	} else {
		for i := 0; i < len(argsWithoutProg); i++ {
			if argsWithoutProg[i] == "--input" {
				inputFile = argsWithoutProg[i+1]
			}
			if argsWithoutProg[i] == "--output" {
				outputFolder = argsWithoutProg[i+1]
			}
			if argsWithoutProg[i] == "--kind" {
				outputKind = argsWithoutProg[i+1]
			}
			if argsWithoutProg[i] == "--imagepath" {
				imagePath = argsWithoutProg[i+1]
			}
		}
	}

	if inputFile != "" && outputFolder != "" {
		converter := converters.NewPlantumlConverter("http://www.plantuml.com/plantuml", outputFolder, imagePath)
		output, err := converter.Convert(inputFile, outputKind)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n%s\n", output)
	}
}
