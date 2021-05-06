package cmd

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kucoin-community-chain/kcc-toolkit/contracts/proposal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var (
	cmdCreateProposal = &cobra.Command{
		Use:   "CreateProposal [targetValidator] [details]",
		Short: "CreateProposal for Validator",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			targetValidator := common.HexToAddress(args[0])
			details := args[1]

			proposalContract := common.HexToAddress(viper.GetString("proposalContractAddr"))
			proposalInstance, err := proposal.NewContracts(proposalContract, client)
			if err != nil {
				log.Fatal(err)
			}

			privateKey, err := crypto.HexToECDSA(viper.GetString("proposalPrivateKey"))
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(viper.GetInt64("chainID")))
			if err != nil {
				log.Fatal(err)
			}
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := proposalInstance.CreateProposal(auth, targetValidator, details)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}

	cmdVoteProposal = &cobra.Command{
		Use:   "VoteProposal [proposalID] [confirm]",
		Short: "confirm = true or false",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var proposalID [32]byte
			copy(proposalID[:], common.HexToHash(args[0]).Bytes())
			confirm, err := strconv.ParseBool(args[1])
			if err != nil {
				log.Fatal(err)
			}

			proposalContract := common.HexToAddress(viper.GetString("proposalContractAddr"))
			proposalInstance, err := proposal.NewContracts(proposalContract, client)
			if err != nil {
				log.Fatal(err)
			}

			privateKey, err := crypto.HexToECDSA(viper.GetString("proposalPrivateKey"))
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(viper.GetInt64("chainID")))
			if err != nil {
				log.Fatal(err)
			}
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := proposalInstance.VoteProposal(auth, proposalID, confirm)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}

	cmdGetProposalInfo = &cobra.Command{
		Use:   "GetProposalInfo [fromBlock] [toBlock]",
		Short: "Print ProposalInfo",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fromBlock, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			toBlock, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
			}

			eventSignature := []byte("LogCreateProposal(bytes32,address,address,uint256)")
			hash := crypto.Keccak256Hash(eventSignature)

			proposalContract := common.HexToAddress(viper.GetString("proposalContractAddr"))
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(toBlock)),
				Addresses: []common.Address{proposalContract},
				Topics:    [][]common.Hash{{hash}},
			}

			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				log.Fatal(err)
			}

			proposalAbi, err := abi.JSON(strings.NewReader(proposal.ContractsABI))
			if err != nil {
				log.Fatal(err)
			}

			for _, vLog := range logs {
				LogCreateProposalEvent := struct {
					Id       [32]byte
					Proposer common.Address
					Dst      common.Address
					Time     *big.Int
				}{}

				err := proposalAbi.UnpackIntoInterface(&LogCreateProposalEvent, "LogCreateProposal", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				println("BlockNumber:", vLog.BlockNumber)
				println("TxHash:", vLog.TxHash.Hex())
				println("ProposalID:", vLog.Topics[1].String())
				println("Proposer:", common.HexToAddress(vLog.Topics[2].Hex()).Hex())
				println("targetValidator:", common.HexToAddress(vLog.Topics[3].Hex()).Hex())
				t := time.Unix(LogCreateProposalEvent.Time.Int64(), 0)
				println("Time:", t.Format(time.RFC3339))
				println("----------------------------------")
			}
		},
	}
)
