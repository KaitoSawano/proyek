// Copyright (c) 2026 AldianOkto. All rights reserved.
// Copyright (c) 2026 Eterbit Core.
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

package crypto_test

import (
	"fmt"
	"testing"
	"github.com/xcosh-chain/xcosh/crypto"
)

func TestCheckTarget(t *testing.T) {
	// Testnet InitialBits (see params.Testnet)
	bits := uint32(0x2000ea44)
	target := crypto.CompactToHash(bits)
	fmt.Printf("Bits 0x%08x -> target: %s\n", bits, target)
	
	// MinBits for testnet
	minBits := uint32(0x207fffff)
	minTarget := crypto.CompactToHash(minBits)
	fmt.Printf("MinBits 0x%08x -> target: %s\n", minBits, minTarget)
	
	// Check various bits values
	for _, b := range []uint32{0x1d00ffff, 0x1e00ffff, 0x1f00ffff, 0x1d0fffff, 0x1e0fffff, 0x1f0fffff, 0x1e03ffff, 0x1f03ffff, 0x2003ffff, 0x1f07ffff, 0x1e07ffff} {
		tgt := crypto.CompactToHash(b)
		fmt.Printf("Bits 0x%08x -> target: %s\n", b, tgt)
	}
	
	// Also check what the big int looks like
	big := crypto.CompactToBig(0x1f03ffff)
	fmt.Printf("\nBits 0x1f03ffff -> big: %s\n", big.Text(16))
	big2 := crypto.CompactToBig(0x1f07ffff)
	fmt.Printf("Bits 0x1f07ffff -> big: %s\n", big2.Text(16))
}
