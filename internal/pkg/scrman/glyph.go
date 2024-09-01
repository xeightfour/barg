package scrman

type (
	Glyph int64
)

var (
	char = [15]rune{'╵', '╶', '╰', '╷', '│', '╭', '├', '╴', '╯', '─', '┴', '╮', '┤', '┬', '┼'}
)

func (gl *Glyph) SetAlpha(x rune) {
	*gl = Glyph(x)
}

func (gl *Glyph) SetMask(x int) {
	*gl |= Glyph((15 & x) << 32)
}

func (gl Glyph) Get() rune {
	if gl < (1 << 32) {
		return rune(gl)
	}
	cut := gl >> 32
	if cut < 16 && cut > 0 {
		return char[cut-1]
	}
	return ' '
}
