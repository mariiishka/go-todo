package tablespretty

import (
	"github.com/alexeyco/simpletable"
	"github.com/mariiishka/go-todo/internal/lib/stringspretty"
)

const (
	borderColor = stringspretty.ColorPink
)

var (
	// StyleDefault
	//
	//
	//  # ║       NAME       ║    TAX
	// ═══║══════════════════║════════════
	//  1 ║ Newton G. Goetz  ║   $ 532.70
	//  2 ║ Rebecca R. Edney ║  $ 1423.25
	//  3 ║ John R. Jackson  ║  $ 7526.12
	//  4 ║ Ron J. Gomes     ║   $ 123.84
	//  5 ║ Penny R. Lewis   ║  $ 3221.11
	// ═══║══════════════════║════════════
	//    ║         Subtotal ║ $ 12827.02
	//
	StyleDefaultColorful *simpletable.Style = &simpletable.Style{
		Border: &simpletable.BorderStyle{
			TopLeft:            " ",
			Top:                " ",
			TopRight:           " ",
			Right:              " ",
			BottomRight:        " ",
			Bottom:             " ",
			BottomLeft:         " ",
			Left:               " ",
			TopIntersection:    " ",
			BottomIntersection: " ",
		},
		Divider: &simpletable.DividerStyle{
			Left:         " ",
			Center:       stringspretty.Color(borderColor, "═"),
			Right:        " ",
			Intersection: stringspretty.Color(borderColor, "║"),
		},
		Cell: stringspretty.Color(borderColor, "║"),
	}
)
