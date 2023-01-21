package main

func main() {

}

func reverseString(input string) string {
	if len(input) == 0 {
		return ""
	}

	output := make([]byte, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return string(output)
}
