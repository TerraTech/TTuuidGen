//==============================================================================
// Copyright (c) 2019, FutureQuest, Inc.
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
	"path/filepath"
	"strconv"

	flag "github.com/TerraTech/pflag"
)

type optionsT struct {
	hyphens, length int
	s1, s4          bool
	showVersion     bool
}

func handleFlags() *optionsT {
	var err error
	var options *optionsT = &optionsT{}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-hV] {[--s<1|4>] | [-H <int>] [length]}\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	//help := flag.Bool("h", false, "help")
	flag.IntVarP(&options.hyphens, "", "H", 0, "number of hyphens to add")
	flag.BoolVar(&options.s1, "s1", false, "emit standard uuid V1 (time based)")
	flag.BoolVar(&options.s4, "s4", false, "emit standard uuid V4 (random based)")
	flag.BoolVarP(&options.showVersion, "version", "V", false, "Show version")
	flag.Parse()

	if options.showVersion {
		fmt.Println(Version())
		os.Exit(0)
	}

	if options.s1 && options.s4 {
		fatal("-s1 and -s4 are mutually exclusive")
	}

	if (options.s1 || options.s4) && options.hyphens > 0 {
		fatal("standard uuid does not support custom hyphenation")
	}

	if options.s1 || options.s4 {
		return options
	}

	if len(flag.Args()) != 0 {
		options.length, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			fatal(err.Error())
		}
	}

	return options
}
