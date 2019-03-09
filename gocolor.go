package gocolor

import (
	"fmt"
	"image/color"
)

// Color is enumration for color name expression
type Color int

const (
	Black Color = iota
	Navy
	Blue
	Green
	Teal
	Lime
	Aqua
	Maroon
	Purple
	Olive
	Gray
	Silver
	Red
	Fuchsia
	Yellow
	White
)

// Get color.RGBA from color and alpha
func Of(c Color, alpha uint8) color.RGBA {
	switch c {
	case Black:
		return color.RGBA{0, 0, 0, alpha}
	case Navy:
		return color.RGBA{0, 0, 128, alpha}
	case Blue:
		return color.RGBA{0, 0, 255, alpha}
	case Green:
		return color.RGBA{0, 128, 0, alpha}
	case Teal:
		return color.RGBA{0, 128, 128, alpha}
	case Lime:
		return color.RGBA{0, 255, 0, alpha}
	case Aqua:
		return color.RGBA{0, 255, 255, alpha}
	case Maroon:
		return color.RGBA{128, 0, 0, alpha}
	case Purple:
		return color.RGBA{128, 0, 128, alpha}
	case Olive:
		return color.RGBA{128, 128, 0, alpha}
	case Gray:
		return color.RGBA{128, 128, 128, alpha}
	case Silver:
		return color.RGBA{192, 192, 192, alpha}
	case Red:
		return color.RGBA{255, 0, 0, alpha}
	case Fuchsia:
		return color.RGBA{255, 0, 255, alpha}
	case Yellow:
		return color.RGBA{255, 255, 0, alpha}
	case White:
		return color.RGBA{255, 255, 255, alpha}
	}

	// Unexpected
	return color.RGBA{255, 255, 255, alpha}
}

// Get color.RGBA from color name and alpha. If any color unmatched with parameter, return error.
func ValueOf(c string, alpha uint8) (color.RGBA, error) {

	switch c {
	case "Black":
		return Of(Black, alpha), nil
	case "Navy":
		return Of(Navy, alpha), nil
	case "Blue":
		return Of(Black, alpha), nil
	case "Green":
		return Of(Green, alpha), nil
	case "Teal":
		return Of(Teal, alpha), nil
	case "Lime":
		return Of(Lime, alpha), nil
	case "Aqua":
		return Of(Aqua, alpha), nil
	case "Maroon":
		return Of(Maroon, alpha), nil
	case "Purple":
		return Of(Purple, alpha), nil
	case "Olive":
		return Of(Olive, alpha), nil
	case "Gray":
		return Of(Gray, alpha), nil
	case "Silver":
		return Of(Silver, alpha), nil
	case "Red":
		return Of(Red, alpha), nil
	case "Fuchsia":
		return Of(Fuchsia, alpha), nil
	case "Yellow":
		return Of(Yellow, alpha), nil
	case "White":
		return Of(White, alpha), nil

	}

	return Of(White, alpha), fmt.Errorf("%v is not implemented", c)
}
