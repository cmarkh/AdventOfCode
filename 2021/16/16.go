package main

import (
	"fmt"
	"log"
	"strconv"
	advent "temp/adventofcode/go"
)

//https://adventofcode.com/2021/day/16

func main() {
	hex, err := input()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(hex)

	Binary(hex)
}

func input() (hex string, err error) {
	lines, err := advent.ReadInput("input.txt")
	if err != nil {
		return
	}
	if len(lines) != 1 {
		err = fmt.Errorf("expected just one line of input. got: %v", lines)
		return
	}
	return lines[0], nil
}

func Binary(hex string) (bits string, err error) {
	encoding := make(map[rune]string)
	encoding['0'] = "0000"
	encoding['1'] = "0001"
	encoding['2'] = "0010"
	encoding['3'] = "0011"
	encoding['4'] = "0100"
	encoding['5'] = "0101"
	encoding['6'] = "0110"
	encoding['7'] = "0111"
	encoding['8'] = "1000"
	encoding['9'] = "1001"
	encoding['A'] = "1010"
	encoding['B'] = "1011"
	encoding['C'] = "1100"
	encoding['D'] = "1101"
	encoding['E'] = "1110"
	encoding['F'] = "1111"

	for _, ch := range hex {
		//fmt.Printf("%c: %v\n", ch, encoding[ch])
		bit, ok := encoding[ch]
		if !ok {
			err = fmt.Errorf("unkown hex symbol: %v", ch)
			return
		}
		bits += bit
	}
	fmt.Println(bits)

	return
}

func ParseBits(bits string) (sumVersionNumbers int64, err error) {
	return parseBits(bits, -1, 0, -1)
}

func parseBits(bits string, packetsRemaining int64, versionSums int64, parentVersion int64) (sumVersionNumbers int64, err error) {
	sumVersionNumbers = versionSums
	if len(bits) == 0 {
		return
	}
	if packetsRemaining == 0 {
		return
	}
	if val, err := strconv.ParseInt(bits, 2, 64); err == nil && val == 0 {
		return sumVersionNumbers, err
	}

	packetsRemaining--

	fmt.Println()
	fmt.Printf("parent version: %v\n", parentVersion)
	fmt.Println(bits)

	version, err := strconv.ParseInt(bits[:3], 2, 64)
	if err != nil {
		return sumVersionNumbers, err
	}
	fmt.Printf("version: %v\n", version)
	sumVersionNumbers += version

	typeID, err := strconv.ParseInt(bits[3:6], 2, 64)
	if err != nil {
		return sumVersionNumbers, err
	}
	fmt.Printf("type ID: %v\n", typeID)

	bits = bits[6:] //remove header

	switch typeID {
	case 4:
		//Packets with type ID 4 represent a literal value. Literal value packets encode a single binary number. To do this, the binary number is padded with leading zeroes until its length is a multiple of four bits, and then it is broken into groups of four bits. Each group is prefixed by a 1 bit except the last group, which is prefixed by a 0 bit
		last := false
		var numberFromBits string
		for !last {
			if bits[0] == '0' {
				last = true
			}
			//fmt.Println(bits[0:5])
			numberFromBits += bits[1:5]
			bits = bits[5:]
		}
		fmt.Printf("number: %v\n", numberFromBits)
		decimal, err := strconv.ParseInt(numberFromBits, 2, 64)
		if err != nil {
			return sumVersionNumbers, err
		}
		fmt.Printf("decimal: %v\n", decimal)

		return parseBits(bits, -1, sumVersionNumbers, version)

	default:
		//Every other type of packet (any packet with a type ID other than 4) represent an operator that performs some calculation on one or more sub-packets contained within
		//An operator packet contains one or more packets. To indicate which subsequent binary data represents its sub-packets, an operator packet can use one of two modes indicated by the bit immediately after the packet header; this is called the length type ID
		lengthTypeID := bits[0]
		fmt.Printf("length type ID: %v\n", string(lengthTypeID))
		bits = bits[1:]

		switch lengthTypeID {
		case '0':
			//If the length type ID is 0, then the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet.
			lengthInBitsOfSubPackets, err := strconv.ParseInt(bits[0:15], 2, 32)
			if err != nil {
				return sumVersionNumbers, err
			}
			fmt.Printf("length (in bits) of sub packets: %v\n", lengthInBitsOfSubPackets)
			bits = bits[15:] //remove any extra bits //15+lengthInBitsOfSubPackets

			return parseBits(bits, -1, sumVersionNumbers, version)

		case '1':
			//If the length type ID is 1, then the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet.
			packetsRemaining, err = strconv.ParseInt(bits[0:11], 2, 32) //note this is same variable used to test for loop condition
			if err != nil {
				return sumVersionNumbers, err
			}
			fmt.Printf("number of sub packets: %v\n", packetsRemaining)
			bits = bits[11:]

			return parseBits(bits, packetsRemaining, sumVersionNumbers, version)

		default:
			err = fmt.Errorf("unkown length type ID: %v", lengthTypeID)
			return
		}

	}
}
