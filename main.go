package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	// Flags / parameters
	filename := flag.String("f", "", "file to search")
	reString := flag.String("r", "(/d)", "regex to search for")
	sortNums := flag.Bool("s", false, "sort numbers in results. This works if capture group catches numbers.")
	verbose := flag.Bool("v", false, "give som extra info")
	flag.Parse()

	err := run(*verbose, *filename, *reString, *sortNums)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

func run(verbose bool, filename string, reString string, sortNums bool) error {
	if verbose {
		fmt.Printf("\n------- PARAMS -------\n")
		fmt.Printf("filename: %s\nregex: %s\nsort: %v\n", filename, reString, sortNums)
		fmt.Printf("\n------ DATA -------\ncapture - frequency\n-------------------\n")
	}

	// Read file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fiStats, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fiStats.Size()

	// Read file into buffer
	bytes := make([]byte, fileSize)
	_, err = file.Read(bytes)
	if err != nil {
		return err
	}
	file.Close()

	// Compile and run regex
	re := regexp.MustCompile(reString)
	res := re.FindAllSubmatch(bytes, -1)

	// Create occurencies map
	unique := make(map[string]int)
	for _, r := range res {
		if unique[string(r[1])] < 1 {
			unique[string(r[1])] = 1
		} else {
			unique[string(r[1])] = unique[string(r[1])] + 1
		}
	}

	// Print in numerical order
	if sortNums {
		keys := make([]int, 0, len(unique))
		for k := range unique {
			id, err := strconv.Atoi(string(k))
			if err != nil {
				return err
			}
			keys = append(keys, id)
		}
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Printf("%d - %d\n", k, unique[fmt.Sprint(k)])
		}
	} else {
		for k, v := range unique {
			fmt.Printf("%s - %d\n", k, v)
		}
	}

	// find total number of products
	count := 0
	for _, v := range unique {
		count += v
	}
	if verbose {
		fmt.Printf("\n------- STATS -------\n")
		kilo := int64(1024)
		mega := int64(kilo * kilo)
		if fileSize > 10*mega {
			fileSizeStr := strconv.FormatInt(fileSize/mega, 10)
			fmt.Printf("size: %s MB (%d Bytes)\n", fileSizeStr, fileSize)
		} else if fileSize > 10*kilo {
			fileSizeStr := strconv.FormatInt(fileSize/kilo, 10)
			fmt.Printf("size: %s KB (%d Bytes)\n", fileSizeStr, fileSize)
		} else {
			fileSizeStr := strconv.FormatInt(fileSize, 10)
			fmt.Printf("size: %s B (%d Bytes)\n", fileSizeStr, fileSize)
		}
		fmt.Printf("unique matches: %d\n", len(unique))
		fmt.Printf("total matches: %d\n", count)
	}
	return nil
}
