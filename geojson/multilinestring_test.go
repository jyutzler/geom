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

const MLS2 = `{"type":"MultiLineString","coordinates":[[[10,20],[30,40]],[[-10,-5],[15,20]]]}`
const MLS2B = `{"type":"MultiLineString","coordinates":[[[10,20],[30,40]],[[-10,-5],[15,20]]],"bbox":[-10,-5,30,40]}`

func TestFromMultiLineString(t *testing.T) {
	var (
		mls     geom.MultiLineString
		text    string
		err     error
		options FromOptions
	)
	mls = core.MultiLineString{LineStrings: []geom.LineString{core.LineString{Points: []geom.Point{core.Point{X: 10, Y: 20},
		core.Point{X: 30, Y: 40}}},
		core.LineString{Points: []geom.Point{core.Point{X: -10, Y: -5},
			core.Point{X: 15, Y: 20}}}}}

	if text, err = FromMultiLineString(mls, options); err == nil {
		if text != MLS2 {
			t.Errorf("Expected %v, received %v.", MLS2, text)
		}
	} else {
		t.Error(err)
	}

	options.BBox = true
	if text, err = FromMultiLineString(mls, options); err == nil {
		if text != MLS2B {
			t.Errorf("Expected %v, received %v.", MLS2B, text)
		}
	} else {
		t.Error(err)
	}
}

func TestToMultiLineString(t *testing.T) {
	var (
		err   error
		mls   geom.MultiLineString
		bytes []byte
		x, y  float64
	)
	if bytes, err = ioutil.ReadFile("test/multilinestring.geojson"); err != nil {
		t.Error(err)
	}
	if mls, err = ToMultiLineString(bytes); err != nil {
		t.Error(err)
	}
	x, y = mls.SubLineStrings()[0].SubPoints()[0].XY()
	if x != 100 {
		t.Errorf("Expected x=100, received %v", x)
	}
	if y != 0 {
		t.Errorf("Expected y=0, received %v", y)
	}
}
