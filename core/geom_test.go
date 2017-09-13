package core

import (
	"testing"

	"github.com/go-spatial/geom"
)

func TestPoint(t *testing.T) {
	var (
		point geom.Point
	)
	point = Point{X: 10, Y: 20}
	point.XY()
	point.BBox()
}

func TestLineString(t *testing.T) {
	var (
		ls     geom.LineString
		points []geom.Point
	)
	points = append(points, Point{X: 10, Y: 20}, Point{X: 30, Y: 40}, Point{X: -10, Y: -5})
	ls = LineString{Points: points}
	ls.SubPoints()
	x1, y1, x2, y2 := ls.BBox()
	if x1 != -10 {
		t.Errorf("Expected x1 = -10, received %v", x1)
	}
	if x2 != 30 {
		t.Errorf("Expected x2 = 30, received %v", x2)
	}
	if y1 != -5 {
		t.Errorf("Expected y1 = -5, received %v", y1)
	}
	if y2 != 40 {
		t.Errorf("Expected y2 = 40, received %v", y2)
	}
}
