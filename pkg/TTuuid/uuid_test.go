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

package TTuuid_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/TerraTech/TTuuidGen/pkg/TTuuid"

	"github.com/stretchr/testify/assert"
)

var testvv bool

func init() {
	if os.Getenv("GOTESTVV") != "" {
		testvv = true
	}
}

func nttError(t *testing.T, h, l int) {
	_, err := TTuuid.New(h, l)
	if assert.Error(t, err) {
		assert.Equal(t, err.(TTuuid.ErrTooManyHyphens), err)
	}
}

func ntt(t *testing.T, h, l int) *TTuuid.TTuuid {
	tt, err := TTuuid.New(h, l)
	if !assert.NoError(t, err) {
		t.Fatal(err)
		return nil
	}

	return tt
}

func testUuidCustomHL(t *testing.T, h, l int) {
	a := ntt(t, h, l).Custom()
	if testvv && testing.Verbose() {
		fmt.Println(a)
	}
	assert.Len(t, a, l)
	assert.Equal(t, h, strings.Count(a, "-"))
}

func testUuidCustomL(t *testing.T, l int) {
	assert.Len(t, ntt(t, 0, l).Custom(), l)
}

func Test_uuidTime(t *testing.T) {
	assert.NotPanics(t, func() { TTuuid.Time() })
}

func Test_uuidRandom(t *testing.T) {
	assert.NotPanics(t, func() { TTuuid.Random() })
}

func Test_uuidCustom_Len50(t *testing.T) {
	testUuidCustomL(t, 50)
}

func Test_uuidCustom_Len100(t *testing.T) {
	testUuidCustomL(t, 100)
}

func Test_uuidCustom_Len1000(t *testing.T) {
	testUuidCustomL(t, 1000)
}

func Test_uuidCustom_Len1000000(t *testing.T) {
	testUuidCustomL(t, 1000000)
}

func Test_uuidCustom_Len9_H2(t *testing.T) {
	testUuidCustomHL(t, 2, 9)
}

func Test_uuidCustom_Len9_H4_ERROR(t *testing.T) {
	nttError(t, 4, 9)
}

func Test_uuidCustom_Len10_H2(t *testing.T) {
	testUuidCustomHL(t, 2, 10)
}

func Test_uuidCustom_Len10_H4(t *testing.T) {
	testUuidCustomHL(t, 4, 10)
}

func Test_uuidCustom_Len10_H5_ERROR(t *testing.T) {
	nttError(t, 5, 10)
}

func Test_uuidCustom_Len50_H4(t *testing.T) {
	testUuidCustomHL(t, 4, 50)
}

func Test_uuidCustom_Len100_H40(t *testing.T) {
	testUuidCustomHL(t, 40, 100)
}

func Test_uuidCustom_Len100_H50_ERROR(t *testing.T) {
	nttError(t, 50, 100)
}

func Test_uuidCustom_Len1000_H49(t *testing.T) {
	testUuidCustomHL(t, 49, 1000)
}

func Test_uuidCustom_Len1000_H50(t *testing.T) {
	testUuidCustomHL(t, 50, 1000)
}
