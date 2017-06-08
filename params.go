// Copyright 2017 Jave Tann. All rights reserved.
// Based on the copyright 2013 Julien Schmidt. All rights reserved.
// Based on the path package, Copyright 2009 The Go Authors.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gomux

/*
 Temporary use of gorilla/context until Go1.7 is out
*/

import (
	"context"
	"net/http"
)

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

type key int

const paramsKey key = 0

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

// GetMuxParams returns the mux params
func GetMuxParams(r *http.Request) *Params {
	if ps := r.Context().Value(paramsKey); ps != nil {
		return ps.(*Params)
	}
	return &Params{}
}

// setMuxParams saves mux params to the context
func setMuxParams(r *http.Request, ps *Params) *http.Request {
	ctx := context.WithValue(r.Context(), paramsKey, ps)
	return r.WithContext(ctx)
}
