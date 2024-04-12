package crypto

import (
	gobytes "bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

func Mine(data string, ruleString string) (int64, string, error) {

	rule, err := strconv.Atoi(ruleString)
	if err != nil {
		return 0, "", err
	}

	bytes := []byte(data)

	var PoW int64 // proof of work
	var hash string

	for {
		PoW += 1
		finalString := gobytes.Join([][]byte{bytes, []byte(strconv.FormatInt(PoW, 10))}, []byte("-"))
		hash = generateHash(finalString)
		if strings.HasPrefix(hash, strconv.FormatInt(int64(rule), 10)) {
			return PoW, hash, nil
		}
	}
}

func generateHash(bytes []byte) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", bytes)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
