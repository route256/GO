package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseLine_ShouldFillPersonFields(t *testing.T) {
	line := "Name:Сергей, Amount:10"

	person, err := parseLine(line)

	assert.NoError(t, err)
	assert.Equal(t,
		Person{
			Name:   "Сергей",
			Amount: 10,
		},
		person,
	)
}

func Test_ParseLine_ShouldFillPersonFields_WhenAmountIsNegative(t *testing.T) {
	line := "Name:Сергей, Amount:-10"

	person, err := parseLine(line)

	assert.NoError(t, err)
	assert.Equal(t, Person{
		Name:   "Сергей",
		Amount: -10,
	}, person)
}

func Test_ParseLine_ShouldReturnError_WhenNameIsAbsent(t *testing.T) {
	line := "Name:, Amount:10"

	_, err := parseLine(line)

	assert.Error(t, err)
}
func Test_ParseLine_ShouldReturnError_WhenAmountIsNotANumber(t *testing.T) {
	line := "Name:Сергей, Amount:1asd0"

	_, err := parseLine(line)

	assert.Error(t, err)
}

func Test_ParseLine_ShouldReturnError_WhenAmountIsEmpty(t *testing.T) {
	line := "Name:Сергей, Amount:"

	_, err := parseLine(line)

	assert.Error(t, err)
}

func Test_ParseLine_ShouldReturnError_WhenAmountTooBig(t *testing.T) {
	line := "Name:Сергей, Amount:1111111111111111111111111111111111111111111111111111111"

	_, err := parseLine(line)

	assert.Equal(t, errCannotParseAmount, err)
}

func Test_ParseReader_WrongLineShouldGiveError(t *testing.T) {
	data := "Nam, Amount:12\n" +
		"Name:Петр, Amount:65\n"
	buf := bytes.NewBufferString(data)

	_, err := ParseReader(buf)

	assert.Error(t, err)
}

func Test_ParseReader_CorrentLineShouldGivePersonsList(t *testing.T) {
	data := "Name:Иван, Amount:12\n" +
		"Name:Петр, Amount:65"
	buf := bytes.NewBufferString(data)

	persons, err := ParseReader(buf)

	assert.NoError(t, err)
	assert.Equal(t,
		[]Person{
			{
				Name:   "Иван",
				Amount: 12,
			},
			{
				Name:   "Петр",
				Amount: 65,
			},
		},
		persons,
	)
}
