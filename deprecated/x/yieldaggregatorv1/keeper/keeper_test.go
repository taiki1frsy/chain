package keeper_test

// import (
// 	"testing"

// 	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/suite"

// 	simapp "github.com/UnUniFi/chain/app"
// )

// type KeeperTestSuite struct {
// 	suite.Suite

// 	ctx sdk.Context
// 	app *simapp.App
// }

// func (suite *KeeperTestSuite) SetupTest() {
// 	isCheckTx := false
// 	app := simapp.Setup(isCheckTx)

// 	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
// 	suite.app = app
// }

// func TestKeeperSuite(t *testing.T) {
// 	suite.Run(t, new(KeeperTestSuite))
// }
