package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input struct{}

func (i Input) NextFloat() (float64, bool) {
	var input float64
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		//fmt.Println("Error:", err)
		return 0, false
	}
	return input, true
}

func (i Input) NextInt() (int, bool) {
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		//fmt.Println("Error:", err)
		return 0, false
	}
	return input, true
}

func (i Input) NextLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}

	input = strings.TrimSuffix(input, "\n")
	return input
}
