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

package logging

import (
	"log/slog"
	"os"
)

var (
	L *slog.Logger

	// DebugMode is set by the -debug CLI flag. When true, subsystems emit
	// hyper-verbose diagnostic output covering block relay, peer topology,
	// sync state, and message flow. This goes beyond slog.LevelDebug by
	// enabling periodic dumps and per-message tracing.
	DebugMode bool

	// StratumDebugMode enables hyper-verbose stratum server diagnostics:
	// every JSON message in/out, share validation byte dumps, header
	// reconstruction details, target comparisons, and job generation.
	// Activated by log-level "stratum" or XCOSH_LOGLEVEL=stratum.
	StratumDebugMode bool
)

func init() {
	L = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// Init replaces the global logger with one configured at the given level and format.
// format may be "text" (default) or "json" for structured JSON output.
// Special level "stratum" enables debug-level logging with StratumDebugMode.
func Init(level string, format ...string) {
	var lvl slog.Level
	switch level {
	case "debug":
		lvl = slog.LevelDebug
	case "stratum":
		lvl = slog.LevelDebug
		StratumDebugMode = true
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: lvl}
	var handler slog.Handler
	if len(format) > 0 && format[0] == "json" {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	} else {
		handler = slog.NewTextHandler(os.Stderr, opts)
	}

	L = slog.New(handler)
	slog.SetDefault(L)
}

// EnableDebug sets DebugMode and forces log level to debug.
func EnableDebug() {
	DebugMode = true
	L = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(L)
}

// P2PSyncDebug emits verbose P2P, header sync, and block-download diagnostics.
// It is a no-op unless DebugMode is true (CLI -debug).
func P2PSyncDebug(msg string, args ...any) {
	if !DebugMode {
		return
	}
	a := append([]any{"component", "p2p_sync"}, args...)
	L.Debug(msg, a...)
}

// ChainSyncDebug emits verbose chain, header-index, reorg, and fork diagnostics.
// It is a no-op unless DebugMode is true (CLI -debug).
func ChainSyncDebug(msg string, args ...any) {
	if !DebugMode {
		return
	}
	a := append([]any{"component", "chain_sync"}, args...)
	L.Debug(msg, a...)
}

// SyncAuditDebug records high-signal sync/IBD decisions (why we are syncing, why
// we dropped a message, header/body phase transitions). Grep for component=sync_audit
// in logs to verify chain convergence behavior on mainnet review.
// No-op unless DebugMode is true (CLI -debug).
func SyncAuditDebug(msg string, args ...any) {
	if !DebugMode {
		return
	}
	a := append([]any{"component", "sync_audit"}, args...)
	L.Debug(msg, a...)
}

// EnableStratumDebug sets StratumDebugMode and forces log level to debug.
func EnableStratumDebug() {
	StratumDebugMode = true
	L = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(L)
}

// StratumDebug emits verbose stratum server diagnostics: message flow,
// share validation, header reconstruction, target comparison.
// No-op unless StratumDebugMode is true.
func StratumDebug(msg string, args ...any) {
	if !StratumDebugMode {
		return
	}
	a := append([]any{"component", "stratum_debug"}, args...)
	L.Debug(msg, a...)
}

// With returns a child logger with additional default attributes.
func With(args ...any) *slog.Logger {
	return L.With(args...)
}
