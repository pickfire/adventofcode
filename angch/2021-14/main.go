package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Edge struct {
	from, to string
}

func day14(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	input := scanner.Text()

	counts := make(map[Edge]int)
	prev := 'x'
	for _, v := range input {
		if prev != 'x' {
			counts[Edge{from: string(prev), to: string(v)}]++
		}
		prev = v
	}

	scanner.Scan()
	scanner.Text()

	rules := make(map[Edge]rune)

	for scanner.Scan() {
		t := scanner.Text()
		rule := strings.Split(t, " -> ")
		from, to := string(rule[0][0]), string(rule[0][1])
		add := rule[1][0]
		rules[Edge{from: from, to: to}] = rune(add)
	}
	total := 1
	for _, v := range counts {
		total += v
	}
	// fmt.Println("Step", 0, counts, total)
	part1, part2 := 0, 0
	for step := 0; step < 40; step++ {
		newcounts := make(map[Edge]int)
		for k, v := range counts {
			newcounts[k] = v
		}
		for rule, to2 := range rules {
			to := string(to2)
			c, ok := counts[Edge{from: rule.from, to: rule.to}]
			if !ok {
				continue
			}
			// fmt.Println("Rule", rule, to)
			newcounts[Edge{from: rule.from, to: rule.to}] -= c
			newcounts[Edge{from: rule.from, to: to}] += c
			newcounts[Edge{from: to, to: rule.to}] += c
			// fmt.Println(" ", newcounts)
		}
		counts = newcounts

		total := 1
		for k, v := range counts {
			total += v
			if v == 0 {
				delete(counts, k)
			}
		}
		// fmt.Println("Step", step+1, counts, total)

		if step == 9 || step == 39 {
			elem := make(map[string]int)
			for c, v := range counts {
				elem[c.from] += v
			}
			elem[string(prev)]++
			// fmt.Println(elem)

			least := 999999999999999
			most := -1
			for _, v := range elem {
				if most < v {
					most = v
				}
				if least > v {
					least = v
				}
			}
			if step == 9 {
				part1 = most - least
			} else {
				//4019217425979
				part2 = most - least
			}
		}
	}
	// fmt.Println("Step", step, counts)

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func main() {
	// day14("test.txt")

	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	day14("input.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
