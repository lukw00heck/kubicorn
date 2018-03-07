// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build go1.8

package http

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
	ocgoogle "go.opencensus.io/plugin/ochttp/propagation/google"
)

func addOCTransport(trans http.RoundTripper) http.RoundTripper {
	return &ochttp.Transport{
		Base: trans,
		// TODO(ramonza): enable stats after census-instrumentation/opencensus-go#302
		NoStats:     true,
		Propagation: &ocgoogle.HTTPFormat{},
	}
}
