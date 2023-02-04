package ibc

import (
	"github.com/misko9/go-substrate-rpc-client/v4/client"
	sdktypes "github.com/cosmos/cosmos-sdk/codec/types"
	prototypes "github.com/gogo/protobuf/types"
)

const (
	generateConnectionHandshakeProofMethod = "ibc_generateConnectionHandshakeProof"
	queryAcknowledgementsMethod            = "ibc_queryAcknowledgements"
	queryBalanceWithAddressMethod          = "ibc_queryBalanceWithAddress"
	queryChannelMethod                     = "ibc_queryChannel"
	queryChannelClientMethod               = "ibc_queryChannelClient"
	queryChannelsMethod                    = "ibc_queryChannels"
	queryClientStateMethod                 = "ibc_queryClientState"
	queryClientConsensusStateMethod        = "ibc_queryClientConsensusState"
	queryClientsMethod                     = "ibc_queryClients"
	queryConnectionMethod                  = "ibc_queryConnection"
	queryConnectionsMethod                 = "ibc_queryConnections"
	queryConnectionChannelsMethod          = "ibc_queryConnectionChannels"
	queryConnectionUsingClientMethod       = "ibc_queryConnectionUsingClient"
	queryConsensusStateMethod              = "ibc_queryConsensusState"
	queryDenomTraceMethod                  = "ibc_queryDenomTrace"
	queryDenomTracesMethod                 = "ibc_queryDenomTraces"
	queryEventsMethod                      = "ibc_queryEvents"
	queryLatestHeightMethod                = "ibc_queryLatestHeight"
	queryNextSeqRecvMethod                 = "ibc_queryNextSeqRecv"
	queryNewlyCreatedClientMethod          = "ibc_queryNewlyCreatedClient"
	queryPacketsMethod                     = "ibc_queryPackets"
	queryPacketCommitmentsMethod           = "ibc_queryPacketCommitments"
	queryPacketAcknowledgementsMethod      = "ibc_queryPacketAcknowledgements"
	queryPacketCommitmentMethod            = "ibc_queryPacketCommitment"
	queryPacketAcknowledgementMethod       = "ibc_queryPacketAcknowledgement"
	queryPacketReceiptMethod               = "ibc_queryPacketReceipt"
	queryProofMethod                       = "ibc_queryProof"
	querySendPackets                       = "ibc_querySendPackets"
	queryRecvPackets                       = "ibc_queryRecvPackets"
	queryTimestampMethod                   = "ibc_queryTimestamp"
	queryUnreceivedAcknowledgementMethod   = "ibc_queryUnreceivedAcknowledgement"
	queryUnreceivedPacketsMethod           = "ibc_queryUnreceivedPackets"
	queryUpgradedClientMethod              = "ibc_queryUpgradedClient"
	queryUpgradedConnectionStateMethod     = "ibc_queryUpgradedConnectionState"
)

// IBC exposes methods for retrieval of chain data
type IBC struct {
	client client.Client
}

// NewIBC creates a new IBC struct
func NewIBC(cl client.Client) *IBC {
	return &IBC{cl}
}

func parseAny(any *prototypes.Any) (*sdktypes.Any, error) {
	message, err := prototypes.EmptyAny(any)
	if err != nil {
		return nil, err
	}

	err = prototypes.UnmarshalAny(any, message)
	if err != nil {
		return nil, err
	}

	return sdktypes.NewAnyWithValue(message)
}
