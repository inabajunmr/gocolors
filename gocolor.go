package gocolor

import "image/color"

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

	// ERROR
	return color.RGBA{255, 255, 255, alpha}
}

func ValueOf(c string, alpha uint8) color.RGBA {

	switch c {
	case "Black":
		return Of(Black, alpha)
	case "Navy":
		return Of(Black, alpha)
	case "Blue":
		return Of(Black, alpha)
	case "Green":
		return Of(Black, alpha)
	case "Teal":
		return Of(Black, alpha)
	case "Lime":
		return Of(Black, alpha)
	case "Aqua":
		return Of(Black, alpha)
	case "Maroon":
		return Of(Black, alpha)
	case "Purple":
		return Of(Black, alpha)
	case "Olive":
		return Of(Black, alpha)
	case "Gray":
		return Of(Black, alpha)
	case "Silver":
		return Of(Black, alpha)
	case "Red":
		return Of(Black, alpha)
	case "Fuchsia":
		return Of(Black, alpha)
	case "Yellow":
		return Of(Black, alpha)
	case "White":
		return Of(Black, alpha)

	}

	// ERROR
	return color.RGBA{255, 255, 255, alpha}
}
