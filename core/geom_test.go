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

func TestLineStringAndMakeBBox(t *testing.T) {
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

func TestMultiPoint(t *testing.T) {
	var (
		mp     geom.MultiPoint
		points []geom.Point
	)
	points = append(points, Point{X: 10, Y: 20}, Point{X: 30, Y: 40}, Point{X: -10, Y: -5})
	mp = MultiPoint{Points: points}
	mp.BBox()
}

func TestPolygon(t *testing.T) {
	var (
		polygon     geom.Polygon
		lineStrings []geom.LineString
	)
	lineStrings = append(lineStrings, LineString{Points: []geom.Point{Point{X: 10, Y: 20},
		Point{X: 30, Y: 40},
		Point{X: -10, Y: -5},
		Point{X: 10, Y: 20}}})
	polygon = Polygon{LineStrings: lineStrings}
	x1, y1, x2, y2 := polygon.BBox()
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

func TestMultiLineString(t *testing.T) {
	var (
		mls         geom.MultiLineString
		lineStrings []geom.LineString
	)
	lineStrings = append(lineStrings, LineString{Points: []geom.Point{Point{X: 10, Y: 20},
		Point{X: 30, Y: 40}}})
	lineStrings = append(lineStrings, LineString{Points: []geom.Point{Point{X: -10, Y: -5},
		Point{X: 15, Y: 20}}})
	mls = MultiLineString{LineStrings: lineStrings}
	x1, y1, x2, y2 := mls.BBox()
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

func TestMultiPolygon(t *testing.T) {
	var (
		mp          geom.MultiPolygon
		lineStrings []geom.LineString
	)
	lineStrings = append(lineStrings, LineString{Points: []geom.Point{Point{X: 10, Y: 20},
		Point{X: 30, Y: 40},
		Point{X: -10, Y: -5},
		Point{X: 15, Y: 20}}})
	mp = MultiPolygon{Polygons: []geom.Polygon{Polygon{LineStrings: lineStrings}}}
	x1, y1, x2, y2 := mp.BBox()
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
