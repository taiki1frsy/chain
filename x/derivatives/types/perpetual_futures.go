package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

var _ PositionInstance = (*PerpetualFuturesPositionInstance)(nil)

func UnpackPerpetualFuturesPositionInstance(positionAny types.Any) PositionInstance {
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualFuturesPositionInstance{}) {
		var position PerpetualFuturesPositionInstance
		err := position.Unmarshal(positionAny.Value)
		if err != nil {
			return nil
		}
		return &position
	}

	return nil
}

// Position Size is considered in denom unit
func NewPerpetualFuturesNetPositionOfMarket(market Market, positionType PositionType, position_size_in_denom_unit sdk.Int) PerpetualFuturesNetPositionOfMarket {
	return PerpetualFuturesNetPositionOfMarket{
		Market:                  market,
		PositionType:            positionType,
		PositionSizeInDenomUnit: position_size_in_denom_unit,
	}
}
