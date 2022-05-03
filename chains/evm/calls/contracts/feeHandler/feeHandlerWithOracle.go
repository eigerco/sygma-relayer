package feeHandler

import (
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/consts"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/contracts"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"strings"
)

type FeeHandlerWithOracleContract struct {
	contracts.Contract
}

func NewFeeHandlerWithOracleContract(
	client calls.ContractCallerDispatcher,
	feeHandlerWithOracleContractAddress common.Address,
	transactor transactor.Transactor,
) *FeeHandlerWithOracleContract {
	a, _ := abi.JSON(strings.NewReader(consts.FeeHandlerWithOracleABI))
	b := common.FromHex(consts.FeeHandlerWithOracleBin)
	return &FeeHandlerWithOracleContract{contracts.NewContract(feeHandlerWithOracleContractAddress, a, b, client, transactor)}
}

func (f *FeeHandlerWithOracleContract) SetFeeOracle(
	oracleAddr common.Address,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().Msgf("setting fee oracle address %s", oracleAddr.String())
	return f.ExecuteTransaction(
		"setFeeOracle",
		opts,
		oracleAddr,
	)
}

func (f *FeeHandlerWithOracleContract) SetFeeProperties(
	gasUsed uint32,
	feePercent uint16,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().Msgf("setting fee properties: gasUsed: %v, feePercent: %v", gasUsed, feePercent)
	return f.ExecuteTransaction(
		"setFeeProperties",
		opts,
		gasUsed,
		feePercent,
	)
}
