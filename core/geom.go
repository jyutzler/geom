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

// XYs returns an array of points, useful for making bounding boxes
func (p Point) XYs() [][2]float64 {
	return [][2]float64{[2]float64{p.X, p.Y}}
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

// XYs returns coordinate pairs
func (ls LineString) XYs() [][2]float64 {
	var result [][2]float64
	for _, point := range ls.SubPoints() {
		result = append(result, point.XYs()...)
	}
	return result
}

// BBox returns x1, y1, x2, y2
func (ls LineString) BBox() (float64, float64, float64, float64) {
	if ls.bbox == nil {
		ls.bbox = MakeBBox(ls.XYs())
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

// XYs returns coordinate pairs
func (mp MultiPoint) XYs() [][2]float64 {
	var result [][2]float64
	for _, point := range mp.SubPoints() {
		result = append(result, point.XYs()...)
	}
	return result
}

// BBox returns x1, y1, x2, y2
func (mp MultiPoint) BBox() (float64, float64, float64, float64) {
	if mp.bbox == nil {
		mp.bbox = MakeBBox(mp.XYs())
	}
	return mp.bbox.BBox()
}

// Polygon is a geometry consisting of multiple closed LineStrings. There must be only one exterior LineString with a clockwise winding order. There may be one or more interior LineStrings with a counterclockwise winding orders.
type Polygon struct {
	LineStrings []geom.LineString
	bbox        geom.BoundingBox
}

// SubLineStrings returns the LineStrings that compose the Polygon
func (polygon Polygon) SubLineStrings() []geom.LineString {
	return polygon.LineStrings
}

// XYs returns coordinate pairs
func (polygon Polygon) XYs() [][2]float64 {
	var result [][2]float64
	for _, lineString := range polygon.SubLineStrings() {
		result = append(result, lineString.XYs()...)
	}
	return result
}

// BBox returns x1, y1, x2, y2
func (polygon Polygon) BBox() (float64, float64, float64, float64) {
	if polygon.bbox == nil {
		polygon.bbox = MakeBBox(polygon.XYs())
	}
	return polygon.bbox.BBox()
}

// MultiLineString is a geometry with multiple LineStrings.
type MultiLineString struct {
	LineStrings []geom.LineString
	bbox        geom.BoundingBox
}

// SubLineStrings returns the LineStrings that compose the Polygon
func (mls MultiLineString) SubLineStrings() []geom.LineString {
	return mls.LineStrings
}

// XYs returns coordinate pairs
func (mls MultiLineString) XYs() [][2]float64 {
	var result [][2]float64
	for _, lineString := range mls.SubLineStrings() {
		result = append(result, lineString.XYs()...)
	}
	return result
}

// BBox returns x1, y1, x2, y2
func (mls MultiLineString) BBox() (float64, float64, float64, float64) {
	if mls.bbox == nil {
		mls.bbox = MakeBBox(mls.XYs())
	}
	return mls.bbox.BBox()
}

// MultiPolygon is a geometry of multiple polygons.
type MultiPolygon struct {
	Polygons []geom.Polygon
	bbox     geom.BoundingBox
}

// SubPolygons returns the Polygons that compose the MultiPolygon
func (mp MultiPolygon) SubPolygons() []geom.Polygon {
	return mp.Polygons
}

// XYs returns coordinate pairs
func (mp MultiPolygon) XYs() [][2]float64 {
	var result [][2]float64
	for _, polygon := range mp.SubPolygons() {
		result = append(result, polygon.XYs()...)
	}
	return result
}

// BBox returns x1, y1, x2, y2
func (mp MultiPolygon) BBox() (float64, float64, float64, float64) {
	if mp.bbox == nil {
		mp.bbox = MakeBBox(mp.XYs())
	}
	return mp.bbox.BBox()
}
