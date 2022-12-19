package figure

import (
	"fmt"
)

var colors = map[string]string{
	"reset":  "\033[0m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"gray":   "\033[37m",
	"white":  "\033[97m",
}

func RGBColor(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func FixedColorizer(color string) colorizer {
	if color == "" {
		return nil
	}
	cstr, ok := colors[color]
	if !ok {
		cstr = color
	}
	return func(cidx, _ int, fragment, text string) string {
		return cstr + fragment + colors["reset"]
	}
}

func FixedRGBColorizer(r, g, b int) colorizer {
	return func(cidx, _ int, fragment, text string) string {
		return fmt.Sprintf("%s%s\033[0m", RGBColor(r, g, b), fragment)
	}
}

func GradientRGBColorizer(r1, g1, b1, r2, g2, b2 int) colorizer {
	return func(cidx, _ int, fragment, text string) string {
		r := r1 + (r2-r1)*cidx/len(text)
		g := g1 + (g2-g1)*cidx/len(text)
		b := b1 + (b2-b1)*cidx/len(text)
		return fmt.Sprintf("%s%s\033[0m", RGBColor(r, g, b), fragment)
	}
}
