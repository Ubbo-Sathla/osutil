// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package userutil

import (
	"testing"

	"github.com/tredoe/osutil"
)

func TestSudo(t *testing.T) {
	if err := MustBeSuperUser(osutil.SystemUndefined); err != nil {
		t.Error(err)
	}
}