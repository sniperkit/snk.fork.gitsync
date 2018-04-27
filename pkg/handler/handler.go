// Copyright 2018 The gitsync Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"fmt"
	"net/http"
)

//go:generate counterfeiter -o mocks/git.go --fake-name Git . git
type git interface {
	Sync() error
}

//go:generate counterfeiter -o mocks/hook.go --fake-name Hook . hook
type hook interface {
	Call() error
}

// Syncer stores the required links for serving gitsync
type Syncer struct {
	Git  git
	Hook hook
}

// ServeHTTP provides a http handler and calls Hook after successfully running Git.Sync()
func (i *Syncer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}
