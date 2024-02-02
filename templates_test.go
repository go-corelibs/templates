// Copyright (c) 2024  The Go-Enjin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

import (
	htmlTemplate "html/template"
	"io"
	"testing"
	textTemplate "text/template"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("AddHtmlParseTree", t, func() {
		var err error
		ht0 := htmlTemplate.New("initial")
		ht0, err = ht0.Parse(`{{ printf "initial test" }}`)
		So(err, ShouldBeNil)
		ht1 := htmlTemplate.New("another")
		ht1, err = ht1.Parse(`{{ printf "another test" }}`)
		So(err, ShouldBeNil)
		So(len(ht0.Tree.Root.Nodes), ShouldEqual, 1)
		err = AddParseTree(ht1, ht0)
		So(err, ShouldBeNil)
		So(ht0.Lookup("another"), ShouldNotBeNil)
		ht2 := htmlTemplate.New("borken")
		ht2, err = ht2.Parse(`{{ printf "borken test" }}`)
		So(err, ShouldBeNil)
		err = ht0.Execute(io.Discard, map[string]interface{}{})
		So(err, ShouldBeNil)
		err = AddParseTree(ht2, ht0)
		So(err, ShouldNotBeNil)
	})

	Convey("LookupHtmlTemplate", t, func() {
		var err error
		ht0 := htmlTemplate.New("initial")
		ht0, err = ht0.Parse(`{{ printf "initial test" }}`)
		So(err, ShouldBeNil)
		ht1 := htmlTemplate.New("another")
		ht1, err = ht1.Parse(`{{ printf "another test" }}`)
		So(err, ShouldBeNil)
		So(len(ht0.Tree.Root.Nodes), ShouldEqual, 1)
		err = AddParseTree(ht1, ht0)
		So(err, ShouldBeNil)
		So(Lookup(ht0, "nope", "not-a-thing"), ShouldBeNil)
		So(Lookup(ht0, "nope", "another"), ShouldNotBeNil)
	})

	Convey("AddTextParseTree", t, func() {
		var err error
		ht0 := textTemplate.New("initial")
		ht0, err = ht0.Parse(`{{ printf "initial test" }}`)
		So(err, ShouldBeNil)
		ht1 := textTemplate.New("another")
		ht1, err = ht1.Parse(`{{ printf "another test" }}`)
		So(err, ShouldBeNil)
		So(len(ht0.Tree.Root.Nodes), ShouldEqual, 1)
		err = AddParseTree(ht1, ht0)
		So(err, ShouldBeNil)
		So(ht0.Lookup("another"), ShouldNotBeNil)
		ht2 := textTemplate.New("borken")
		ht2, err = ht2.Parse(`{{ printf "borken test" }}`)
		So(err, ShouldBeNil)
		err = ht0.Execute(io.Discard, map[string]interface{}{})
		So(err, ShouldBeNil)
		err = AddParseTree(ht2, ht0)
		So(err, ShouldBeNil)
	})

	Convey("LookupTextTemplate", t, func() {
		var err error
		ht0 := textTemplate.New("initial")
		ht0, err = ht0.Parse(`{{ printf "initial test" }}`)
		So(err, ShouldBeNil)
		ht1 := textTemplate.New("another")
		ht1, err = ht1.Parse(`{{ printf "another test" }}`)
		So(err, ShouldBeNil)
		So(len(ht0.Tree.Root.Nodes), ShouldEqual, 1)
		err = AddParseTree(ht1, ht0)
		So(err, ShouldBeNil)
		So(Lookup(ht0, "nope", "not-a-thing"), ShouldBeNil)
		So(Lookup(ht0, "nope", "another"), ShouldNotBeNil)
	})
}
