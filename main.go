package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	s := flag.String("start", "", "starter names, order retained")
	n := flag.String("names", "", "names of standup members separated by comma, randomized")
	e := flag.String("end", "", "ending names, order retained")
	flag.Parse()

	starts := parseNames(s)
	names := parseNames(n)
	ends := parseNames(e)

	// randomize names
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	// merge all the names
	all := append(starts, names...)
	all = append(all, ends...)
	if len(all) == 0 {
		fmt.Printf("Nobody is coming to this standup :(\n")
		return
	}

	// print names
	fmt.Printf("Today is a workday:\n")
	for i, name := range all {
		fmt.Printf("[%d] - %s\n", i, name)
	}
}

func parseNames(n *string) []string {
	var names = []string{}
	if n == nil {
		return names
	}

	// Split by comman and trim whitespace
	rawNames := strings.Split(*n, ",")
	for _, name := range rawNames {
		trimmed := strings.TrimSpace(name)
		if trimmed == "" {
			continue
		}
		names = append(names, name)
	}

	return names
}
