package main

import (
	"fmt"
	"path"

	"github.com/phaus/platinum/converters"
	"github.com/phaus/platinum/utils"
)

func main() {

	input, err := utils.ReadFromFile(path.Join("testdata", "test1.md"), "")
	if err != nil {
		panic(err)
	}

	converter := converters.NewPlantumlConverter("localhost")
	output, err := converter.Convert(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n", output)
}
