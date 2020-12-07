package advent2020

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Bag struct {
	description string
	content     map[string]int
}

func parseBag(spec string) (Bag, error) {
	description := ""
	content := map[string]int{}

	for i, c := range spec {
		dLen := len(description)
		if dLen > 4 {
			if description[dLen-4:dLen] == "bags" {
				spec = spec[i:]
				break
			}
		}

		description += string(c)
	}

	if len(spec) == 0 {
		return Bag{}, fmt.Errorf("Bag '%s' doesn't contain anything?", description)
	}

	spec = strings.TrimSpace(spec)
	if !strings.HasPrefix(spec, "contain ") {
		return Bag{}, fmt.Errorf("For bag '%s', expected content to start with 'contain', got '%s'", description, spec)
	}

	bags := strings.Split(strings.TrimSuffix(spec[7:], "."), ",")
	for _, bagSpec := range bags {
		bagSpec = strings.TrimSpace(bagSpec)
		if bagSpec == "no other bags" {
			break
		}

		if strings.HasSuffix(bagSpec, "bag") {
			bagSpec += "s"
		}

		parts := strings.SplitN(bagSpec, " ", 2)
		if len(parts) != 2 {
			return Bag{}, fmt.Errorf("(1) for bag '%s', could not parse specification '%s'", description, bagSpec)
		}

		count, err := strconv.Atoi(parts[0])
		if err != nil {
			return Bag{}, fmt.Errorf("(2) for bag '%s', could not parse specification '%s'", description, bagSpec)
		}

		content[parts[1]] = count
	}

	return Bag{
		description: description,
		content:     content,
	}, nil
}

func invertContainers(bags []Bag) map[string][]string {
	bagParents := map[string][]string{}

	for _, bag := range bags {
		for container := range bag.content {
			if _, exists := bagParents[container]; !exists {
				bagParents[container] = []string{bag.description}
			} else {
				bagParents[container] = append(bagParents[container], bag.description)
			}
		}
	}

	return bagParents
}

func ParentsOf(bagDesc string, bags []Bag) []string {
	invertedTree := invertContainers(bags)

	baseList := invertedTree[bagDesc]
	list := make([]string, len(baseList))
	copy(list, baseList)
	for _, parent := range baseList {
		list = append(list, ParentsOf(parent, bags)...)
	}
	sort.Strings(list)
	cleanedList := make([]string, 0, len(list))
	for _, bag := range list {
		if len(cleanedList) != 0 && bag == cleanedList[len(cleanedList)-1] {
			continue
		}
		cleanedList = append(cleanedList, bag)
	}
	return cleanedList
}

func ChildCount(bagDesc string, bags []Bag) int {
	bagMap := map[string]Bag{}
	for _, bag := range bags {
		bagMap[bag.description] = bag
	}

	total := 0
	for childBag, count := range bagMap[bagDesc].content {
		total += count + count*ChildCount(childBag, bags)
	}
	return total
}

func AllBags(stream <-chan string) ([]Bag, error) {
	bags := []Bag{}
	for spec := range stream {
		bag, err := parseBag(spec)
		if err != nil {
			return nil, err
		}
		bags = append(bags, bag)
	}
	return bags, nil
}
