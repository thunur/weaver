// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/thunur/weaver"
)

//go:generate ../../cmd/weaver/weaver generate

//go:embed index.html
var indexHtml string // index.html served on "/"

func main() {
	// Initialize the Service Weaver application.
	flag.Parse()
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type server struct {
	weaver.Implements[weaver.Main]
	reverser weaver.Ref[Reverser]
	lis      weaver.Listener `weaver:"reverser"`
}

func serve(ctx context.Context, s *server) error {
	// Setup the HTTP handler.
	var mux http.ServeMux
	mux.Handle("/", weaver.InstrumentHandlerFunc("root", s.handleRoot))
	mux.Handle("/reverse", weaver.InstrumentHandlerFunc("reverse", s.handleReverse))
	mux.HandleFunc(weaver.HealthzURL, weaver.HealthzHandler)
	s.Logger(ctx).Info("Reverser server running", "address", s.lis)
	return http.Serve(s.lis, &mux)
}

// Handle requests to the "/" endpoint.
func (s *server) handleRoot(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, indexHtml)
}

// Handle requests to the "/reverse?s=<string>" endpoint.
func (s *server) handleReverse(w http.ResponseWriter, r *http.Request) {
	reversed, err := s.reverser.Get().Reverse(r.Context(), r.URL.Query().Get("s"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, html.EscapeString(reversed))
}
