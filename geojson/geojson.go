/*
Copyright 2016, RadiantBlue Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package geojson

// GeoJSON Constants
const (
	TYPE               = "type"
	BBOX               = "bbox"
	COORDINATES        = "coordinates"
	POINT              = "Point"
	LINESTRING         = "LineString"
	POLYGON            = "Polygon"
	MULTIPOINT         = "MultiPoint"
	MULTILINESTRING    = "MultiLineString"
	MULTIPOLYGON       = "MultiPolygon"
	GEOMETRYCOLLECTION = "GeometryCollection"
	GEOMETRIES         = "geometries"
)

// FromOptions contains the various options for the From* functions
type FromOptions struct {
	BBox bool
}
