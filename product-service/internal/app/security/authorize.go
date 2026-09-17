package security

import (
	"bufio"

	"os"
	"strings"

	"github.com/go-openapi/errors"
)

func Authorize(token string) (interface{}, error) {

	if _, found := authorizedKeys[token]; found {
		return &DummyPrincipal{}, nil
	}
	return nil, errors.Unauthenticated("invalid API key")
}

const (
	secret = "secret/authorized.txt"
)

type DummyPrincipal struct{}

var authorizedKeys = make(map[string]struct{})

func init() {
	if _, err := os.Stat(secret); err != nil {
		panic("unable to read secret/authorized.txt")
	}

	file, _ := os.Open(secret)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		authorizedKeys[strings.TrimSpace(scanner.Text())] = struct{}{}
	}
}
