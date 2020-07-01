package system

import (
	eos "github.com/fioprotocol/fio-go/imports/eos-fio"
	"github.com/fioprotocol/fio-go/imports/eos-fio/fecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewSetProds(producers []ProducerKey) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setprods"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `eosio.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    eos.AccountName `json:"producer_name"`
	BlockSigningKey fecc.PublicKey  `json:"block_signing_key"`
}
