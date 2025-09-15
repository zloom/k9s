package tchart_test

import (
	"testing"

	"github.com/zloom/k9s/internal/tchart"
	"github.com/derailed/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestCoSeriesColorNames(t *testing.T) {
	c := tchart.NewComponent("fred")

	c.SetSeriesColors(tcell.ColorGreen, tcell.ColorBlue, tcell.ColorRed)

	assert.Equal(t, []string{"green", "blue", "red"}, c.GetSeriesColorNames())
}
