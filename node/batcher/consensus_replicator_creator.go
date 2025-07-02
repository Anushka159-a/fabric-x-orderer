/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package batcher

import (
	"github.com/hyperledger/fabric-x-orderer/common/types"
	node_config "github.com/hyperledger/fabric-x-orderer/node/config"
	"github.com/hyperledger/fabric-x-orderer/node/delivery"
)

//go:generate counterfeiter -o mocks/consensus_state_replicator_creator.go . ConsensusStateReplicatorCreator
type ConsensusStateReplicatorCreator interface {
	CreateStateConsensusReplicator(conf *node_config.BatcherNodeConfig, logger types.Logger) StateReplicator
}

type ConsensusStateReplicatorFactory struct{}

func (c *ConsensusStateReplicatorFactory) CreateStateConsensusReplicator(config *node_config.BatcherNodeConfig, logger types.Logger) StateReplicator {
	var endpoint string
	var tlsCAs []node_config.RawBytes
	for i := 0; i < len(config.Consenters); i++ {
		consenter := config.Consenters[i]
		if consenter.PartyID == config.PartyId {
			endpoint = consenter.Endpoint
			tlsCAs = consenter.TLSCACerts
		}
	}

	if endpoint == "" || len(tlsCAs) == 0 {
		logger.Panicf("Failed finding endpoint and TLS CAs for party %d", config.PartyId)
	}
	return delivery.NewConsensusStateReplicator(tlsCAs, config.TLSPrivateKeyFile, config.TLSCertificateFile, endpoint, nil, logger) // TODO use a real ledger instead of nil
}
