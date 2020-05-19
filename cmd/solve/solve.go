package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hsmtkk/paiza_c051/pkg/solve"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	ss := strings.Split(string(input), " ")
	if len(ss) != 4 {
		log.Fatal("failed to parse input")
	}
	numbers := make([]int, 4)
	for i, s := range ss {
		numbers[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	max := solve.Solve(numbers)
	fmt.Fprintf(os.Stdout, "%d", max)
}
