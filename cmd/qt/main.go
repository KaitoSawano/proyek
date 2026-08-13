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
package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/xcosh-chain/xcosh/internal/coinparams"
	"github.com/xcosh-chain/xcosh/params"
	"github.com/xcosh-chain/xcosh/version"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// defaultNetwork is unused but kept for backward-compatible ldflags builds.
// Network selection is fully runtime — see networkForBuild().
var defaultNetwork string

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var appIconPNG []byte

//go:embed assets/trayicon.png
var trayIconPNG []byte

func buildAppMenu(app *App) *menu.Menu {
	appMenu := menu.NewMenu()

	fileMenu := appMenu.AddSubmenu("File")
	fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		wailsRuntime.Quit(app.ctx)
	})

	walletMenu := appMenu.AddSubmenu("Wallet")
	walletMenu.AddText("Wallet security...", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "menu:wallet-security")
	})
	walletMenu.AddText("Encrypt Wallet...", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "menu:encrypt-wallet")
	})
	walletMenu.AddText("Change Passphrase...", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "menu:change-passphrase")
	})
	walletMenu.AddSeparator()
	walletMenu.AddText("Sign Message...", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "menu:sign-message")
	})
	walletMenu.AddText("Verify Message...", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "menu:verify-message")
	})

	viewMenu := appMenu.AddSubmenu("View")
	viewMenu.AddText("Block Explorer", nil, func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "menu:block-explorer")
	})

	helpMenu := appMenu.AddSubmenu("Help")
	helpMenu.AddText("About "+coinparams.Name+" Wallet", nil, func(_ *menu.CallbackData) {
		_, _ = wailsRuntime.MessageDialog(app.ctx, wailsRuntime.MessageDialogOptions{
			Type:    wailsRuntime.InfoDialog,
			Title:   "About " + coinparams.Name + " Wallet",
			Message: coinparams.Name + " Wallet v" + version.String() + "\n\n" + coinparams.CopyrightHolder + "\nDistributed under the MIT software license.",
		})
	})
	helpMenu.AddText("Debug Window", keys.Key("f12"), func(_ *menu.CallbackData) {
		wailsRuntime.EventsEmit(app.ctx, "menu:debug-window")
	})

	return appMenu
}

func networkForBuild() string {
	// Explicit env var always wins — never ignore a testnet/regtest override.
	if env := strings.TrimSpace(os.Getenv("XCOSH_NETWORK")); env != "" {
		return strings.ToLower(env)
	}
	if cliNetwork := networkFromArgs(os.Args[1:]); cliNetwork != "" {
		return cliNetwork
	}

	// Runtime auto-detect: mainnet activates once MiningStartTime has passed.
	if params.Mainnet.MiningStartTime > 0 && time.Now().Unix() >= params.Mainnet.MiningStartTime {
		return "mainnet"
	}
	return "testnet"
}

func networkFromArgs(args []string) string {
	for i := 0; i < len(args); i++ {
		arg := strings.TrimSpace(args[i])
		switch {
		case arg == "-testnet" || arg == "--testnet":
			return "testnet"
		case arg == "-regtest" || arg == "--regtest":
			return "regtest"
		case arg == "-mainnet" || arg == "--mainnet":
			return "mainnet"
		case arg == "-network" || arg == "--network":
			if i+1 < len(args) {
				return strings.ToLower(strings.TrimSpace(args[i+1]))
			}
		case strings.HasPrefix(arg, "-network="):
			return strings.ToLower(strings.TrimSpace(strings.TrimPrefix(arg, "-network=")))
		case strings.HasPrefix(arg, "--network="):
			return strings.ToLower(strings.TrimSpace(strings.TrimPrefix(arg, "--network=")))
		}
	}
	return ""
}

func windowTitle() string {
	net := networkForBuild()
	if net == "mainnet" {
		return coinparams.Name + " Wallet"
	}
	return coinparams.Name + " Wallet [" + net + "]"
}

func main() {
	app := NewApp()

	if err := wails.Run(&options.App{
		Title:             windowTitle(),
		Width:             1200,
		Height:            800,
		MinWidth:          900,
		MinHeight:         600,
		HideWindowOnClose: true,
		Menu:              buildAppMenu(app),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon: appIconPNG,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   windowTitle(),
				Message: "Version " + version.String(),
				Icon:    appIconPNG,
			},
		},
	}); err != nil {
		fmt.Fprintf(os.Stderr, "%s Wallet v%s: %v\n", coinparams.Name, version.String(), err)
		os.Exit(1)
	}
}
