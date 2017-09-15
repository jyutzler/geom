// Package geom describes geometry interfaces.
package geom

// BoundingBox describes the extents of a geometry
type BoundingBox interface {
	BBox() (float64, float64, float64, float64)
}

// Geometry describes a geometry.
type Geometry interface {
	BoundingBox
	XYs() [][2]float64
}

// Point is a point with two dimensions.
type Point interface {
	Geometry
	XY() (float64, float64)
}

// MultiPoint is a geometry with multiple points.
type MultiPoint interface {
	Geometry
	SubPoints() []Point
}

// LineString is a line of two or more points
type LineString interface {
	Geometry
	SubPoints() []Point
}

// MultiLineString is a geometry with multiple LineStrings.
type MultiLineString interface {
	Geometry
	SubLineStrings() []LineString
}

// Polygon is a geometry consisting of multiple closed LineStrings. There must be only one exterior LineString with a clockwise winding order. There may be one or more interior LineStrings with a counterclockwise winding orders.
type Polygon interface {
	Geometry
	SubLineStrings() []LineString
}

// MultiPolygon is a geometry of multiple polygons.
type MultiPolygon interface {
	Geometry
	SubPolygons() []Polygon
}

// Collection is a collections of different geometries.
type Collection interface {
	Geometry
	Geometries() []Geometry
}
