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
package discovery

import (
	"log"

	"github.com/xcosh-chain/xcosh/logging"
	"github.com/xcosh-chain/xcosh/internal/store"
)

// Discovery manages peer address discovery from multiple sources.
type Discovery struct {
	peerStore store.PeerStore
	seeds     []string
}

// New creates a new Discovery instance.
func New(ps store.PeerStore, seeds []string) *Discovery {
	return &Discovery{
		peerStore: ps,
		seeds:     seeds,
	}
}

// Bootstrap returns the initial set of peer addresses to connect to.
// Combines static seeds with persisted peers.
func (d *Discovery) Bootstrap() []string {
	var addrs []string
	addrs = append(addrs, d.seeds...)

	stored, err := d.peerStore.GetPeers()
	if err != nil {
		log.Printf("[discovery] failed to load stored peers: %v", err)
	} else {
		addrs = append(addrs, stored...)
	}

	out := deduplicate(addrs)
	logging.P2PSyncDebug("discovery.Bootstrap",
		"seed_count", len(d.seeds),
		"total_addrs", len(out))
	return out
}

// AddPeer persists a newly discovered peer address.
func (d *Discovery) AddPeer(addr string) {
	logging.P2PSyncDebug("discovery.AddPeer", "addr", addr)
	d.peerStore.PutPeer(addr)
}

// RemovePeer removes a peer address from the persistent store.
func (d *Discovery) RemovePeer(addr string) {
	logging.P2PSyncDebug("discovery.RemovePeer", "addr", addr)
	d.peerStore.RemovePeer(addr)
}

func deduplicate(addrs []string) []string {
	seen := make(map[string]struct{}, len(addrs))
	result := make([]string, 0, len(addrs))
	for _, a := range addrs {
		if _, ok := seen[a]; !ok {
			seen[a] = struct{}{}
			result = append(result, a)
		}
	}
	return result
}
