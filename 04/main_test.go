package main

import (
	"fmt"
	"testing"
)

func TestCountXMAS(t *testing.T) {
	lines := [][]rune{
		[]rune("XX.."),
		[]rune("XMAS"),
		[]rune(".AA."),
		[]rune(".S.S"),
	}
	c := CountAllDirections(lines, "XMAS")
	if c != 3 {
		t.Errorf("Count returned %d, expected 3", c)
	}
}

func TestGetDiags(t *testing.T) {
	lines := [][]rune{
		[]rune("XX.."),
		[]rune("XMAS"),
		[]rune(".AA."),
		[]rune(".S.S"),
	}
	se, sw := GetDiags(lines)
	if len(se) != 7 {
		t.Errorf("GetDiags returned %d, expected 7", len(se))
	}
	if len(sw) != 7 {
		t.Errorf("GetDiags returned %d, expected 7", len(sw))
	}

	if string(se[0]) != ".AA." {
		t.Errorf("GetDiags returned %s, expected .AA.", string(se[0]))
	}
	if string(se[1]) != ".M." {
		t.Errorf("GetDiags returned %s, expected .M.", string(se[1]))
	}
	if string(sw[0]) != "XMAS" {
		t.Errorf("GetDiags returned %s, expected XMAS", string(sw[0]))
	}
	if string(sw[1]) != "XA." {
		t.Errorf("GetDiags returned %s, expected XA.", string(sw[1]))
	}
	// print as strings
	fmt.Println("SE:")
	for _, l := range se {
		fmt.Println(string(l))
	}
	fmt.Println("SW:")
	for _, l := range sw {
		fmt.Println(string(l))
	}
}
