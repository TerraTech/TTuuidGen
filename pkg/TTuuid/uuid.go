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

package TTuuid

import (
	"bytes"
	"encoding/hex"
	"math"

	goUuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
)

// DefaultLength is the standard UUID length.
const DefaultLength = 36

// TTuuid holds info for both Hyphens and Length
type TTuuid struct {
	Hyphens int
	Length  int
}

// New returns a new TTuuid object for given number of hyphens and length of resultant uuid string.
//
// length == 0, will use default: 36
func New(hyphens, length int) (*TTuuid, error) {
	if length == 0 {
		length = DefaultLength
	}

	c := length / 2
	if hyphens >= c {
		return nil, ErrTooManyHyphens(c - 1)
	}

	return &TTuuid{
		Hyphens: hyphens,
		Length:  length,
	}, nil
}

// Custom returns a mixed UUIDv1 + UUIDv4 string with custom number of hyphens and specified overall length.
func (ttuuid *TTuuid) Custom() string {
	ttuuid.Length -= ttuuid.Hyphens
	out := make([]byte, int(math.Ceil(float64(ttuuid.Length)/2.0)))
	c1 := sha3.NewCShake256([]byte("LOGUUIDGEN"), Time().Bytes())
	c1.Write(Random().Bytes())
	c1.Read(out)

	hexString := (hex.EncodeToString(out))[:ttuuid.Length]
	l := len(hexString)

	var chunk int = int(math.Floor(float64(l) / (float64(ttuuid.Hyphens) + 1.0)))

	if chunk == 0 {
		return ""
	}

	h := ttuuid.Hyphens
	if ttuuid.Hyphens > 0 && chunk < l {
		var b bytes.Buffer
		for i, j := 0, chunk; j <= l; j += chunk {
			b.WriteString(hexString[i:j])
			b.WriteString("-")
			h--
			if h == 0 || j+chunk >= l {
				b.WriteString(hexString[j:])
				break
			}
			i = j
		}
		hexString = b.String()
	}

	return hexString
}

// Random returns a standard random based UUIDv4 object.
func Random() goUuid.UUID {
	return goUuid.Must(goUuid.NewV4())
}

// Time returns a standard time based UUIDv1 object.
func Time() goUuid.UUID {
	return goUuid.Must(goUuid.NewV1())
}
