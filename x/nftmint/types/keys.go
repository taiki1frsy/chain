package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "nftmint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for nftmint
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines thee in-memory store key
	MemStoreKey = "mem_nftmint"
)

var (
	// KeyPrefixClassAttributes defines prefix key for ClassAttributes
	KeyPrefixClassAttributes = []byte{0x01}

	// KeyPrefixNFTAttributes defines prefix key for NFTAttributes
	KeyPrefixNFTAttributes = []byte{0x02}

	// KeyPrefixOwningClassList defines prefix key for OwningClassList
	KeyPrefixOwningClassList = []byte{0x03}
)

func NFTAttributesKey(classID, nftID string) []byte {
	nftIdentifier := classID + nftID
	return []byte(nftIdentifier)
}

func OwningClassListKey(owner string) []byte {
	ownerAddr, _ := sdk.AccAddressFromBech32(owner)
	return ownerAddr.Bytes()
}
