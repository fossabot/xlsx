package format

import (
	"strings"
	"github.com/plandem/xlsx/internal/ml"
)

var (
	indexedColors []string
)

func newColor(color string) *ml.Color {
	color = normalizeColor(color)

	//check if it's indexed color
	for i, c := range indexedColors {
		if color == c {
			return &ml.Color{ Indexed: &i }
		}
	}

	return &ml.Color{ RGB: color }
}

//normalizeColor check if color in #RGB format and convert it into ARGB format
func normalizeColor(color string) string {
	//normalize color
	if len(color) > 1 {
		if color[0] == '#' {
			color = "FF" + color[1:]
		}
	}

	return strings.ToUpper(color)
}

func init() {
	indexedColors = []string{
		normalizeColor("#000000"),
		normalizeColor("#FFFFFF"),
		normalizeColor("#FF0000"),
		normalizeColor("#00FF00"),
		normalizeColor("#0000FF"),
		normalizeColor("#FFFF00"),
		normalizeColor("#FF00FF"),
		normalizeColor("#00FFFF"),
		normalizeColor("#000000"),
		normalizeColor("#FFFFFF"),
		normalizeColor("#FF0000"),
		normalizeColor("#00FF00"),
		normalizeColor("#0000FF"),
		normalizeColor("#FFFF00"),
		normalizeColor("#FF00FF"),
		normalizeColor("#00FFFF"),
		normalizeColor("#800000"),
		normalizeColor("#008000"),
		normalizeColor("#000080"),
		normalizeColor("#808000"),
		normalizeColor("#800080"),
		normalizeColor("#008080"),
		normalizeColor("#C0C0C0"),
		normalizeColor("#808080"),
		normalizeColor("#9999FF"),
		normalizeColor("#993366"),
		normalizeColor("#FFFFCC"),
		normalizeColor("#CCFFFF"),
		normalizeColor("#660066"),
		normalizeColor("#FF8080"),
		normalizeColor("#0066CC"),
		normalizeColor("#CCCCFF"),
		normalizeColor("#000080"),
		normalizeColor("#FF00FF"),
		normalizeColor("#FFFF00"),
		normalizeColor("#00FFFF"),
		normalizeColor("#800080"),
		normalizeColor("#800000"),
		normalizeColor("#008080"),
		normalizeColor("#0000FF"),
		normalizeColor("#00CCFF"),
		normalizeColor("#CCFFFF"),
		normalizeColor("#CCFFCC"),
		normalizeColor("#FFFF99"),
		normalizeColor("#99CCFF"),
		normalizeColor("#FF99CC"),
		normalizeColor("#CC99FF"),
		normalizeColor("#FFCC99"),
		normalizeColor("#3366FF"),
		normalizeColor("#33CCCC"),
		normalizeColor("#99CC00"),
		normalizeColor("#FFCC00"),
		normalizeColor("#FF9900"),
		normalizeColor("#FF6600"),
		normalizeColor("#666699"),
		normalizeColor("#969696"),
		normalizeColor("#003366"),
		normalizeColor("#339966"),
		normalizeColor("#003300"),
		normalizeColor("#333300"),
		normalizeColor("#993300"),
		normalizeColor("#993366"),
		normalizeColor("#333399"),
		normalizeColor("#333333"),
	}
}
