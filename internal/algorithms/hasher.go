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
package algorithms

import (
	"fmt"

	"github.com/xcosh-chain/xcosh/internal/algorithms/argon2id"
	"github.com/xcosh-chain/xcosh/internal/algorithms/scrypt"
	"github.com/xcosh-chain/xcosh/internal/algorithms/sha256d"
	"github.com/xcosh-chain/xcosh/internal/algorithms/sha256mem"
	"github.com/xcosh-chain/xcosh/params"
	"github.com/xcosh-chain/xcosh/types"
)

// Hasher computes the proof-of-work hash for a serialized block header.
// Implementations must be deterministic: same input always produces same output.
// Implementations must be safe for concurrent use.
type Hasher interface {
	// PoWHash computes the proof-of-work hash of the given data.
	// For header validation, data is the canonical 80-byte serialized header.
	PoWHash(data []byte) types.Hash

	// Name returns the algorithm identifier (e.g., "sha256d", "argon2id").
	Name() string
}

// GetHasher returns a Hasher for the named algorithm.
// For sha256mem, pass chain params via GetHasherForChain so height-gated variants work.
func GetHasher(name string) (Hasher, error) {
	return GetHasherForChain(name, nil)
}

// GetHasherForChain returns a Hasher, binding chain params when the algorithm supports forks.
func GetHasherForChain(name string, p *params.ChainParams) (Hasher, error) {
	switch name {
	case "sha256d":
		return sha256d.New(), nil
	case "argon2id":
		return argon2id.New(), nil
	case "scrypt":
		return scrypt.New(), nil
	case "sha256mem":
		var act uint32
		if p != nil {
			act = p.ActivationHeights[sha256mem.ActivationKey]
		}
		return sha256mem.NewChainHasher(act), nil
	default:
		return nil, fmt.Errorf("unknown PoW algorithm: %q", name)
	}
}
