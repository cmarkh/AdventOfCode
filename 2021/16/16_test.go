package main

import (
	"fmt"
	"testing"
)

func testPacket(hex string, t *testing.T) {
	fmt.Printf("Packet %s:\n", hex)
	bits, err := Binary(hex)
	if err != nil {
		t.Fatal(err)
	}
	versionSum, err := ParseBits(bits)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\nversion sum: %v\n\n", versionSum)
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
