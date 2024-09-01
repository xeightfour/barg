package scrman

import (
	"fmt"
)

const (
	Width  = 65
	Height = 35
)

type (
	Screen struct {
		Table [Height][Width]Glyph
	}
)

func (sc *Screen) setAlpha(i int, j int, x rune) {
	sc.Table[i][j].SetAlpha(x)
}

func (sc *Screen) setMask(i int, j int, x int) {
	sc.Table[i][j].SetMask(x)
}

func (sc *Screen) HLine(row int, from int, to int) error {
	if min(row, from, to) < 0 || row >= Height || max(from, to) >= Width {
		return fmt.Errorf("screen.go:HLine: Line doesn't fit within screen")
	}
	if from == to {
		return nil
	}
	if from > to {
		from, to = to, from
	}
	sc.setMask(row, from, 2)
	sc.setMask(row, to, 8)
	for i := from + 1; i < to; i++ {
		sc.setMask(row, i, 10)
	}
	return nil
}

func (sc *Screen) VLine(col int, from int, to int) error {
	if min(col, from, to) < 0 || col >= Width || max(from, to) >= Height {
		return fmt.Errorf("screen.go:VLine: Line doesn't fit within screen")
	}
	if from == to {
		return nil
	}
	if from > to {
		from, to = to, from
	}
	sc.setMask(from, col, 4)
	sc.setMask(to, col, 1)
	for i := from + 1; i < to; i++ {
		sc.setMask(i, col, 5)
	}
	return nil
}

func (sc *Screen) Box(sy int, ey int, sx int, ex int) error {
	if min(sy, ey, sx, ex) < 0 || max(sy, ey) >= Height || max(sx, ex) >= Width {
		return fmt.Errorf("screen.go:Box: Box doesn't fit within screen")
	}
	sc.HLine(sy, sx, ex)
	sc.HLine(ey, sx, ex)
	sc.VLine(sx, sy, ey)
	sc.VLine(ex, sy, ey)
	return nil
}

func (sc *Screen) Text(sy int, sx int, str string) error {
	ex := sx + len(str) - 1
	if min(sy, sx, ex) < 0 || sy >= Height || max(sx, ex) >= Width {
		return fmt.Errorf("screen.go:Text: Text doesn't fit within screen")
	}
	for i := 0; i < len(str); i++ {
		sc.setAlpha(sy, sx+i, rune(str[i]))
	}
	return nil
}

func (sc *Screen) TexBox(sy int, sx int, str string) error {
	ex := sx + len(str) + 1
	ey := sy + 2
	if min(sy, ey, sx, ex) < 0 || max(sy, ey) >= Height || max(sx, ex) >= Width {
		return fmt.Errorf("screen.go:TexBox: Text box doesn't fit within screen")
	}
	if err := sc.Box(sy, ey, sx, ex); err != nil {
		return err
	}
	if err := sc.Text(sy+1, sx+1, str); err != nil {
		return err
	}
	return nil
}

func (sc *Screen) Show() {
	var table [Height][Width]rune
	for i := range table {
		for j := 0; j < Width; j++ {
			table[i][j] = sc.Table[i][j].Get()
		}
	}
	for _, s := range table {
		last := -1
		for i, c := range s {
			if c != ' ' {
				last = i + 1
			}
		}
		if last != -1 {
			str := string(s[:last])
			fmt.Println(str)
		}
	}
}

func (sc *Screen) Init() {
	for i := range sc.Table {
		for j := 0; j < Width; j++ {
			sc.Table[i][j].SetAlpha(' ')
		}
	}
}
