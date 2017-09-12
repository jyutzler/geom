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
	"testing"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/core"
)

// func TestRTPoint(t *testing.T) {
// 	var (
// 		gj     interface{}
// 		err    error
// 		m      map[string]interface{}
// 		p1     *Point
// 		p2     *Point
// 		result = `{"type":"Point","coordinates":[100,0]}`
// 	)
// 	if gj, err = ParseFile("test/point.geojson"); err != nil {
// 		t.Errorf("Failed to parse file: %v", err)
// 	}
// 	p1 = gj.(*Point)
// 	m = p1.Map()
// 	p2 = FromMap(m).(*Point)
// 	if p2.String() != result {
// 		t.Errorf("Round trip point failed: %v", p2.String())
// 	}
// }

const POINT2 = `{"type":"Point","coordinates":[10,20]}`
const POINT2B = `{"type":"Point","coordinates":[10,20],"bbox":[10,20,10,20]}`

func TestFromPoint(t *testing.T) {
	var (
		point geom.Point
		text  string
		err   error
		fpo   FromPointOptions
	)
	point = core.Point{X: 10, Y: 20}

	if text, err = FromPoint(point, fpo); err == nil {
		if text != POINT2 {
			t.Errorf("Expected %v, received %v.", POINT2, text)
		}
	} else {
		t.Error(err)
	}

	fpo.BBox = true
	if text, err = FromPoint(point, fpo); err == nil {
		if text != POINT2B {
			t.Errorf("Expected %v, received %v.", POINT2B, text)
		}
	} else {
		t.Error(err)
	}
}
