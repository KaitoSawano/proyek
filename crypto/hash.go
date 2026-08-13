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

package crypto

import (
	"crypto/sha256"

	"github.com/xcosh-chain/xcosh/types"
)

// DoubleSHA256 computes SHA256(SHA256(data)), the standard consensus hash
// used for block identity, transaction hashing, and merkle roots.
// This is NOT the PoW hash — the PoW hash is provided by the consensus
// engine's Hasher and may use a different algorithm (e.g., Argon2id).
func DoubleSHA256(data []byte) types.Hash {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return types.HashFromBytes(second[:]).Reversed()
}

// HashBlockHeader computes the double-SHA256 of the canonical 80-byte header.
// This is the block's identity hash used for chain indexing, prevblock references,
// and RPC responses. It is always DoubleSHA256 regardless of the PoW algorithm.
func HashBlockHeader(h *types.BlockHeader) types.Hash {
	return DoubleSHA256(h.SerializeToBytes())
}

// HashTransaction computes the double-SHA256 of the canonical transaction bytes.
func HashTransaction(tx *types.Transaction) (types.Hash, error) {
	data, err := tx.SerializeToBytes()
	if err != nil {
		return types.ZeroHash, err
	}
	return DoubleSHA256(data), nil
}
