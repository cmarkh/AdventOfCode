package day8

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func input(filename string) (Instructions, Map) {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return ParseInput(string(input))
}

func TestInput(t *testing.T) {
	instructions, mapIn := input("ex1.txt")
	fmt.Println(instructions)
	fmt.Println(mapIn)
}

/*
#[case("ex1.txt" => 2)]
    #[case("ex2.txt" => 6)]
    #[case("input.txt" => 20777)]
    fn test_part1(input: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/{}", input)).unwrap();
        let (instructions, map) = parse_input(&input);
        part1(instructions, map)
    }
*/

func TestEx1(t *testing.T) {
	instructions, mapIn := input("ex1.txt")
	res := Part1(instructions, mapIn)
	expected := 2
	if res != uint64(expected) {
		t.Fatalf("expected %d, got %d", expected, res)
	}
}

func TestEx2(t *testing.T) {
	instructions, mapIn := input("ex2.txt")
	res := Part1(instructions, mapIn)
	expected := 6
	if res != uint64(expected) {
		t.Fatalf("expected %d, got %d", expected, res)
	}
}

func TestPart1(t *testing.T) {
	instructions, mapIn := input("input.txt")
	res := Part1(instructions, mapIn)
	expected := 20777
	if res != uint64(expected) {
		t.Fatalf("expected %d, got %d", expected, res)
	}
}
