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

package params

import (
	"github.com/xcosh-chain/xcosh/types"
)

// GenesisConfig holds the inputs needed to construct a genesis block.
// This is separate from ChainParams so the genesis mining tool can
// operate on a config before the full params are finalized.
type GenesisConfig struct {
	NetworkName     string
	CoinbaseMessage []byte   // Arbitrary data embedded in the coinbase (e.g., headline).
	Timestamp       uint32   // Unix timestamp for the genesis block.
	Bits            uint32   // Initial difficulty target in compact form.
	Version         uint32   // Block version.
	Reward          uint64   // Coinbase reward value.
	RewardScript    []byte   // PkScript for the coinbase output (recipient placeholder).
	ExtraOutputs    []types.TxOutput // Additional coinbase outputs (e.g., premine burn).
}

// BuildGenesisBlock constructs a genesis block from config.
// The nonce is set to 0; the caller must mine it to find a valid nonce.
func BuildGenesisBlock(cfg GenesisConfig) types.Block {
	outputs := []types.TxOutput{
		{
			Value:    cfg.Reward,
			PkScript: cfg.RewardScript,
		},
	}
	outputs = append(outputs, cfg.ExtraOutputs...)

	coinbaseTx := types.Transaction{
		Version: 1,
		Inputs: []types.TxInput{
			{
				PreviousOutPoint: types.CoinbaseOutPoint,
				SignatureScript:  cfg.CoinbaseMessage,
				Sequence:         0xFFFFFFFF,
			},
		},
		Outputs:  outputs,
		LockTime: 0,
	}

	return types.Block{
		Header: types.BlockHeader{
			Version:    cfg.Version,
			PrevBlock:  types.ZeroHash,
			MerkleRoot: types.ZeroHash, // Must be computed after tx hashing.
			Timestamp:  cfg.Timestamp,
			Bits:       cfg.Bits,
			Nonce:      0,
		},
		Transactions: []types.Transaction{coinbaseTx},
	}
}
