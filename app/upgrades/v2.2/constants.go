package v2_2

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/UnUniFi/chain/app/upgrades"
)

const UpgradeName string = "v2_2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
