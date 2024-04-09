package main

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

type Person struct {
	Name   string
	Amount int
}

var lineRe = regexp.MustCompile(`^Name:([^,]+), Amount:(-?\d+)\n?$`)

var errIncorrectLine = errors.New("the line is incorrect")
var errCannotParseAmount = errors.New("cannot parse amount")

func parseLineV1(line string) (Person, error) {
	matches := lineRe.FindStringSubmatch(line)
	if len(matches) < 3 {
		return Person{}, errIncorrectLine
	}

	nameStr := matches[1]
	amountStr := matches[2]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return Person{}, errCannotParseAmount
	}

	return Person{
		Name:   nameStr,
		Amount: amount,
	}, nil
}

func parseLineV2(line string) (Person, error) {
	matches := lineRe.FindStringSubmatch(line)
	if len(matches) < 3 {
		return Person{}, errors.New("incorrect line")
	}

	nameStr := matches[1]
	amountStr := matches[2]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return Person{}, errors.Wrap(err, "incorrect amount")
	}

	return Person{
		Name:   nameStr,
		Amount: amount,
	}, nil
}
