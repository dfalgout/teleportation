package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dfalgout/teleportation/graph"
)

var (
	portalGraph = graph.NewGraph()
)

func main() {
	var (
		file = flag.String("file", "input.txt", "input file")
	)
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("Error Reading input file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		do(scanner.Text())
	}
}

func do(command string) {
	var result string

	// regex to capture Proper Nouns in sentences
	pn := regexp.MustCompile("(\\b[A-Z][a-z]*\\b)")
	// regex to capture numbers
	d := regexp.MustCompile("([0-9].)")

	// fmt.Println(command)
	if strings.Contains(command, "-") {
		parts := strings.Split(command, " - ")
		portalGraph.AddTarget(parts[0], parts[1])

		// Nothing to print
		return
	} else if strings.Contains(command, "jumps") {
		city := strings.Join(
			pn.FindAllString(command, -1), " ")
		depth, err := strconv.Atoi(
			strings.Trim(
				d.FindString(command), " "))
		if err != nil {
			log.Fatalf("Error parsing string to int: %v\n", err)
		}

		edges := portalGraph.GetEdgesAtDepth(city, depth)
		result = strings.Join(edges, ", ")
	} else if strings.Contains(command, "teleport") {
		excludeI := strings.Split(command, "from ")[1]
		cities := strings.Split(excludeI, " to ")
		fromCity := cities[0]
		toCity := cities[1]

		connected := portalGraph.TeleportToBFS(fromCity, toCity)
		result = "no"
		if connected {
			result = "yes"
		}
	} else if strings.Contains(command, "loop") {
		city := strings.Join(
			pn.FindAllString(command, -1), " ")

		hasCycle := portalGraph.CheckUniqueCycle(city)
		result = "no"
		if hasCycle {
			result = "yes"
		}
	}

	fmt.Printf("%s: %s\n", command, result)
}
