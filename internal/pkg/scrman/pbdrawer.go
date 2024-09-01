package scrman

import (
	"github.com/xeightfour/barg/internal/pkg/prober"
)

func drawNode(pb *prober.Prober, root int, top int, left int, sc *Screen) (int, int, error) {
	var err error
	right := left - 1
	cLeft, cRight := -1, -1

	// Draw children
	for _, v := range pb.Graph[root] {
		var pos int
		right, pos, err = drawNode(pb, v, top+4, right+1, sc)
		if err != nil {
			return 0, 0, err
		}
		length := len(pb.Nodes[v].Name)
		if cLeft == -1 {
			cLeft = pos + (length+1)/2
		}
		cRight = pos + (length+1)/2
	}

	// Find root's position
	length := len(pb.Nodes[root].Name)
	right = max(right, left+length+1)
	pos := (right+left-length-1)/2

	// Draw root
	if err := sc.TexBox(top, pos, pb.Nodes[root].Name); err != nil {
		return 0, 0, err
	}

	// Draw lower edges
	if len(pb.Graph[root]) > 0 {
		center := pos + (length+1)/2
		if err := sc.VLine(center, top+2, top+3); err != nil {
			return 0, 0, err
		}
		if err := sc.HLine(top+3, cLeft, center); err != nil {
			return 0, 0, err
		}
		if err := sc.HLine(top+3, center, cRight); err != nil {
			return 0, 0, err
		}
	}

	// Draw upper edges
	if pb.Nodes[root].Type != 1 {
		center := pos + (length+1)/2
		if err := sc.VLine(center, top-1, top); err != nil {
			return 0, 0, err
		}
	}

	return right, pos, nil
}

func DrawProber(pb *prober.Prober) error {
	for i, v := range pb.Nodes {
		if v.Type == 1 {
			var sc Screen
			sc.Init()
			if _, _, err := drawNode(pb, i, 0, 0, &sc); err != nil {
				return err
			}
			sc.Show()
		}
	}
	return nil
}
