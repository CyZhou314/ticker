package watchlist

import (
	"fmt"
	"strconv"
	"strings"
	"ticker-tape/internal/quote"
	"ticker-tape/internal/ui/util"

	"github.com/muesli/reflow/ansi"
)

var (
	styleNeutral       = util.NewStyle("#d4d4d4", "", false)
	styleNeutralBold   = util.NewStyle("#d4d4d4", "", true)
	styleNeutralFaded  = util.NewStyle("#555761", "", false)
	stylePricePositive = util.NewStyle("#d1ff82", "", false)
	stylePriceNegative = util.NewStyle("#ff8c82", "", false)
)

const (
	footerHeight = 1
)

type Model struct {
	Width  int
	Quotes []quote.Quote
}

// NewModel returns a model with default values.
func NewModel() Model {
	return Model{
		Width: 100,
	}
}

func (m Model) View() string {
	return watchlist(m.Quotes, m.Width)
}

func watchlist(q []quote.Quote, elementWidth int) string {
	quoteSummaries := ""
	for _, quote := range q {
		quoteSummaries = quoteSummaries + "\n" + quoteSummary(quote, elementWidth)
	}
	return quoteSummaries
}

func quoteSummary(q quote.Quote, elementWidth int) string {

	firstLine := lineWithGap(
		styleNeutralBold(q.Symbol),
		styleNeutral(convertFloatToString(q.RegularMarketPrice)),
		elementWidth,
	)
	secondLine := lineWithGap(
		styleNeutralFaded(q.ShortName),
		priceText(q.RegularMarketChange, q.RegularMarketChangePercent),
		elementWidth,
	)

	return fmt.Sprintf("%s\n%s", firstLine, secondLine)
}

func priceText(change float64, changePercent float64) string {
	if change == 0.0 {
		return styleNeutral("  " + convertFloatToString(change) + "  (" + convertFloatToString(changePercent) + "%)")
	}

	if change > 0.0 {
		return stylePricePositive("↑ " + convertFloatToString(change) + "  (" + convertFloatToString(changePercent) + "%)")
	}

	return stylePriceNegative("↓ " + convertFloatToString(change) + " (" + convertFloatToString(changePercent) + "%)")
}

// util
func lineWithGap(leftText string, rightText string, elementWidth int) string {
	innerGapWidth := elementWidth - ansi.PrintableRuneWidth(leftText) - ansi.PrintableRuneWidth(rightText)
	if innerGapWidth > 0 {
		return leftText + strings.Repeat(" ", innerGapWidth) + rightText
	}

	return leftText + " " + rightText
}

func convertFloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}