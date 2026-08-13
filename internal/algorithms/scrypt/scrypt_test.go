// Copyright (c) 2026 AldianOkto. All rights reserved.
// Copyright (c) 2026 Xcosh Core.
// Use of this source code is governed by the Apache License.
// that can be found in the root directory of this repository.
// Project: Xcosh / Blockchain Core
//
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at. <http://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scrypt

import (
	"testing"

	"github.com/xcosh-chain/xcosh/types"
)

func TestPoWHashDeterministic(t *testing.T) {
	h := New()
	input := []byte("test vector for scrypt pow hash")
	got1 := h.PoWHash(input)
	got2 := h.PoWHash(input)

	if got1 == types.ZeroHash {
		t.Fatal("PoWHash returned zero hash")
	}
	if got1 != got2 {
		t.Fatal("PoWHash is not deterministic")
	}
}

func TestPoWHashDifferentInputs(t *testing.T) {
	h := New()
	a := h.PoWHash([]byte("input A"))
	b := h.PoWHash([]byte("input B"))

	if a == b {
		t.Fatal("different inputs produced the same hash")
	}
}

func TestPoWHashEmptyInput(t *testing.T) {
	h := New()
	got := h.PoWHash([]byte{})
	if got == types.ZeroHash {
		t.Fatal("PoWHash of empty input returned zero hash")
	}
}

func TestName(t *testing.T) {
	h := New()
	if h.Name() != "scrypt" {
		t.Fatalf("expected name scrypt, got %s", h.Name())
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
			for j := 0; j < 20; j++ {
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
