package main

import (
	"bufio"
	"fmt"
	"os"
)

type Scope = map[string]string

func NewScope() Scope {
	scopeElement := make(Scope)
	return scopeElement
}

func main() {
	var sliceScopes []Scope
	var fileName string

	fmt.Print("Name file: ")
	if _, err := fmt.Scan(&fileName); err != nil {
		fmt.Println("Error fmt.Scan - ", err)
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error open file - %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "{" {
			sliceScopes = append(sliceScopes, NewScope())
		} else if line == "}" {
			sliceScopes = sliceScopes[:len(sliceScopes)-1]
		} else if line == "ShowVar;" {
			showVariables := make(map[string]string)

			for _, element := range sliceScopes {
				for key, value := range element {
					showVariables[key] = value
				}
			}

			fmt.Println("ShowVar: ", showVariables)
		} else {
			variableName := ""
			variableValue := ""

			eqBool := false
			for _, char := range line {
				if char == '=' {
					eqBool = true
				}

				if !eqBool {
					variableName += string(char)
				} else {
					variableValue += string(char)
				}
			}

			scope := sliceScopes[len(sliceScopes)-1]
			scope[variableName] = variableValue
		}
	}

}
