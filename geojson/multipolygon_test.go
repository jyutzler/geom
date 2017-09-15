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
	"io/ioutil"
	"testing"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/core"
)

const MPOLY2 = `{"type":"MultiPolygon","coordinates":[[[[10,20],[30,40],[-10,-5],[15,20]]]]}`
const MPOLY2B = `{"type":"MultiPolygon","coordinates":[[[[10,20],[30,40],[-10,-5],[15,20]]]],"bbox":[-10,-5,30,40]}`

func TestFromMultiPolygon(t *testing.T) {
	var (
		mp      geom.MultiPolygon
		text    string
		err     error
		options FromOptions
	)
	mp = core.MultiPolygon{Polygons: []geom.Polygon{core.Polygon{LineStrings: []geom.LineString{core.LineString{Points: []geom.Point{core.Point{X: 10, Y: 20},
		core.Point{X: 30, Y: 40},
		core.Point{X: -10, Y: -5},
		core.Point{X: 15, Y: 20}}}}}}}

	if text, err = FromMultiPolygon(mp, options); err == nil {
		if text != MPOLY2 {
			t.Errorf("Expected %v, received %v.", MPOLY2, text)
		}
	} else {
		t.Error(err)
	}

	options.BBox = true
	if text, err = FromMultiPolygon(mp, options); err == nil {
		if text != MPOLY2B {
			t.Errorf("Expected %v, received %v.", MPOLY2B, text)
		}
	} else {
		t.Error(err)
	}
}

func TestToMultiPolygon(t *testing.T) {
	var (
		err   error
		mp    geom.MultiPolygon
		bytes []byte
		x, y  float64
	)
	if bytes, err = ioutil.ReadFile("test/multipolygon.geojson"); err != nil {
		t.Error(err)
	}
	if mp, err = ToMultiPolygon(bytes); err != nil {
		t.Error(err)
	}
	x, y = mp.SubPolygons()[0].SubLineStrings()[0].SubPoints()[0].XY()
	if x != 102 {
		t.Errorf("Expected x=102, received %v", x)
	}
	if y != 2 {
		t.Errorf("Expected y=2, received %v", y)
	}
	x, y = mp.SubPolygons()[1].SubLineStrings()[0].SubPoints()[0].XY()
	if x != 100 {
		t.Errorf("Expected x=100, received %v", x)
	}
	if y != 0 {
		t.Errorf("Expected y=0, received %v", y)
	}
}
