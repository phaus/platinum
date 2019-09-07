package converters

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/phaus/platinum/utils"
)

const (
	startMarker = "```plantuml"
	endMarker   = "```"
)

//PlantumlConverter - a basic PlantUml Converter
type PlantumlConverter struct {
	serverURL string
}

//NewPlantumlConverter - create a new converter
func NewPlantumlConverter(url string) *PlantumlConverter {
	return &PlantumlConverter{
		serverURL: url,
	}
}

//Convert - run the converter
func (converter *PlantumlConverter) Convert(input string) (string, error) {

	var uml, line string
	var inPlantUML bool

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line == startMarker {
			fmt.Println("plantuml start")
			inPlantUML = true
			uml = ""
			line = ""
		}
		if line == endMarker {
			fmt.Println("plantuml end")
			inPlantUML = false
			encoded, err := compress(uml)
			if err != nil {
				return "", err
			}
			fmt.Printf("uml:\n%s\n", encoded)
			line = ""
		}
		if inPlantUML {
			uml += line + "\n"
		} else {
			if line != endMarker && line != startMarker {
				fmt.Println(line)
			}
		}
	}
	return "", nil
}

func compress(content string) (string, error) {
	str := utils.EncodeP(content)
	return str, nil
}
