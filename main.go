package main

import (
	"fmt"
	"github.com/mattrajca/GoEV3/Button"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	goPath := os.Getenv("GOPATH")

	if goPath == "" {
		log.Fatal("Cannot find GOPATH")
	}

	files, err := ioutil.ReadDir(fmt.Sprintf("%s/bin", goPath))

	if err != nil {
		log.Fatal("Cannot find bin")
	}

	selection := int(0)

	for {
		for i, f := range files {
			checked := " "

			if i == selection {
				checked = "X"
			}

			fmt.Printf("[%s] %s\n", checked, f.Name())
		}

		fmt.Print("\n")

		button := Button.WaitAny()

		if button == Button.Down {
			selection++
		} else if button == Button.Up {
			selection--
		} else if button == Button.Enter {
			c1 := exec.Command(fmt.Sprintf("%s/bin/%s", goPath, files[selection].Name()))
			_ = c1.Run()
		}
	}
}
