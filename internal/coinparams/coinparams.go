// Copyright (c) 2024-2026 The Fairchain Contributors
// Fairchain is an experiment in modularity, designed to improve on the work
// of Satoshi Nakamoto and to inspire more creative genius in the space.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

package coinparams

import "math"

var (
	// CoinsPerBaseUnit is 10^Decimals, computed at init.
	CoinsPerBaseUnit float64
)

func init() {
	CoinsPerBaseUnit = math.Pow(10, float64(Decimals))
}

const (
	// Name is the human-readable coin name (e.g., "Fairchain", "Bitcoin", "Litecoin").
	Name = "Xcosh"

	// NameLower is the lowercase form used in paths, config files, and CLI.
	NameLower = "xcosh"

	// Ticker is the exchange ticker symbol (e.g., "FAIR", "BTC", "LTC").
	Ticker = "XCOSH"

	// DaemonName is the binary name for the full node daemon.
	DaemonName = "xcoshd"

	// CLIName is the binary name for the command-line RPC client.
	CLIName = "xcosh-cli"

	// GenesisToolName is the binary name for the genesis mining tool.
	GenesisToolName = "xcosh-genesis"

	// AdversaryToolName is the binary name for the adversary testing tool.
	AdversaryToolName = "xcosh-adversary"

	// GUIName is the binary name for the GUI wallet (e.g., "fairchain-qt").
	GUIName = "xcosh-qt"

	// DefaultDataDirName is the hidden directory name in the user's home (e.g., ".fairchain").
	DefaultDataDirName = ".xcosh"

	// ConfFileName is the INI-style config file name (e.g., "fairchain.conf").
	ConfFileName = "xcosh.conf"

	// CoinbaseTag is the ASCII tag embedded in coinbase transactions.
	CoinbaseTag = "xcosh"

	// RPCRealm is the HTTP Basic Auth realm for the RPC server.
	RPCRealm = "xcosh-rpc"

	// UserAgentPrefix is the BIP-style user agent prefix (e.g., "/fairchain:").
	UserAgentPrefix = "/xcosh:"

	// CopyrightHolder is the name used in LICENSE and legal notices.
	CopyrightHolder = "Xcosh Contributors"

	// BaseUnitName is the name of the smallest indivisible unit (e.g., "satoshi").
	BaseUnitName = "unit"

	// DisplayUnitName is the name of the display unit (e.g., "BTC", "FAIR").
	// Used in RPC responses like "balance_fair" instead of "balance_btc".
	DisplayUnitName = "xcosh"

	// Decimals is the number of decimal places between the display unit and
	// the smallest indivisible base unit. For example, Bitcoin uses 8 (1 BTC
	// = 100,000,000 satoshi). Valid range: 0–18. Changing this is a
	// consensus / hard-fork change.
	Decimals = 8

	// Algorithm is the PoW hash algorithm name. Must match a registered
	// algorithm in internal/algorithms/. Changing this is a hard fork.
	// Options: "sha256d" (Bitcoin-compatible), "argon2id" (CPU-fair, RFC 9106),
	//          "scrypt" (Litecoin-style), "sha256mem" (memory-hard SHA256)
	Algorithm = "sha256mem"

	// DifficultyAlgorithm is the difficulty retargeting algorithm name.
	// Must match a registered algorithm in internal/difficulty/.
	// Changing this is a consensus / hard-fork change.
	// Options: "bitcoin"    (Nakamoto-style epoch retarget with EDA),
	//          "lwma"       (zawy12 LWMA-1, per-block weighted moving average),
	//          "dgw"        (Dark Gravity Wave v3, per-block averaging, Dash-style),
	//          "digishield" (DigiShield v3, per-block asymmetric dampening, Dogecoin/Zcash-style)
	DifficultyAlgorithm = "lwma"
)
