package models

import "encoding/json"

type LRPStartAuctionState int

const (
	LRPStartAuctionStateInvalid LRPStartAuctionState = iota
	LRPStartAuctionStatePending
	LRPStartAuctionStateClaimed
)

type LRPStartAuction struct {
	DesiredLRP DesiredLRP `json:"desired_lrp"`

	InstanceGuid string `json:"instance_guid"`
	Index        int    `json:"index"`

	State     LRPStartAuctionState `json:"state"`
	UpdatedAt int64                `json:"updated_at"`
}

func NewLRPStartAuctionFromJSON(payload []byte) (LRPStartAuction, error) {
	auction := LRPStartAuction{}

	err := json.Unmarshal(payload, &auction)
	if err != nil {
		return LRPStartAuction{}, err
	}

	return auction, auction.Validate()
}

func (auction LRPStartAuction) Validate() error {
	var validationError ValidationError

	if auction.InstanceGuid == "" {
		validationError = append(validationError, ErrInvalidJSONMessage{"instance_guid"})
	}

	err := auction.DesiredLRP.Validate()
	if err != nil {
		validationError = append(validationError, err)
	}

	if len(validationError) > 0 {
		return validationError
	}

	return nil
}

func (auction LRPStartAuction) ToJSON() []byte {
	bytes, err := json.Marshal(auction)
	if err != nil {
		panic(err)
	}

	return bytes
}
