// Copyright 2017 The rkt Authors
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

// +build linux

package netinfo

import (
	"encoding/json"
	"os"
	"syscall"
)

func LoadAt(cdirfd int) ([]NetInfo, error) {
	fd, err := syscall.Openat(cdirfd, filename, syscall.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	f := os.NewFile(uintptr(fd), filename)

	var info []NetInfo
	err = json.NewDecoder(f).Decode(&info)
	return info, err
}
