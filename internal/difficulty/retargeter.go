// Copyright (c) 2026 AldianOkto. All rights reserved.
// Copyright (c) 2026 Xcosh Core.
// Use of this source code is governed by the Apache License.
// that can be found in the root directory of this repository.
// Project: Eterbit / Blockchain Core
//
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at. <http://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package difficulty

import (
	"fmt"

	"github.com/xcosh-chain/xcosh/internal/difficulty/xcosh"
	"github.com/xcosh-chain/xcosh/internal/difficulty/dgw"
	"github.com/xcosh-chain/xcosh/internal/difficulty/digishield"
	"github.com/xcosh-chain/xcosh/internal/difficulty/lwma"
	"github.com/xcosh-chain/xcosh/params"
	"github.com/xcosh-chain/xcosh/types"
)

// Retargeter computes the next difficulty target for a blockchain.
// Implementations must be deterministic: same chain state always produces the
// same result. Implementations must be safe for concurrent use.
type Retargeter interface {
	// CalcNextBits computes the compact target (nBits) for the next block
	// given the current tip, its height, an ancestor lookup function, and
	// the chain parameters.
	CalcNextBits(tip *types.BlockHeader, tipHeight uint32, getAncestor func(height uint32) *types.BlockHeader, p *params.ChainParams) uint32

	// Name returns the algorithm identifier (e.g., "xcosh", "lwma", "digishield").
	Name() string
}

// GetRetargeter returns a Retargeter for the named difficulty algorithm.
// Adding a new algorithm requires a new sub-package and a new case here.
func GetRetargeter(name string) (Retargeter, error) {
	switch name {
	case "xcosh":
		return xcosh.New(), nil
	case "lwma":
		return lwma.New(), nil
	case "dgw":
		return dgw.New(), nil
	case "digishield":
		return digishield.New(), nil
	default:
		return nil, fmt.Errorf("unknown difficulty algorithm: %q", name)
	}
}
