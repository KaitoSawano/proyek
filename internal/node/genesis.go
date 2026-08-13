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

package node

import (
	"fmt"

	"github.com/xcosh-chain/xcosh/internal/algorithms"
	"github.com/xcosh-chain/xcosh/internal/coinparams"
	"github.com/xcosh-chain/xcosh/consensus/pow"
	"github.com/xcosh-chain/xcosh/crypto"
	"github.com/xcosh-chain/xcosh/internal/difficulty"
	"github.com/xcosh-chain/xcosh/logging"
	"github.com/xcosh-chain/xcosh/params"
)

// initNetworkGenesis verifies or mines the genesis block for the given network.
func initNetworkGenesis(p *params.ChainParams, hasher algorithms.Hasher, retargeter difficulty.Retargeter) error {
	if !p.GenesisHash.IsZero() {
		computed := crypto.HashBlockHeader(&p.GenesisBlock.Header)
		if computed != p.GenesisHash {
			return fmt.Errorf("genesis hash verification failed for %s: expected %s, computed %s",
				p.Name, p.GenesisHash.ReverseString(), computed.ReverseString())
		}
		return nil
	}

	if p.Name == "mainnet" {
		return fmt.Errorf("mainnet requires a hardcoded genesis block in params")
	}

	cfg := params.GenesisConfig{
		NetworkName:     p.Name,
		CoinbaseMessage: []byte(fmt.Sprintf("%s %s genesis", coinparams.NameLower, p.Name)),
		Timestamp:       1773212462,
		Bits:            p.InitialBits,
		Version:         1,
		Reward:          p.InitialSubsidy,
		RewardScript:    []byte{0x00},
	}

	block := params.BuildGenesisBlock(cfg)
	genesisEngine := pow.New(hasher, retargeter)
	if err := genesisEngine.MineGenesis(&block); err != nil {
		return fmt.Errorf("mine genesis: %w", err)
	}

	hash := crypto.HashBlockHeader(&block.Header)
	params.InitGenesis(p, block, hash)
	logging.L.Info("genesis block", "hash", hash.ReverseString(), "nonce", block.Header.Nonce)
	return nil
}
