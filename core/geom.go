// Package core contains basic implementations of structs that implement geometry interfaces.
package core

import (
	"reflect"

	"github.com/go-spatial/geom"
)

// BoundingBox contains x1, y1, x2, y2
type BoundingBox struct {
	coordinates [4]float64
}

// BBox returns x1, y1, x2, y2
func (bbox BoundingBox) BBox() (float64, float64, float64, float64) {
	return bbox.coordinates[0],
		bbox.coordinates[1],
		bbox.coordinates[2],
		bbox.coordinates[3]
}

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
	points []geom.Point
	bbox   geom.BoundingBox
}

// SubPoints returns the points that compose the LineString
func (ls LineString) SubPoints() []geom.Point {
	return ls.points
}

// BBox returns x1, y1, x2, y2
func (ls LineString) BBox() (float64, float64, float64, float64) {
	if ls.bbox == nil {
		ls.bbox = MakeBBox(ls.SubPoints())
	}
	return ls.bbox.BBox()
}

// MakeBBox creates something implementing the geom.BoundingBox interface from the input
func MakeBBox(input interface{}) geom.BoundingBox {
	var (
		x1, x2, y1, y2, newx1, newx2, newy1, newy2 float64
	)

	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice:
		sl := reflect.ValueOf(input)

		for inx := 0; inx < sl.Len(); inx++ {
			if bboxer, ok := (sl.Index(inx).Interface()).(geom.BoundingBox); ok {
				newx1, newy1, newx2, newy2 = bboxer.BBox()
				if inx == 0 {
					x1 = newx1
					y1 = newy1
					x2 = newx2
					y2 = newy2
					continue
				}
				if newx1 < x1 {
					x1 = newx1
				}
				if newx2 > x2 {
					x2 = newx2
				}
				if newy1 < y1 {
					y1 = newy1
				}
				if newy2 > y2 {
					y2 = newy2
				}
			}
		}
	}
	return BoundingBox{coordinates: [4]float64{x1, y1, x2, y2}}
}
