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
	"golang.org/x/crypto/scrypt"

	"github.com/xcosh-chain/xcosh/types"
)

// Consensus-critical Scrypt parameters. Changing any of these is a hard fork.
// These match Litecoin's Scrypt parameters (N=1024, r=1, p=1, keyLen=32).
const (
	N      = 1024 // CPU/memory cost parameter
	R      = 1    // block size parameter
	P      = 1    // parallelization parameter
	KeyLen = 32   // 256-bit output to match types.Hash
)

// Hasher implements Scrypt proof-of-work hashing.
// Scrypt is memory-hard, making it more ASIC-resistant than SHA256d while
// remaining faster to validate than Argon2id. Used by Litecoin and many altcoins.
type Hasher struct{}

func New() *Hasher { return &Hasher{} }

func (h *Hasher) PoWHash(data []byte) types.Hash {
	out, err := scrypt.Key(data, data, N, R, P, KeyLen)
	if err != nil {
		panic("scrypt.Key failed with consensus parameters: " + err.Error())
	}
	return types.HashFromBytes(out).Reversed()
}

func (h *Hasher) Name() string { return "scrypt" }
