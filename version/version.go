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
package version

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xcosh-chain/xcosh/internal/coinparams"
)

const (
	// Major is the major version component (breaking protocol changes).
	Major = 0

	// Minor is the minor version component (new features, backward-compatible).
	Minor = 15

	// Patch is the patch version component (bug fixes).
	Patch = 0

	// ProtocolVersion is the peer-to-peer wire protocol version.
	// Increment when the wire format changes in a backward-incompatible way.
	ProtocolVersion uint32 = 12

	// MinProtocolVersion is the minimum wire protocol version accepted from peers.
	// Keep in lockstep with ProtocolVersion for releases that break the old wire.
	MinProtocolVersion uint32 = 12

	// ClientName identifies this implementation.
	ClientName = coinparams.NameLower

	// ReleasesURL is the URL where users can download the latest release.
	//
	// FORKERS: Update this URL to point to YOUR project's releases page so
	// that out-of-date notifications direct users to the correct download.
	ReleasesURL = "https://github.com/bams-repo/go-chain/releases"
)

// String returns the semantic version string (e.g. "0.1.0").
func String() string {
	return fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
}

// UserAgent returns the BIP-style user agent (e.g. "/xcosh:0.1.0/").
func UserAgent() string {
	return fmt.Sprintf("%s%s/", coinparams.UserAgentPrefix, String())
}

// SemVer holds a parsed major.minor.patch triple.
type SemVer struct {
	Major, Minor, Patch int
}

// ParseSemVer parses a "major.minor.patch" string. Returns ok=false on failure.
func ParseSemVer(s string) (SemVer, bool) {
	parts := strings.SplitN(s, ".", 3)
	if len(parts) != 3 {
		return SemVer{}, false
	}
	maj, err1 := strconv.Atoi(parts[0])
	min, err2 := strconv.Atoi(parts[1])
	pat, err3 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return SemVer{}, false
	}
	return SemVer{Major: maj, Minor: min, Patch: pat}, true
}

// IsNewerThan returns true if v is strictly newer than other.
func (v SemVer) IsNewerThan(other SemVer) bool {
	if v.Major != other.Major {
		return v.Major > other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor > other.Minor
	}
	return v.Patch > other.Patch
}

func (v SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// ExtractVersionFromUserAgent extracts the semver from a BIP-style user agent
// string like "/xcosh:0.8.1/". Returns the version string and ok=true on
// success.
func ExtractVersionFromUserAgent(ua string) (string, bool) {
	ua = strings.TrimPrefix(ua, "/")
	ua = strings.TrimSuffix(ua, "/")
	idx := strings.LastIndex(ua, ":")
	if idx < 0 || idx >= len(ua)-1 {
		return "", false
	}
	return ua[idx+1:], true
}

// Current returns the running node's version as a SemVer.
func Current() SemVer {
	return SemVer{Major: Major, Minor: Minor, Patch: Patch}
}
