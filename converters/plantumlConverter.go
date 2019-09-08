package converters

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/phaus/platinum/utils"
)

const (
	defaultServer    = "http://www.plantuml.com/plantuml"
	defaultFolder    = "build"
	defaultImagePath = "images"
	startMarker      = "```plantuml"
	endMarker        = "```"
	// PngKind type server/png
	PngKind = "png"
	// SvgKind type server/svg
	SvgKind = "svg"
	// TxtKind type server/txt
	TxtKind = "txt"
)

//PlantumlConverter - a basic PlantUml Converter
type PlantumlConverter struct {
	serverURL    string
	outputFolder string
	imagePath    string
}

//NewPlantumlConverter - create a new converter
func NewPlantumlConverter(url string, outputFolder string, imagePath string) *PlantumlConverter {
	var imgFolder string
	if url == "" {
		url = defaultServer
	}
	if outputFolder == "" {
		outputFolder = defaultFolder
	}
	if imagePath == "" {
		imagePath = defaultImagePath
	}
	imgFolder = path.Join(outputFolder, "images")
	log.Printf("Convert images via %s into %s", url, imgFolder)
	if _, err := os.Stat(imgFolder); os.IsNotExist(err) {
		os.MkdirAll(imgFolder, 0777)
	}
	return &PlantumlConverter{
		serverURL:    url,
		outputFolder: outputFolder,
		imagePath:    imagePath,
	}
}

//Convert - run the converter
func (converter *PlantumlConverter) Convert(inputFile string, kind string) (string, error) {

	var uml, line string
	var inPlantUML bool

	dir, filename := converter.outputFile(inputFile)
	input, err := utils.ReadFromFile(inputFile, "")
	if err != nil {
		panic(err)
	}
	file, err := os.Create(path.Join(dir, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedWriter := bufio.NewWriter(file)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line == startMarker {
			inPlantUML = true
			uml = ""
			line = ""
		}
		if line == endMarker {
			inPlantUML = false
			encoded, err := compress(uml)
			if err != nil {
				return "", err
			}
			content, err := converter.call(encoded, kind, path.Join(dir, "images"))
			if err != nil {
				return "", err
			}
			line = content
		}
		if inPlantUML {
			uml += line + "\n"
		} else {
			if line != endMarker && line != startMarker {
				_, err = bufferedWriter.Write(
					[]byte(
						fmt.Sprintf("%s\n", line)))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	bufferedWriter.Flush()
	return "", nil
}

func compress(content string) (string, error) {
	str := utils.EncodeP(content)
	return str, nil
}

func (converter *PlantumlConverter) outputFile(inputFile string) (string, string) {
	out := ""
	dir, file := path.Split(inputFile)
	parts := strings.Split(dir, string(os.PathSeparator))
	parts[0] = converter.outputFolder
	for _, part := range parts {
		out = path.Join(out, part)
	}
	if _, err := os.Stat(path.Join(out, "images")); os.IsNotExist(err) {
		os.MkdirAll(path.Join(out, "images"), 0777)
	}
	return out, file
}

func (converter *PlantumlConverter) call(encoded string, kind string, imageFolder string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", converter.serverURL, kind, encoded)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	switch kind {
	case TxtKind:
		return fmt.Sprintf("<pre>%s</pre>", string(body)), nil
	case SvgKind:
		fallthrough
	case PngKind:
		h := hash(body)
		dst := path.Join(imageFolder, fmt.Sprintf("%s.%s", h, kind))

		file, err := os.Create(dst)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		bufferedWriter := bufio.NewWriter(file)
		_, err = bufferedWriter.Write(body)
		if err != nil {
			log.Fatal(err)
		}
		bufferedWriter.Flush()
		return fmt.Sprintf("![%s](%s/%s.%s)", h, converter.imagePath, h, kind), nil
	}
	return "", nil
}

func hash(content []byte) string {
	h := sha1.New()
	h.Write(content)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
