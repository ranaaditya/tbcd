package tbcd

import (
	"errors"
	"strings"
)

var TbcdSymbolMap map[string]string = map[string]string {
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"*": "1010",
	"#": "1011",
	"a": "1100",
	"b": "1101",
	"c": "1110",
	"f": "1111",
}

var reverseTbcdSymbolMap map[string]string = map[string]string{
	"0000": "0",
	"0001": "1",
	"0010": "2",
	"0011": "3",
	"0100": "4",
	"0101": "5",
	"0110": "6",
	"0111": "7",
	"1000": "8",
	"1001": "9",
	"1010": "*",
	"1011": "#",
	"1100": "a",
	"1101": "b",
	"1110": "c",
	"1111": "f",
}

var TbcdSymbols = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "*", "a", "b", "c"}

func EncodeToTBCD(input string) (string, error) {
	input = strings.ToLower(input)
	var output string = ""
	if len(input) < 1 { return output, errors.New("given string cannot be empty")}
	if len(input)%2 != 0 {
		input += "f"
	} else  {
		// nothing special
	}

	index := 0
	for index < len(input) {
		output += TbcdSymbolMap[string(input[index+1])] + TbcdSymbolMap[string(input[index])]
		index += 2
	}
	return output, nil
}

func DecodeToTBCD(input string) (string, error) {
	var output string = ""
	if len(input) < 4 { return output, errors.New("given string cannot be empty")}

	index := 0
	for index < len(input) {
		output += reverseTbcdSymbolMap[string(input[index+4])+string(input[index+5])+string(input[index+6])+string(input[index+7])] + reverseTbcdSymbolMap[string(input[index])+string(input[index+1])+string(input[index+2])+string(input[index+3])]
		index += 8
	}

	if output[len(output)-1] == 'f' {output = output[:len(output)-1]}
	return output, nil
}