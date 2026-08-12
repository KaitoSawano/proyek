// Copyright (c) 2024-2026 The Xcosh Contributors
// Distributed under the MIT software license.

package algorithms

import "github.com/xcosh-chain/xcosh/types"

// HeightHasher extends Hasher with height-gated PoW (consensus forks).
// Activation rules are configured when the hasher is constructed.
type HeightHasher interface {
	Hasher
	PoWHashAtHeight(data []byte, height uint32) types.Hash
}
