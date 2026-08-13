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
package consensus

import (
	"math/big"

	"github.com/xcosh-chain/xcosh/internal/algorithms"
	"github.com/xcosh-chain/xcosh/params"
	"github.com/xcosh-chain/xcosh/types"
)

// Engine defines the pluggable consensus interface.
// The baseline implementation is PoW. Future implementations may include
// identity-bound ticket-based sequential-work consensus.
//
// Design boundary: the Engine is responsible for:
//   - validating that a header satisfies the consensus challenge
//   - validating block-level consensus rules
//   - computing retarget adjustments
//   - preparing block templates (challenge parameters)
//
// The Engine is NOT responsible for:
//   - transaction validation beyond coinbase rules
//   - chain selection (that's the chain manager's job)
//   - networking
type Engine interface {
	// ValidateHeader checks that the header satisfies consensus rules
	// (e.g., PoW target, timestamp bounds, correct difficulty) given the parent header and params.
	// getAncestor returns the header at a given main-chain height (needed for difficulty calculation).
	ValidateHeader(header *types.BlockHeader, parent *types.BlockHeader, height uint32, getAncestor func(uint32) *types.BlockHeader, p *params.ChainParams) error

	// ValidateBlock checks block-level consensus rules: coinbase, merkle root,
	// transaction ordering, size limits, etc.
	ValidateBlock(block *types.Block, height uint32, p *params.ChainParams) error

	// CalcNextBits computes the difficulty bits for the next block given
	// the chain state at the current tip.
	CalcNextBits(tip *types.BlockHeader, tipHeight uint32, getAncestor func(height uint32) *types.BlockHeader, p *params.ChainParams) uint32

	// PrepareHeader fills in consensus-specific fields on a new block header
	// being constructed for mining (e.g., sets bits).
	PrepareHeader(header *types.BlockHeader, parent *types.BlockHeader, parentHeight uint32, getAncestor func(height uint32) *types.BlockHeader, p *params.ChainParams) error

	// SealHeader attempts to find a valid nonce for the header at the given height.
	// Returns true if a valid nonce was found within maxIterations.
	// The header's Nonce field is updated in place.
	SealHeader(header *types.BlockHeader, target types.Hash, height uint32, p *params.ChainParams, maxIterations uint64) (bool, error)

	// CalcBlockWeight returns the consensus weight contributed by a single block.
	// For PoW, this is the work implied by the header's difficulty bits.
	// For other engines, this could be ticket count, VRF score, etc.
	// The chain manager accumulates these to determine the heaviest chain.
	CalcBlockWeight(header *types.BlockHeader) *big.Int

	// Hasher returns the PoW hash algorithm used by this engine.
	// The PoW hash is distinct from the block identity hash (DoubleSHA256).
	Hasher() algorithms.Hasher

	// Name returns the consensus engine name for logging/identification.
	Name() string
}
