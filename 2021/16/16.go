package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	advent "temp/adventofcode/go/2021"
)

//https://adventofcode.com/2021/day/16

func main() {
	hex, err := input()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(hex)

	fmt.Printf("Packet %s:\n\n", hex)
	bits, err := Binary(hex)
	if err != nil {
		log.Fatal(err)
	}
	parentPacket, err := ParseBits(bits)
	if err != nil {
		log.Fatal(err)
	}
	parentPacket.Print()
	fmt.Printf("Version sum: %v\n\n", parentPacket.SumVersions())

	fmt.Printf("Outermost packet value: %v\n", parentPacket.CalculateValue())
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
	//fmt.Println(bits)

	return
}

func ParseBits(bits string) (packet Packet, err error) {
	packet, _, err = parseBits(bits)
	return
}

func parseBits(bits string) (packet Packet, bitsLeft string, err error) {
	packet.Bits = bits

	bits, err = packet.header(bits)
	if err != nil {
		return
	}

	switch packet.TypeID {
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
		packet.Value, err = strconv.ParseInt(numberFromBits, 2, 64)
		if err != nil {
			return
		}

	default:
		//Every other type of packet (any packet with a type ID other than 4) represent an operator that performs some calculation on one or more sub-packets contained within
		//An operator packet contains one or more packets. To indicate which subsequent binary data represents its sub-packets, an operator packet can use one of two modes indicated by the bit immediately after the packet header; this is called the length type ID
		packet.LengthTypeID = string(bits[0])
		bits = bits[1:]

		switch packet.LengthTypeID {
		case "0":
			packet.BitLengthOfSubPackets, err = strconv.ParseInt(bits[0:15], 2, 64)
			if err != nil {
				return
			}

			childBits := bits[15 : 15+packet.BitLengthOfSubPackets]
			bits = bits[15+packet.BitLengthOfSubPackets:]

			for {
				var childPacket Packet
				childPacket, childBits, err = parseBits(childBits)
				if err != nil {
					return packet, "", err
				}
				packet.ChildPackets = append(packet.ChildPackets, childPacket)

				if len(childBits) == 0 {
					break
				}
			}

		case "1":
			packet.NumberOfSubPackets, err = strconv.ParseInt(bits[0:11], 2, 64)
			if err != nil {
				return
			}
			bits = bits[11:]

			for i := 0; i < int(packet.NumberOfSubPackets); i++ {
				var childPacket Packet
				childPacket, bits, err = parseBits(bits)
				if err != nil {
					return packet, "", err
				}
				packet.ChildPackets = append(packet.ChildPackets, childPacket)
			}

		default:
			err = fmt.Errorf("unkown length type ID: %v", packet.LengthTypeID)
			return
		}
	}

	return packet, bits, nil
}

func (packet *Packet) CalculateValue() (val int64) {
	switch packet.TypeID {
	case 0:
		packet.Value = 0
		for _, child := range packet.ChildPackets {
			packet.Value += child.CalculateValue()
		}

	case 1:
		packet.Value = 1 //value is product of child packets
		for _, child := range packet.ChildPackets {
			packet.Value *= child.CalculateValue()
		}

	case 2:
		minimum := int64(math.MaxInt64 / 2)
		for _, child := range packet.ChildPackets {
			childVal := child.CalculateValue()
			if childVal < minimum {
				minimum = childVal
			}
		}
		packet.Value = minimum

	case 3:
		var max int64 = 0
		for _, child := range packet.ChildPackets {
			childVal := child.CalculateValue()
			if childVal > max {
				max = childVal
			}
		}
		packet.Value = max

	case 4:
		break //packet.Value is already set for literal value packets

	case 5:
		//These packets always have exactly two sub-packets
		if packet.ChildPackets[0].CalculateValue() > packet.ChildPackets[1].CalculateValue() {
			packet.Value = 1
		} else {
			packet.Value = 0
		}

	case 6:
		//These packets always have exactly two sub-packets
		if packet.ChildPackets[0].CalculateValue() < packet.ChildPackets[1].CalculateValue() {
			packet.Value = 1
		} else {
			packet.Value = 0
		}

	case 7:
		//These packets always have exactly two sub-packets
		if packet.ChildPackets[0].CalculateValue() == packet.ChildPackets[1].CalculateValue() {
			packet.Value = 1
		} else {
			packet.Value = 0
		}

	}

	return packet.Value
}

type Packet struct {
	Bits string
	Header
	OperatingPacket
	Value int64
}
type Header struct {
	Version int64
	TypeID  int64
}
type OperatingPacket struct {
	LengthTypeID          string
	BitLengthOfSubPackets int64
	NumberOfSubPackets    int64
	ChildPackets          []Packet
}

func (packet *Packet) header(bits string) (bitsLeft string, err error) {
	packet.Version, err = strconv.ParseInt(bits[:3], 2, 64)
	if err != nil {
		return
	}
	packet.TypeID, err = strconv.ParseInt(bits[3:6], 2, 64)
	if err != nil {
		return
	}
	bits = bits[6:] //remove header
	return bits, nil
}

func (packet Packet) SumVersions() (sum int64) {
	sum = packet.Version
	for _, child := range packet.ChildPackets {
		sum += child.SumVersions()
	}
	return
}

func (packet Packet) Print() {
	fmt.Println(packet.Bits)
	fmt.Printf("Version: %v\n", packet.Version)
	fmt.Printf("TypeID: %v\n", packet.TypeID)
	if packet.BitLengthOfSubPackets == 0 && packet.NumberOfSubPackets == 0 {
		fmt.Printf("Value: %v\n", packet.Value)
	} else {
		fmt.Printf("Number of child packets: %v\n", len(packet.ChildPackets))
		fmt.Printf("Length Type ID: %v\n", packet.LengthTypeID)
		if packet.BitLengthOfSubPackets > 0 {
			fmt.Printf("Bit length of sub packets: %v\n", packet.BitLengthOfSubPackets)
		}
		if packet.NumberOfSubPackets > 0 {
			fmt.Printf("Number of sub packets: %v\n", packet.NumberOfSubPackets)
		}
	}
	fmt.Println()

	for _, child := range packet.ChildPackets {
		fmt.Printf("Parent version: %v\n", packet.Version)
		child.Print()
	}
}
