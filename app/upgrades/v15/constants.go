package v15

import (
	store "cosmossdk.io/store/types"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/packetforward/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v7/types"

	"github.com/osmosis-labs/osmosis/v24/app/upgrades"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v24/x/poolmanager/types"
	protorevtypes "github.com/osmosis-labs/osmosis/v24/x/protorev/types"
	valsetpreftypes "github.com/osmosis-labs/osmosis/v24/x/valset-pref/types"
)

// UpgradeName defines the on-chain upgrade name for the Osmosis v15 upgrade.
const UpgradeName = "v15"

// pool ids to migrate
const (
	stOSMO_OSMOPoolId   = 833
	stJUNO_JUNOPoolId   = 817
	stSTARS_STARSPoolId = 810
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{poolmanagertypes.StoreKey, valsetpreftypes.StoreKey, protorevtypes.StoreKey, icqtypes.StoreKey, packetforwardtypes.StoreKey},
		Deleted: []string{},
	},
}
