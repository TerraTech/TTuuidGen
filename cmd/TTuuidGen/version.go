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

// +build tools

package main

import (
	_ "github.com/TerraTech/FQversion"
	_ "github.com/TerraTech/FQversion/tools/genVersion"
)

//go:generate go run -mod=readonly -tags=generate github.com/TerraTech/FQversion/tools/genVersion -package "$GOPACKAGE" -prog "$PROG" -version "$VERSION" -build "$BUILD" -import_FQversion "$IMPFQVERSION" -lib "$LIB"
