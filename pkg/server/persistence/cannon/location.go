package cannon

import (
	"fmt"
	"time"

	"github.com/ethpandaops/xatu/pkg/proto/xatu"
	"google.golang.org/protobuf/encoding/protojson"
)

type Location struct {
	// LocationID is the location id.
	LocationID interface{} `json:"locationId" db:"location_id"`
	// CreateTime is the timestamp of when the execution record was created.
	CreateTime time.Time `json:"createTime" db:"create_time" fieldopt:"omitempty"`
	// UpdateTime is the timestamp of when the activity record was updated.
	UpdateTime time.Time `json:"updateTime" db:"update_time" fieldopt:"omitempty"`
	// NetworkId is the network id of the location.
	NetworkID string `json:"networkId" db:"network_id"`
	// Type is the type of the location.
	Type string `json:"type" db:"type"`
	// Value is the value of the location.
	Value string `json:"value" db:"value"`
}

// MarshalValueFromProto marshals a proto message into the Value field.
func (l *Location) Marshal(msg *xatu.CannonLocation) error {
	l.NetworkID = msg.NetworkId

	switch msg.Type {
	case xatu.CannonType_BEACON_API_ETH_V2_BEACON_BLOCK_VOLUNTARY_EXIT:
		l.Type = "BEACON_API_ETH_V2_BEACON_BLOCK_VOLUNTARY_EXIT"

		data := msg.GetEthV2BeaconBlockVoluntaryExit()

		b, err := protojson.Marshal(data)
		if err != nil {
			return err
		}

		l.Value = string(b)
	case xatu.CannonType_BEACON_API_ETH_V2_BEACON_BLOCK_PROPOSER_SLASHING:
		l.Type = "BEACON_API_ETH_V2_BEACON_BLOCK_PROPOSER_SLASHING"

		data := msg.GetEthV2BeaconBlockProposerSlashing()

		b, err := protojson.Marshal(data)
		if err != nil {
			return err
		}

		l.Value = string(b)
	case xatu.CannonType_BEACON_API_ETH_V2_BEACON_BLOCK_DEPOSIT:
		l.Type = "BEACON_API_ETH_V2_BEACON_BLOCK_DEPOSIT"

		data := msg.GetEthV2BeaconBlockDeposit()

		b, err := protojson.Marshal(data)
		if err != nil {
			return err
		}

		l.Value = string(b)
	default:
		return fmt.Errorf("unknown type: %s", msg.Type)
	}

	return nil
}

func (l *Location) Unmarshal() (*xatu.CannonLocation, error) {
	msg := &xatu.CannonLocation{
		NetworkId: l.NetworkID,
	}

	switch l.Type {
	case "BEACON_API_ETH_V2_BEACON_BLOCK_VOLUNTARY_EXIT":
		msg.Type = xatu.CannonType_BEACON_API_ETH_V2_BEACON_BLOCK_VOLUNTARY_EXIT

		data := &xatu.CannonLocationEthV2BeaconBlockVoluntaryExit{}

		err := protojson.Unmarshal([]byte(l.Value), data)
		if err != nil {
			return nil, err
		}

		msg.Data = &xatu.CannonLocation_EthV2BeaconBlockVoluntaryExit{
			EthV2BeaconBlockVoluntaryExit: data,
		}
	case "BEACON_API_ETH_V2_BEACON_BLOCK_PROPOSER_SLASHING":
		msg.Type = xatu.CannonType_BEACON_API_ETH_V2_BEACON_BLOCK_PROPOSER_SLASHING

		data := &xatu.CannonLocationEthV2BeaconBlockProposerSlashing{}

		err := protojson.Unmarshal([]byte(l.Value), data)
		if err != nil {
			return nil, err
		}

		msg.Data = &xatu.CannonLocation_EthV2BeaconBlockProposerSlashing{
			EthV2BeaconBlockProposerSlashing: data,
		}
	case "BEACON_API_ETH_V2_BEACON_BLOCK_DEPOSIT":
		msg.Type = xatu.CannonType_BEACON_API_ETH_V2_BEACON_BLOCK_DEPOSIT

		data := &xatu.CannonLocationEthV2BeaconBlockDeposit{}

		err := protojson.Unmarshal([]byte(l.Value), data)
		if err != nil {
			return nil, err
		}

		msg.Data = &xatu.CannonLocation_EthV2BeaconBlockDeposit{
			EthV2BeaconBlockDeposit: data,
		}
	default:
		return nil, fmt.Errorf("unknown type: %s", l.Type)
	}

	return msg, nil
}
