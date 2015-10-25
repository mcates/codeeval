package main
import "testing"

func TestPenultimate(t *testing.T) {
	inputs := []string{"some line with text", "another line"}
	val := Penultimate(inputs[0])
	if val != "with" {
		t.Errorf("%s: should return %s, but returns %s", inputs[0], "with", val)
	}
	val = Penultimate(inputs[1])
	if val != "another" {
		t.Errorf("%s: should return %s, but returns %s", inputs[1], "another", val)
	}
}
