package functions

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetColorCode(colorName string) string {
	colorMap := map[string]string{
		"blue":         "\033[34m",
		"brightblue":   "\033[38;5;33m",
		"brightcyan":   "\033[38;5;51m",
		"brightgreen":  "\033[38;5;46m",
		"brightred":    "\033[38;5;196m",
		"brightwhite":  "\033[38;5;15m",
		"brightyellow": "\033[38;5;226m",
		"brown":        "\033[38;5;130m",
		"chartreuse":   "\033[38;5;154m",
		"reset":        "\033[0m",
		"darkpurple":   "\033[38;5;53m",
		"darkred":      "\033[91m",
		"gold":         "\033[38;5;220m",
		"gray":         "\033[90m",
		"green":        "\033[32m",
		"indigo":       "\033[38;5;54m",
		"lime":         "\033[38;5;46m",
		"lightblue":    "\033[38;5;39m",
		"magenta":      "\033[38;5;200m",
		"maram":        "\033[38;5;206m",
		"olive":        "\033[38;5;100m",
		"orange":       "\033[38;5;208m",
		"peach":        "\033[38;5;222m",
		"periwinkle":   "\033[38;5;159m",
		"pink":         "\033[38;5;206m",
		"purple":       "\033[35m",
		"red":          "\033[31m",
		"salmon":       "\033[38;5;203m",
		"skyblue":      "\033[38;5;111m",
		"teal":         "\033[38;5;51m",
		"turquoise":    "\033[38;5;87m",
		"white":        "\033[37m",
		"yellow":       "\033[33m",
	}
	

	colorName = strings.ToLower(colorName)
	if colorCode, ok := colorMap[colorName]; ok {
		return colorCode
	}

	rgbRegex := regexp.MustCompile(`^rgb\((\d{1,3}),\s?(\d{1,3}),\s?(\d{1,3})\)$`)
	rgbMatches := rgbRegex.FindStringSubmatch(colorName)
	if len(rgbMatches) == 4 {
		r, _ := strconv.Atoi(rgbMatches[1])
		g, _ := strconv.Atoi(rgbMatches[2])
		b, _ := strconv.Atoi(rgbMatches[3])
		return CreateRGBColorCode(r, g, b)
	}

	if IsHexColor(colorName) {
		r, g, b := HexToRGB(colorName)
		return CreateRGBColorCode(r, g, b)
	}

	if IsHSLColor(colorName) {
		r, g, b := HSLToRGB(colorName)
		return CreateRGBColorCode(r, g, b)
	}

	fmt.Println("Unrecognized color:", colorName)
	return ""
}

func ConvertRGBToValues(color string) (int, int, int) {
	rgbPattern := regexp.MustCompile(`(\d+),\s*(\d+),\s*(\d+)`)
	matches := rgbPattern.FindStringSubmatch(color)
	if len(matches) != 4 {
		return 0, 0, 0
	}

	r, _ := strconv.Atoi(matches[1])
	g, _ := strconv.Atoi(matches[2])
	b, _ := strconv.Atoi(matches[3])

	return r, g, b
}

func CreateRGBColorCode(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func IsHexColor(color string) bool {
	match, _ := regexp.MatchString(`^#([a-fA-F0-9]{6})$`, color)
	return match
}

func HexToRGB(color string) (int, int, int) {
	color = color[1:]

	r, _ := strconv.ParseInt(color[:2], 16, 0)
	g, _ := strconv.ParseInt(color[2:4], 16, 0)
	b, _ := strconv.ParseInt(color[4:6], 16, 0)

	return int(r), int(g), int(b)
}

func IsHSLColor(color string) bool {
	match, _ := regexp.MatchString(`^hsl\(\d+,\s*\d+%?,\s*\d+%?\)$`, color)
	return match
}

func HSLToRGB(color string) (int, int, int) {
	hslPattern := regexp.MustCompile(`(\d+),\s*(\d+)%?,\s*(\d+)%?`)
	matches := hslPattern.FindStringSubmatch(color)
	if len(matches) != 4 {
		return 0, 0, 0
	}

	h, _ := strconv.Atoi(matches[1])
	s, _ := strconv.Atoi(matches[2])
	l, _ := strconv.Atoi(matches[3])

	h = h % 360
	s = int(math.Max(0, math.Min(100, float64(s))))
	l = int(math.Max(0, math.Min(100, float64(l))))

	c := (1 - math.Abs(float64(2*l-100))/100) * float64(s) / 100
	x := c * (1 - math.Abs(float64((h/60)%2-1)))
	m := float64(l)/100 - c/2

	var r, g, b float64
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	r = (r + m) * 255
	g = (g + m) * 255
	b = (b + m) * 255

	return int(r), int(g), int(b)
}
