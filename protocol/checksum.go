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

package protocol

import "crypto/sha256"

// doubleSHA256Checksum computes SHA256(SHA256(data)) and returns the first 4 bytes.
// This is used for wire message checksums only — not consensus hashing.
func doubleSHA256Checksum(data []byte) [4]byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	var checksum [4]byte
	copy(checksum[:], second[:4])
	return checksum
}
