package main

import (
	"bufio"
	"log"
	"strings"
)

type item struct {
	Text     string
	Children []*item
}

func parse(data string) item {
	scanner := bufio.NewScanner(strings.NewReader(data))

	var rootItem item
	var currentLevelItem *item
	lastLevel := 0
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		level := getLevel(line)

		lineItem := item{
			Text:     strings.TrimLeft(line, " "),
			Children: make([]*item, 0),
		}

		if level == 0 {
			rootItem = lineItem
			currentLevelItem = &rootItem
			lastLevel = 1
		} else {
			if rootItem.Text == "" || level-lastLevel > 1 {
				log.Fatal("syntax error in line: ", lineNumber)
			}

			currentLevelItem = &rootItem
			// as we start from root level, start with level - 1
			for i := level - 1; i > 0; i-- {
				currentLevelItem = currentLevelItem.Children[len(currentLevelItem.Children)-1]
			}

			// log.Println("append at ", currentLevelItem)
			currentLevelItem.Children = append(currentLevelItem.Children, &lineItem)
			lastLevel = level
		}

	}

	return rootItem
}

func getLevel(line string) int {
	level := 0
	for {
		if strings.HasPrefix(line, "  ") {
			line = strings.Replace(line, "  ", "", 1)
			level++
		} else {
			break
		}
	}
	return level
}
