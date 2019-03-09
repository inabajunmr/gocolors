package gocolor

import (
	"fmt"
	"testing"
)

func TestOf(t *testing.T) {

	color := Of(Black, 255)
	fmt.Println(color)
	if color.A != 255 {
		t.Fatal("Return unexpected color.")
	}
	if color.B != 0 {
		t.Fatal("Return unexpected color.")
	}
	if color.G != 0 {
		t.Fatal("Return unexpected color.")
	}
	if color.R != 0 {
		t.Fatal("Return unexpected color.")
	}
}

func TestOfValue(t *testing.T) {
	color, _ := ValueOf("Black", 255)

	if color != Of(Black, 255) {
		t.Fatal("Return unexpected color.")
	}

}

func TestOfValue_UnknownColor(t *testing.T) {
	_, err := ValueOf("Unknown", 255)

	if err == nil {
		t.Fatal("Not return error by unknon color.")
	}

}
