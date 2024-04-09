package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Person struct {
	Name   string
	Amount int
}

func main() {
	fmt.Println(sumAmountsFromFile("./data1.txt"))
}

func ParseFile(filename string) ([]Person, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}
	defer f.Close()

	return ParseReader(f)
}

func ParseReader(rd io.Reader) ([]Person, error) {
	data, err := io.ReadAll(rd)
	if err != nil {
		return nil, errors.Wrap(err, "reading data")
	}

	lines := strings.Split(string(data), "\n")
	persons := make([]Person, 0, len(lines))
	for _, line := range lines {
		person, err := parseLine(line)
		if err != nil {
			var myErr incorrectLineError
			if errors.As(err, &myErr) {
				return nil, errors.New("incorrect line:" + myErr.line)
			}
			return nil, errors.Wrap(err, "parsing person")
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func sumAmountsFromFile(filename string) (map[string]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}
	defer f.Close()

	return sumAmountsFromReader(f)
}

func sumAmountsFromReader(r io.Reader) (map[string]int, error) {
	buf := bufio.NewReader(r)
	aggregations := make(map[string]int)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.Wrap(err, "reading line")
		}

		line = strings.TrimSuffix(line, "\n")

		person, err := parseLine(line)
		if err != nil {
			return nil, errors.Wrap(err, "parsing line")
		}

		fmt.Println(person.Name, person.Amount)
		aggregations[person.Name] += person.Amount
	}
	return aggregations, nil
}

var lineRe = regexp.MustCompile(`^Name:([^,]+), Amount:(-?\d+)$`)

// var errIncorrectLine = errors.New("the line is incorrect")
var errCannotParseAmount = errors.New("cannot parse amount")

type incorrectLineError struct {
	line string
}

func (e incorrectLineError) Error() string {
	return "incorrect line"
}

func parseLine(line string) (Person, error) {
	matches := lineRe.FindStringSubmatch(line)
	if len(matches) < 3 {
		return Person{}, incorrectLineError{
			line: line,
		}
	}

	name := matches[1]
	amountStr := matches[2]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return Person{}, errCannotParseAmount
	}

	return Person{
		Name:   name,
		Amount: amount,
	}, nil
}
