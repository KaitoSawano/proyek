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
package dgw

import (
	"math/big"
	"time"

	"github.com/xcosh-chain/xcosh/crypto"
	"github.com/xcosh-chain/xcosh/logging"
	"github.com/xcosh-chain/xcosh/params"
	"github.com/xcosh-chain/xcosh/types"
)

// Retargeter implements Dark Gravity Wave v3, the per-block difficulty
// adjustment algorithm created by Evan Duffield for Dash. DGW uses an
// exponentially-weighted moving average of the last N block targets and
// solve times to compute the next difficulty, adjusting every block.
//
// DGW v3 is superior to simple epoch-based retargeting for small networks
// because it responds to hash rate changes within a single block interval
// rather than waiting for a full retarget epoch. It is more resistant to
// time-warp attacks than v1/v2 due to its use of actual vs expected
// timespan ratios with clamping.
//
// Reference: Dash Core src/pow.cpp DarkGravityWave()
//
// The window size N is taken from ChainParams.RetargetInterval.
// The target solve time T is taken from ChainParams.TargetBlockSpacing.
type Retargeter struct{}

func New() *Retargeter { return &Retargeter{} }

func (r *Retargeter) Name() string { return "dgw" }

// CalcNextBits computes the compact target for the next block using DGW v3.
//
// Algorithm:
//  1. Collect the last N blocks (window = RetargetInterval).
//  2. Compute the average target across the window using a running sum.
//  3. Compute the actual timespan across the window.
//  4. Clamp the actual timespan to [expectedTimespan/3, expectedTimespan*3]
//     to prevent extreme adjustments from timestamp manipulation.
//  5. Adjust: nextTarget = avgTarget * actualTimespan / expectedTimespan.
//  6. Clamp to MinBits (maximum allowed target / minimum difficulty).
func (r *Retargeter) CalcNextBits(
	tip *types.BlockHeader,
	tipHeight uint32,
	getAncestor func(height uint32) *types.BlockHeader,
	p *params.ChainParams,
) uint32 {
	if p.NoRetarget {
		return p.InitialBits
	}

	N := p.RetargetInterval

	// Not enough history — return initial difficulty.
	if tipHeight < N {
		return p.InitialBits
	}

	T := int64(p.TargetBlockSpacing / time.Second)

	// Walk backwards from the tip, collecting N blocks (indices 0..N-1).
	// We also need the block at (tip - N) for the timespan start.
	//
	// blockN   = tip
	// block0   = tip - (N-1)
	// timePrev = tip - N       (only used for its timestamp)
	sumTarget := new(big.Int)
	current := tip
	currentHeight := tipHeight

	for i := uint32(0); i < N; i++ {
		if current == nil {
			logging.L.Error("nil block in DGW window — possible data corruption",
				"component", "difficulty", "height", currentHeight)
			return tip.Bits
		}
		blockTarget := crypto.CompactToBig(current.Bits)
		sumTarget.Add(sumTarget, blockTarget)

		if i < N-1 {
			currentHeight--
			current = getAncestor(currentHeight)
		}
	}

	// current is now at height (tipHeight - N + 1), which is block0.
	// We need the block before block0 for the timespan start timestamp.
	windowStartHeight := tipHeight - N
	windowStart := getAncestor(windowStartHeight)
	if windowStart == nil {
		logging.L.Error("nil ancestor at DGW window start — possible data corruption",
			"component", "difficulty", "height", windowStartHeight)
		return tip.Bits
	}

	// Average target = sumTarget / N.
	avgTarget := new(big.Int).Div(sumTarget, big.NewInt(int64(N)))

	// Actual timespan across the N-block window.
	actualTimespan := int64(tip.Timestamp) - int64(windowStart.Timestamp)

	// Expected timespan = N * T.
	expectedTimespan := int64(N) * T

	// Clamp actual timespan to [expected/3, expected*3] per DGW v3.
	// This prevents extreme difficulty swings from timestamp manipulation.
	if actualTimespan < expectedTimespan/3 {
		actualTimespan = expectedTimespan / 3
	}
	if actualTimespan > expectedTimespan*3 {
		actualTimespan = expectedTimespan * 3
	}

	// nextTarget = avgTarget * actualTimespan / expectedTimespan
	nextTarget := new(big.Int).Mul(avgTarget, big.NewInt(actualTimespan))
	nextTarget.Div(nextTarget, big.NewInt(expectedTimespan))

	// Floor: target must be at least 1.
	if nextTarget.Sign() <= 0 {
		nextTarget.SetInt64(1)
	}

	// Clamp to minimum difficulty (maximum target).
	maxTarget := crypto.CompactToBig(p.MinBits)
	if nextTarget.Cmp(maxTarget) > 0 {
		nextTarget = maxTarget
	}

	return crypto.BigToCompact(nextTarget)
}
