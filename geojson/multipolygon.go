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

import (
	"encoding/json"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/core"
)

// MultiPolygon is a geometry of multiple polygons.
type MultiPolygon struct {
	Type        string          `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
	BBox        BoundingBox     `json:"bbox,omitempty"`
}

// ToMultiPolygon returns a MultiPolygon for the GeoJSON input
func ToMultiPolygon(bytes []byte) (geom.MultiPolygon, error) {
	var (
		result      geom.MultiPolygon
		gjMP        MultiPolygon
		err         error
		points      []geom.Point
		lineStrings []geom.LineString
		polygons    []geom.Polygon
	)
	if err = json.Unmarshal(bytes, &gjMP); err == nil {

		for _, coords3 := range gjMP.Coordinates {
			lineStrings = nil
			for _, coords2 := range coords3 {
				points = nil
				for _, coords := range coords2 {
					points = append(points, core.Point{X: coords[0], Y: coords[1]})
				}
				lineStrings = append(lineStrings, core.LineString{Points: points})
			}

			polygons = append(polygons, core.Polygon{LineStrings: lineStrings})
		}
		result = core.MultiPolygon{Polygons: polygons}
	}
	return result, err
}

// FromMultiPolygon returns GeoJSON for the input multipolygon
func FromMultiPolygon(input geom.MultiPolygon, options FromOptions) (string, error) {
	var (
		result       string
		err          error
		mp           MultiPolygon
		coordinates4 [][][][]float64
		coordinates3 [][][]float64
		coordinates2 [][]float64
		bytes        []byte
	)
	for _, polygon := range input.SubPolygons() {
		coordinates3 = nil
		for _, lineString := range polygon.SubLineStrings() {
			coordinates2 = nil
			for _, point := range lineString.SubPoints() {
				x, y := point.XY()
				coordinates2 = append(coordinates2, []float64{x, y})
			}
			coordinates3 = append(coordinates3, coordinates2)
		}
		coordinates4 = append(coordinates4, coordinates3)
	}
	mp = MultiPolygon{Type: MULTIPOLYGON, Coordinates: coordinates4}
	if options.BBox {
		x1, y1, x2, y2 := input.BBox()
		mp.BBox = BoundingBox{x1, y1, x2, y2}
	}
	if bytes, err = json.Marshal(mp); err == nil {
		result = string(bytes)
	}
	return result, err
}
