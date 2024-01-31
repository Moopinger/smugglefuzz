/*
2024 Moopinger
*/

package lib

import (
	"fmt"
	"strings"
)

func OutputParser(name string, response string, colourDisable bool, filter string) string {

	var filterReturn string

	if filter != "" && !strings.Contains(response, filter) {
		filterReturn = "\r"
	} else {
		filterReturn = "\n"
	}

	if len(response) < 3 {
		return "ERROR: Bad response argument passed"
	}

	if colourDisable == true {
		return fmt.Sprintf("%-30s [%s]%s", name, response, filterReturn) //no colour
	}

	if response[:3] == "SUC" {
		return fmt.Sprintf("\033[32m%-30v [%s]\033[0m%40s", name, response, filterReturn) // successful confirmation - green
	} else if response[:3] == "RST" {
		return fmt.Sprintf("\033[33m%-30v [%s]\033[0m%40s", name, response, filterReturn) // RST_STREAM yellow

	} else if response[:3] == "[+]" {
		return fmt.Sprintf("\033[32m%-30v [%s]\033[0m%40s", name, response, filterReturn) // success - green
	} else if response[:3] == "*TI" {

		return fmt.Sprintf("\033[36m%-30v [%s]\033[0m%40s", name, response, filterReturn) // TIMEOUT - cyan
	} else if response[:3] == "*Se" {

		return fmt.Sprintf("\033[35m%-30s [%s]\033[0m\n", name, response) //"sending confirmation message" - pink no filter
	} else if response[:3] == "GOA" {

		return fmt.Sprintf("\033[35m%-30s [%s]\033[0m%40s", name, response, filterReturn) // GOAWAY - pink
	} else {
		return fmt.Sprintf("\033[31m%-30s [%s]\033[0m\n", name, response) // unknown so red
	}

}
