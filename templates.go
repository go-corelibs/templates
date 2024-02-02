// Copyright (c) 2024  The Go-CoreLibs Authors
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

// Package templates provides text and html template utilities
package templates

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"
	"text/template/parse"
)

// Type is a generic type constraint for the [text/template.Template] and
// [html/template.Template] concrete types
type Type interface {
	*textTemplate.Template | *htmlTemplate.Template
}

// Template is a generic type interface for interacting with either text or
// html template instances
type Template[V Type] interface {
	Name() string
	Templates() []V
	Lookup(name string) (t V)
	AddParseTree(name string, tree *parse.Tree) (t V, err error)
}

// GetParseTree returns the [text/template/parse.Tree] associated with the
// given template
func GetParseTree[T Type](t T) (tree *parse.Tree) {
	switch tt := interface{}(t).(type) {
	case *textTemplate.Template:
		tree = tt.Tree
	case *htmlTemplate.Template:
		tree = tt.Tree
	}
	return
}

// AddParseTree adds all the individual `src.Templates()` and adds their parse
// trees to the `dst` template instance
func AddParseTree[V Type, T Template[V]](src, dst T) (err error) {
	for _, srcTmpl := range src.Templates() {
		if tt, ok := interface{}(srcTmpl).(Template[V]); ok {
			name := tt.Name()
			tree := GetParseTree(srcTmpl)
			if _, err = dst.AddParseTree(name, tree); err != nil {
				err = fmt.Errorf("error adding %v parse tree to %v template: %v", name, dst.Name(), err)
				return
			}
		}
	}
	return
}

// Lookup returns the first template matching one of the names given
func Lookup[T Type](t T, names ...string) (tmpl T) {
	if tt, ok := interface{}(t).(Template[T]); ok {
		for _, name := range names {
			if tmpl = tt.Lookup(name); tmpl != nil {
				return
			}
		}
	}
	return
}
