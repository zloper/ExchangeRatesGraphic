package collors

import "github.com/fogleman/gg"

func Red(painter *gg.Context) *gg.Context {
	painter.SetRGBA(1, 0, 0, 0.7)
	return painter
}

func Blue(painter *gg.Context) *gg.Context {
	painter.SetRGBA(0, 0, 1, 0.7)
	return painter
}

func Green(painter *gg.Context) *gg.Context {
	painter.SetRGBA(0, 1, 0, 0.7)
	return painter
}

func Cyan(painter *gg.Context) *gg.Context {
	painter.SetRGBA(0, 1, 1, 0.7)
	return painter
}

func Purple(painter *gg.Context) *gg.Context {
	painter.SetRGBA(1, 0, 1, 0.7)
	return painter
}
