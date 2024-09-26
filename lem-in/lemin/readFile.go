package lemin

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadData(filePath string) (linksMap map[string][]string, start, end string, antsNum int, input string, err error) {

	linksMap = make(map[string][]string)

	file, err := os.Open(filePath)
	if err != nil {
		file, err = os.Open("./examples/"+ filePath) 
		if err != nil {
			return nil, "", "", 0, "", fmt.Errorf("can't open your input file")
		}		
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, "", "", 0, "", fmt.Errorf("can't read your input file")
	}

	input = string(fileBytes)
	inputSlice := strings.Split(input, "\n")

	for i, line := range inputSlice {

		if i == 0 {
			var atoierr error
			antsNum, atoierr = strconv.Atoi(line)
			if atoierr != nil || antsNum == 0 {
				return nil, "", "", 0, "", fmt.Errorf("error reading ants number from your input file")
			}
			continue
		}
		if line == "##start" {
			start, err = CheckStartorEnd("Start", i, inputSlice)
			if err != nil {
				return nil, "", "", 0, "", err
			}
			continue
		}
		if line == "##end" {
			end, err = CheckStartorEnd("End", i, inputSlice)
			if err != nil {
				return nil, "", "", 0, "", err
			}
			continue
		}
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "L") || strings.HasPrefix(line, " ") {
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			from := parts[0]
			to := parts[1]
			linksMap[from] = append(linksMap[from], to)
			linksMap[to] = append(linksMap[to], from) // Add the reverse link
		}
	}

	if start == end {
		return nil, "", "", 0, "", fmt.Errorf("wrong start/end room")
	}

	return linksMap, start, end, antsNum, input, nil
}

func CheckStartorEnd(or string, i int, inputSlice []string) (string, error) {
	if i == len(inputSlice)-1 {
		return "", fmt.Errorf("%s room is missing", or)
	}
	startORendRoom := strings.Fields(inputSlice[i+1])
	if len(startORendRoom) != 3 || startORendRoom == nil {
		return "", fmt.Errorf("%s room coordinates are not correctly formated", or)
	}
	startORend := startORendRoom[0]
	if strings.HasPrefix(startORend, "#") || strings.HasPrefix(startORend, "L") || strings.Contains(startORend, " ") {
		return "", fmt.Errorf("%s room is missing", or)
	}
	return startORend, nil
}
