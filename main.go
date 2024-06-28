package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

const APP_NAME string = "Hearthstone.exe"
const SEARCH_COL_NAME = "\"PID\""

func findStringIndex(slice []string, searchVal string) (int, error) {
	var result int = -1
	var err error

	for i, val := range slice {
		if val == searchVal {
			result = i
			break
		}
	}

	if result == -1 {
		err = errors.New("cannot find value in the given slice")
	}

	// consider splitting result with error
	return result, err
}

func getProcessPidByName(imageName string) (string, error) {
	cmd := exec.Command("cmd", "/C", "tasklist", "/FI", fmt.Sprint("ImageName eq ", imageName), "/FI", "Status eq Running", "/FO", "CSV")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	stringsArr := strings.Split(string(output), "\n")

	// consider looking for many search rows
	headingRowArr := strings.Split(stringsArr[0], ",")
	searchRowLine := strings.Split(stringsArr[1], ",")

	colIndex, err := findStringIndex(headingRowArr, SEARCH_COL_NAME)

	if err != nil {
		return "", err
	}

	return searchRowLine[colIndex], nil
}

func main() {
	pid, err := getProcessPidByName(APP_NAME)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pid)
}
