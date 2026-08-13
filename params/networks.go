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
package params

import (
	"time"

	"github.com/xcosh-chain/xcosh/types"
)

const (
	// MinedSupply is the total coins producible through mining alone (785 million coins).
	MinedSupply = 78500000000000000

	// MainnetPremineAmount is 20% of the mined supply, added on top at block 1.
	MainnetPremineAmount = MinedSupply / 5

	// MaxMoneyValue is the absolute maximum base units that can ever exist,
	// including the mainnet premine. Consensus and mempool checks use this.
	MaxMoneyValue = MinedSupply + MainnetPremineAmount

	// MaxTxSize is the maximum serialized size of a single transaction in bytes.
	MaxTxSize = 100000

	// 20% premine on top of mined supply for testnet.
	TestnetPremineAmount = MinedSupply / 5
)

var (
	// Hardcoded burn marker script for trackable burns/premine accounting.
	TestnetBurnScript = []byte("burn:testnet:premine:v1")
)

// Mainnet is the primary network.
var Mainnet = &ChainParams{
	Name:         "mainnet",
	DataDirName:  "",
	NetworkMagic: [4]byte{0xFA, 0x1C, 0xC0, 0x01},
	DefaultPort:  19333,
	AddressPrefix: 0x00,

	GenesisBlock: types.Block{
		Header: types.BlockHeader{
			Version:    1,
			PrevBlock: types.ZeroHash,
			MerkleRoot: types.Hash{
				0x2f, 0xd5, 0x06, 0x2d, 0x25, 0xf4, 0x21, 0x67,
				0x59, 0x23, 0x2c, 0x8c, 0xb8, 0x65, 0x76, 0x59,
				0xbd, 0xf8, 0x54, 0xa0, 0xf0, 0x0d, 0xfa, 0x63,
				0x63, 0x9b, 0xea, 0x6d, 0x0a, 0x1e, 0xea, 0xc6,
			},
			Timestamp: 1774175035,
			Bits:      0x1e346dbd,
			Nonce:     2147488231,
		},
		Transactions: []types.Transaction{{
			Version: 1,
			Inputs: []types.TxInput{{
				PreviousOutPoint: types.CoinbaseOutPoint,
				SignatureScript:  []byte("xcosh genesis"),
				Sequence:         0xFFFFFFFF,
			}},
			Outputs: []types.TxOutput{{
				Value:    50_0000_0000,
				PkScript: []byte{0x00},
			}},
			LockTime: 0,
		}},
	},
	GenesisHash: types.Hash{
		0x76, 0x95, 0x98, 0x49, 0x69, 0x49, 0xe8, 0xc2,
		0x55, 0x32, 0xc7, 0x69, 0xac, 0x16, 0xb8, 0xb7,
		0xb6, 0x74, 0x89, 0x99, 0xdb, 0x45, 0x40, 0x58,
		0x46, 0xc4, 0xed, 0x8a, 0xc0, 0xc6, 0xe5, 0x25,
	},

	TargetBlockSpacing:   35 * time.Second,
	RetargetInterval:     144,
	TargetTimespan:       144 * 35 * time.Second,
	MaxTimeFutureDrift:   2 * time.Hour,
	MinTimestampRule:     "median-11",

	InitialBits:      0x1e346dbd,
	MinBits:          0x207fffff,
	NoRetarget:       false,

	MaxBlockSize:     1000000,
	MaxBlockTxCount:  10000,

	InitialSubsidy:         5000000000,
	SubsidyHalvingInterval: 7850000,

	CoinbaseMaturity: 100,

	MaxReorgDepth: 288,

	TimewarpGracePeriod: 35 * time.Second, // Updated to 1 block spacing
	PeerStoreMaxSize:    4096,

	MaxMempoolSize:    5000,
	MinRelayTxFee:     1000,
	MinRelayTxFeeRate: 1,
	MempoolExpiry:     336 * time.Hour,

	SeedNodes: []string{
		"95.179.203.47:19333",
		"207.246.117.14:19333",
		"[2001:19f0:5400:3322:5400:06ff:fe1d:ce90]:19333",
	},

	MiningStartTime: 1777338000,

	PremineHeight: 1,
	PremineAmount: MainnetPremineAmount,
	PremineScript: []byte{
		0x76, 0xa9, 0x14,
		0xc7, 0xff, 0xfe, 0xed, 0x7b, 0x2b, 0x51, 0x77,
		0x94, 0x93, 0x73, 0x99, 0x84, 0xa4, 0x51, 0xf6,
		0x16, 0xd9, 0x9c, 0x64,
		0x88, 0xac,
	},

	ActivationHeights: map[string]uint32{
		"locktime": 1,
		"timewarp": 1,
	},
}

