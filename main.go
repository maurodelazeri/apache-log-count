package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readFile("file")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	m := make(map[string]int)
	for _, text := range lines {
		data := strings.Split(text, " ")
		action := data[5][1:]
		m[data[0]+":"+action]++
	}
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	total := 0
	for _, k := range a {
		for _, s := range n[k] {
			if total > 10 {
				return
			}
			values := strings.Split(s, ":")
			fmt.Printf("%s %s %d\n", values[0], values[1], k)
			total++
		}
	}
}
