// Copyright 2015 The go-sabom Authors
// This file is part of the go-sabom library.
//
// The go-sabom library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-sabom library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-sabom library. If not, see <http://www.gnu.org/licenses/>.

package flags

import (
	"os"
	"os/user"
	"testing"
)

func TestPathExpansion(t *testing.T) {
	user, _ := user.Current()
	tests := map[string]string{
		"/home/someuser/tmp": "/home/someuser/tmp",
		"~/tmp":              user.HomeDir + "/tmp",
		"~thisOtherUser/b/":  "~thisOtherUser/b",
		"$DDDXXX/a/b":        "/tmp/a/b",
		"/a/b/":              "/a/b",
	}
	os.Setenv("DDDXXX", "/tmp")
	for test, expected := range tests {
		got := expandPath(test)
		if got != expected {
			t.Errorf("test %s, got %s, expected %s\n", test, got, expected)
		}
	}
}