// Testnet is the public test network.
var Testnet = &ChainParams{
	Name:         "testnet",
	DataDirName:  "testnet12",
	NetworkMagic: [4]byte{0xFA, 0x1C, 0xC0, 0x03},
	DefaultPort:  19334,
	AddressPrefix: 0x6F,

	GenesisBlock: types.Block{
		Header: types.BlockHeader{
			Version:    1,
			PrevBlock: types.ZeroHash,
			MerkleRoot: types.Hash{
				0xb1, 0x88, 0x4d, 0xb0, 0x63, 0x4b, 0xe0, 0x81,
				0x08, 0x02, 0x9f, 0x73, 0xfe, 0x53, 0xdc, 0xb0,
				0x93, 0xeb, 0x40, 0xdd, 0xf7, 0x54, 0x23, 0x6c,
				0x65, 0xbc, 0x4f, 0x2f, 0xc2, 0x1e, 0xe5, 0xe2,
			},
			Timestamp: 1744325400,
			Bits:      0x1f666659,
			Nonce:     1289,
		},
		Transactions: []types.Transaction{{
			Version: 1,
			Inputs: []types.TxInput{{
				PreviousOutPoint: types.CoinbaseOutPoint,
				SignatureScript:  []byte("xcosh testnet1 genesis"),
				Sequence:         0xFFFFFFFF,
			}},
			Outputs: []types.TxOutput{
				{
					Value:    5000000000,
					PkScript: []byte{0x00},
				},
				{
					Value:    TestnetPremineAmount,
					PkScript: TestnetBurnScript,
				},
			},
			LockTime: 0,
		}},
	},
	GenesisHash: types.Hash{
		0xc1, 0xc0, 0x86, 0xe2, 0x5b, 0x2f, 0xc0, 0x9d,
		0xf7, 0x57, 0x2b, 0xd4, 0xed, 0xd0, 0x34, 0x1f,
		0x64, 0xf2, 0x17, 0x1b, 0xcd, 0xe2, 0x6d, 0xed,
		0x71, 0x8d, 0xc0, 0xeb, 0x6c, 0x39, 0x2c, 0x24,
	},

	TargetBlockSpacing:   35 * time.Second,
	RetargetInterval:     60,
	TargetTimespan:       60 * 35 * time.Second,
	MaxTimeFutureDrift:   105 * time.Second, // N*T/20 = 60*35/20
	MinTimestampRule:     "median-11",

	InitialBits:              0x1f666659,
	MinBits:                  0x207fffff,
	NoRetarget:               false,
	AllowMinDifficultyBlocks: true,
	MinDifficultyGap:         5 * time.Minute,

	MaxBlockSize:     2000000,
	MaxBlockTxCount:  10000,

	InitialSubsidy:         5000000000,
	SubsidyHalvingInterval: 7850000,

	CoinbaseMaturity: 100,

	MaxReorgDepth: 1000,

	TimewarpGracePeriod: 35 * time.Second,
	PeerStoreMaxSize:    1024,

	MaxMempoolSize:    5000,
	MinRelayTxFee:     1000,
	MinRelayTxFeeRate: 1,
	MempoolExpiry:     336 * time.Hour,

	SeedNodes: []string{
		"95.179.203.47:19334",
		"207.246.117.14:19334",
	},

	ActivationHeights: map[string]uint32{
		"locktime":      1,
		"mindiffblocks": 1,
		"timewarp":      1,
		"sha256mem":     85000,
	},
}

// Regtest is a local regression-test network.
var Regtest = &ChainParams{
	Name:         "regtest",
	DataDirName:  "regtest",
	NetworkMagic: [4]byte{0xFA, 0x1C, 0xC0, 0xFF},
	DefaultPort:  19444,
	AddressPrefix: 0x6F,

	TargetBlockSpacing:   1 * time.Second,
	RetargetInterval:     1,
	TargetTimespan:       1 * time.Second,
	MaxTimeFutureDrift:   10 * time.Minute,
	MinTimestampRule:     "prev+1",

	InitialBits:      0x207fffff,
	MinBits:          0x207fffff,
	NoRetarget:       true,

	MaxBlockSize:     4000000,
	MaxBlockTxCount:  50000,

	InitialSubsidy:         5000000000,
	SubsidyHalvingInterval: 150,

	CoinbaseMaturity: 1,

	MaxReorgDepth: 0,

	TimewarpGracePeriod: 10 * time.Minute,
	PeerStoreMaxSize:    512,

	MaxMempoolSize:    10000,
	MinRelayTxFee:     0,
	MinRelayTxFeeRate: 0,
	MempoolExpiry:     1 * time.Hour,

	SeedNodes: []string{},

	ActivationHeights: map[string]uint32{
		"locktime": 1,
		"timewarp": 1,
	},
}

// NetworkByName returns chain params by network name.
func NetworkByName(name string) *ChainParams {
	switch name {
	case "mainnet":
		return Mainnet
	case "testnet":
		return Testnet
	case "regtest":
		return Regtest
	default:
		return nil
	}
}

// InitGenesis computes and sets the genesis block and hash for the given params.
func InitGenesis(p *ChainParams, genesisBlock types.Block, genesisHash types.Hash) {
	p.GenesisBlock = genesisBlock
	p.GenesisHash = genesisHash
}
