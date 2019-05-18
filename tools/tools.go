package tools

import (
	"ExchangeRatesGraphic/collors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

const W = 1000
const H = 500

func PaintCurrencyScheme(paramsLst map[string][]float64) string {
	painter := createGraph(W, H)

	collor := []func(*gg.Context) *gg.Context{collors.Red, collors.Blue, collors.Green, collors.Purple, collors.Cyan}
	i := 0
	for key, valuesLst := range paramsLst {
		painter = collor[i](painter)
		i++
		drawCurrency(painter, valuesLst, key, i)
	}

	painter.SavePNG("out.png")
	return "Complete"
}

func drawCurrency(dc *gg.Context, currencyValues []float64, name string, namePos int) {
	dc.SetLineWidth(5)
	lenght := len(currencyValues) + 1
	pointBlock := float64(W / lenght)
	piece := float64(W / 6)
	zeroX := float64(W / 60)
	zeroY := float64(H / 1.05)
	for i := range currencyValues {
		lenght -= 1
		endPointX := zeroX + pointBlock
		endPointY := H - float64(currencyValues[i])*2.5

		dc.DrawLine(zeroX, zeroY, endPointX, endPointY)
		fmt.Println(endPointY)
		dc.DrawCircle(endPointX, endPointY, 5)

		strValue := fmt.Sprintf("%f", currencyValues[i])
		paintText(dc, strValue, endPointX, endPointY-H/50, 10)

		dc.Stroke()
		zeroX = endPointX
		zeroY = endPointY
	}

	posX := piece * float64(namePos)
	paintText(dc, name, posX, H/20, 25)
	dc.DrawCircle(posX-(W/50), H/30, 10)
	dc.Stroke()
}

func paintText(dc *gg.Context, text string, x float64, y float64, s float64) {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: s,
	})
	dc.SetFontFace(face)
	dc.DrawString(text, x, y)
	dc.Stroke()
}

func createGraph(width int, height int) *gg.Context {
	//Background main lines
	painter := gg.NewContext(width, height)
	painter.SetRGB(1, 1, 1)
	painter.Clear()
	painter.DrawLine(W/60, 0, W/60, H)
	painter.DrawLine(0, H/1.05, W, H/1.05)
	painter.SetRGBA(0, 0, 0, 0.5)
	painter.SetLineWidth(5)
	painter.Stroke()

	//Background opt lines
	painter.SetRGBA(0, 0, 0, 0.01)
	for i := 1; i <= 10; i++ {
		num := float64(i) * 50
		painter.DrawLine(W/60, H-num, W, H-num)
	}

	painter.SetRGBA(0, 0, 0, 0.25)
	painter.SetLineWidth(2)
	painter.Stroke()

	painter.SetRGBA(0, 0, 1, 0.25)
	painter.DrawPoint(W/60, H/1.05, 4)
	painter.Stroke()
	painter.SetRGB(0, 0, 0)
	painter.Stroke()

	return painter
}
