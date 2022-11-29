package main

import (
	"fmt"
	"testing"
)

func testPacket(hex string, t *testing.T) {
	fmt.Printf("Packet %s:\n\n", hex)
	bits, err := Binary(hex)
	if err != nil {
		t.Fatal(err)
	}
	parentPacket, err := ParseBits(bits)
	if err != nil {
		t.Fatal(err)
	}
	parentPacket.Print()
	fmt.Printf("Version sum: %v\n\n", parentPacket.SumVersions())
}

func TestLiteralPacket1(t *testing.T) {
	testPacket("D2FE28", t)
}

func TestOperatorPacket1(t *testing.T) {
	testPacket("38006F45291200", t)
}

func TestOperatorPacket2(t *testing.T) {
	testPacket("EE00D40C823060", t)
}

func TestOperatorPacket3(t *testing.T) {
	testPacket("8A004A801A8002F478", t)
}

func TestOperatorPacket4(t *testing.T) {
	testPacket("620080001611562C8802118E34", t)
}

func TestOperatorPacket5(t *testing.T) {
	testPacket("C0015000016115A2E0802F182340", t)
}

func TestOperatorPacket6(t *testing.T) {
	testPacket("A0016C880162017C3686B18A3D4780", t)
}

func testCalculateValues(hex string, t *testing.T) {
	bits, err := Binary(hex)
	if err != nil {
		t.Fatal(err)
	}
	parentPacket, err := ParseBits(bits)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Parent packet %s value: %v\n", hex, parentPacket.CalculateValue())
}

func TestCalculateValues1(t *testing.T) {
	testCalculateValues("C200B40A82", t)
}

func TestCalculateValues2(t *testing.T) {
	testCalculateValues("04005AC33890", t)
}

func TestCalculateValues3(t *testing.T) {
	testCalculateValues("880086C3E88112", t)
}

func TestCalculateValues4(t *testing.T) {
	testCalculateValues("CE00C43D881120", t)
}

func TestCalculateValues5(t *testing.T) {
	testCalculateValues("D8005AC2A8F0", t)
}

func TestCalculateValues6(t *testing.T) {
	testCalculateValues("F600BC2D8F", t)
}

func TestCalculateValues7(t *testing.T) {
	testCalculateValues("9C005AC2F8F0", t)
}

func TestCalculateValues8(t *testing.T) {
	testCalculateValues("9C0141080250320F1802104A08", t)
}
