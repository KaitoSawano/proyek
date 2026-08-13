// Copyright (c) 2026 AldianOkto. All rights reserved.
// Copyright (c) 2026 Xcosh Core.
// Use of this source code is governed by the Apache License.
// that can be found in the root directory of this repository.
// Project: Eterbit / Blockchain Core
//
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package sha256d

import (
	"encoding/hex"
	"testing"

	"github.com/xcosh-chain/xcosh/types"
)

func TestPoWHashMatchesDoubleSHA256(t *testing.T) {
	h := New()

	input := []byte("test vector for sha256d pow hash")
	got := h.PoWHash(input)

	if got == types.ZeroHash {
		t.Fatal("PoWHash returned zero hash")
	}

	got2 := h.PoWHash(input)
	if got != got2 {
		t.Fatal("PoWHash is not deterministic")
	}
}

func TestPoWHashKnownVector(t *testing.T) {
	h := New()

	// SHA256("") = e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	// SHA256(above bytes) = 5df6e0e2...9456 (BE)
	// PoWHash returns LE (reversed): 56944c5d...f65d
	input := []byte{}
	got := h.PoWHash(input)

	expected, _ := hex.DecodeString("56944c5d3f98413ef45cf54545538103cc9f298e0575820ad3591376e2e0f65d")
	var want types.Hash
	copy(want[:], expected)

	if got != want {
		t.Fatalf("PoWHash empty input:\n  got  %s\n  want %s", got, want)
	}
}

func TestName(t *testing.T) {
	h := New()
	if h.Name() != "sha256d" {
		t.Fatalf("expected name sha256d, got %s", h.Name())
	}
}

func TestConcurrentSafety(t *testing.T) {
	h := New()
	input := []byte("concurrent test data")
	expected := h.PoWHash(input)

	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- struct{}{} }()
			for j := 0; j < 100; j++ {
				got := h.PoWHash(input)
				if got != expected {
					t.Errorf("concurrent PoWHash mismatch")
					return
				}
			}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
