/*
Copyright 2024 Chainguard, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/chainguard-dev/clog"
)

func main() {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clog.InfoContextf(r.Context(), string(b))
		w.WriteHeader(http.StatusOK)
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
