package evm

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/events"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var TestTimeout = time.Second * 600

type Client interface {
	LatestBlock() (*big.Int, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
}

func WaitForProposalExecuted(client Client, bridge common.Address) error {
	startBlock, _ := client.LatestBlock()

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		Addresses: []common.Address{bridge},
		Topics: [][]common.Hash{
			{events.ProposalExecutionSig.GetTopic()},
		},
	}
	timeout := time.After(TestTimeout)
	ch := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		return err
	}

	defer sub.Unsubscribe()
	for {
		select {
		case <-ch:
			return nil
		case err := <-sub.Err():
			if err != nil {
				return err
			}
		case <-timeout:
			return errors.New("test timed out waiting for proposal execution event")
		}
	}
}

func WaitUntilBridgeReady(client Client, bridge common.Address) error {
	startBlock, _ := client.LatestBlock()
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		Addresses: []common.Address{bridge},
		Topics: [][]common.Hash{
			{events.FeeHandlerChangedSig.GetTopic()},
		},
	}

	ch := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	for {
		select {
		case <-ch:
			return nil
		case err := <-sub.Err():
			if err != nil {
				return err
			}
		}
	}
}
