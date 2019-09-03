package converters

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
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
		}
		if line == endMarker {
			fmt.Println("plantuml end")
			inPlantUML = false
			encoded, err := compress(uml)
			if err != nil {
				return "", err
			}
			fmt.Printf("uml:\n%s\n", encoded)
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
	//	return "", fmt.Errorf("error converting %d bytes", len(input))
}

func compress(content string) (string, error) {
	var b bytes.Buffer
	const dict = `>>` + `--` + `==` + ` ` + `<<`
	// Compress the data using the specially crafted dictionary.
	zw, err := flate.NewWriterDict(&b, flate.BestCompression, []byte(dict))
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(zw, strings.NewReader(content)); err != nil {
		return "", err
	}
	if err := zw.Close(); err != nil {
		return "", err
	}

	str := base64.URLEncoding.EncodeToString(b.Bytes())
	return str, nil
}
