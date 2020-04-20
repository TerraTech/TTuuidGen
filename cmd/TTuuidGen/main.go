//==============================================================================
// Copyright (c) 2019-2020, FutureQuest, Inc.
//   https://www.FutureQuest.net
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
//==============================================================================

package main

import (
	"fmt"
	"os"

	"github.com/TerraTech/TTuuidGen/pkg/TTuuid"
)

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "[FATAL]", msg)
	os.Exit(1)
}

func main() {
	options := handleFlags()

	if options.s1 {
		fmt.Println(TTuuid.Time())
		os.Exit(0)
	}

	if options.s4 {
		fmt.Println(TTuuid.Random())
		os.Exit(0)
	}

	ttuuid, err := TTuuid.New(options.hyphens, options.length)
	if err != nil {
		fatal(err.Error())
	}

	fmt.Println(ttuuid.Custom())
}
