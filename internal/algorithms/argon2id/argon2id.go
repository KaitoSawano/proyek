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

package argon2id

import (
	"golang.org/x/crypto/argon2"

	"github.com/xcosh-chain/xcosh/types"
)

// Consensus-critical Argon2id parameters. Changing any of these is a hard fork.
const (
	TimeCost    = 1   // single pass — fast enough for block validation
	MemoryCost  = 256 // 256 KiB — light for validation, heavy enough for ASIC resistance
	Parallelism = 1   // single-threaded for determinism
	KeyLen      = 32  // 256-bit output to match types.Hash
)

// Hasher implements Argon2id proof-of-work hashing.
// Uses Argon2id (RFC 9106 recommended hybrid) which combines data-independent
// memory access (ASIC resistance) with data-dependent access (GPU resistance).
type Hasher struct{}

func New() *Hasher { return &Hasher{} }

func (h *Hasher) PoWHash(data []byte) types.Hash {
	out := argon2.IDKey(data, data, TimeCost, MemoryCost, Parallelism, KeyLen)
	return types.HashFromBytes(out).Reversed()
}

func (h *Hasher) Name() string { return "argon2id" }
