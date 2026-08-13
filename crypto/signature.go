// Copyright (c) 2026 AldianOkto. All rights reserved.
// Copyright (c) 2026 Eterbit Core.
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

package crypto

import (
	"fmt"

	"github.com/xcosh-chain/xcosh/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

// SignTransaction signs a specific input of a transaction using SIGHASH_ALL.
// Returns the DER-encoded signature with the sighash type byte appended.
func SignTransaction(tx *types.Transaction, inputIdx int, subscript []byte, privKey *secp256k1.PrivateKey) ([]byte, error) {
	sigHash, err := ComputeSigHash(tx, inputIdx, subscript)
	if err != nil {
		return nil, fmt.Errorf("compute sighash: %w", err)
	}

	sig := ecdsa.Sign(privKey, sigHash[:])
	derSig := sig.Serialize()

	// Append SIGHASH_ALL type byte.
	return append(derSig, SigHashAll), nil
}

// VerifySignature verifies a DER-encoded signature (with sighash type byte)
// against the computed sighash for the given input.
func VerifySignature(tx *types.Transaction, inputIdx int, subscript []byte, sigWithHashType []byte, pubKey *secp256k1.PublicKey) error {
	if len(sigWithHashType) < 2 {
		return fmt.Errorf("signature too short: %d bytes", len(sigWithHashType))
	}

	hashType := sigWithHashType[len(sigWithHashType)-1]
	if hashType != SigHashAll {
		return fmt.Errorf("unsupported sighash type: 0x%02x (only SIGHASH_ALL=0x01 supported)", hashType)
	}

	derSig := sigWithHashType[:len(sigWithHashType)-1]
	sig, err := ecdsa.ParseDERSignature(derSig)
	if err != nil {
		return fmt.Errorf("parse DER signature: %w", err)
	}

	sigHash, err := ComputeSigHash(tx, inputIdx, subscript)
	if err != nil {
		return fmt.Errorf("compute sighash: %w", err)
	}

	if !sig.Verify(sigHash[:], pubKey) {
		return fmt.Errorf("signature verification failed")
	}

	return nil
}

// SignInput is a convenience that signs a transaction input and builds the
// complete P2PKH signature script (sig + pubkey).
func SignInput(tx *types.Transaction, inputIdx int, prevPkScript []byte, privKey *secp256k1.PrivateKey) ([]byte, error) {
	sig, err := SignTransaction(tx, inputIdx, prevPkScript, privKey)
	if err != nil {
		return nil, err
	}
	pubKey := privKey.PubKey().SerializeCompressed()
	return MakeP2PKHSignatureScript(sig, pubKey), nil
}
