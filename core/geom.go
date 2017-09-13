// Package core contains basic implementations of structs that implement geometry interfaces.
package core

import (
	"github.com/go-spatial/geom"
)

// The Point object contains a single position
type Point struct {
	X float64
	Y float64
}

// XY returns the X and Y values in a format suitable for certain functions
func (p Point) XY() (float64, float64) {
	return p.X, p.Y
}

// BBox returns x1, y1, x2, y2
func (p Point) BBox() (float64, float64, float64, float64) {
	return p.X, p.Y, p.X, p.Y
}

// LineString is a line of two or more points
type LineString struct {
	Points []geom.Point
	bbox   geom.BoundingBox
}

// SubPoints returns the points that compose the LineString
func (ls LineString) SubPoints() []geom.Point {
	return ls.Points
}

// BBox returns x1, y1, x2, y2
func (ls LineString) BBox() (float64, float64, float64, float64) {
	if ls.bbox == nil {
		ls.bbox = MakeBBox(ls.SubPoints())
	}
	return ls.bbox.BBox()
}

// MultiPoint is a geometry with multiple points
type MultiPoint struct {
	Points []geom.Point
	bbox   geom.BoundingBox
}

// SubPoints returns the points that compose the MultiPoint
func (mp MultiPoint) SubPoints() []geom.Point {
	return mp.Points
}

// BBox returns x1, y1, x2, y2
func (mp MultiPoint) BBox() (float64, float64, float64, float64) {
	if mp.bbox == nil {
		mp.bbox = MakeBBox(mp.SubPoints())
	}
	return mp.bbox.BBox()
}
