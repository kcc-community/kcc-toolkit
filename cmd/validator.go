package cmd

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kucoin-community-chain/kcc-toolkit/contracts/validator"
	"github.com/kucoin-community-chain/kcc-toolkit/types"
	"github.com/kucoin-community-chain/kcc-toolkit/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"strconv"
	"strings"
)

var (
	cmdActiveValidators = &cobra.Command{
		Use:   "ActiveValidators",
		Short: "Print ActiveValidators",
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			validators, err := validatorInstance.GetActiveValidators(nil)
			util.CheckErr(err)

			for i, address := range validators {
				println(i, address.Hex())
			}
		},
	}

	cmdTopValidators = &cobra.Command{
		Use:   "TopValidators",
		Short: "Print TopValidators",
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			validators, err := validatorInstance.GetTopValidators(nil)
			util.CheckErr(err)

			for i, address := range validators {
				println(i, address.Hex())
			}
		},
	}

	cmdStakingInfo = &cobra.Command{
		Use:   "StakingInfo [staker] [targetValidator]",
		Short: "Print StakingInfo",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			staker := common.HexToAddress(args[0])
			targetValidator := common.HexToAddress(args[1])

			coins, unStakeBlock, index, err := validatorInstance.GetStakingInfo(nil, staker, targetValidator)
			util.CheckErr(err)

			println("Stake Amount:", util.ToDecimal(coins, 18).String())
			println("unStake Block:", unStakeBlock.Uint64())
			println("Stake Index:", index.Uint64())
		},
	}

	cmdValidatorInfo = &cobra.Command{
		Use:   "ValidatorInfo [targetValidator]",
		Short: "Print ValidatorInfo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			targetValidator := common.HexToAddress(args[0])

			feeAddr, status, coins, Income, totalJailed, lastWithdrawProfitsBlock, stakers, err := validatorInstance.GetValidatorInfo(nil, targetValidator)
			util.CheckErr(err)

			println("feeAddr:", feeAddr.Hex())
			println("status:", status)
			println("Stake Amount:", util.ToDecimal(coins, 18).String())
			println("Income:", util.ToDecimal(Income, 18).String())
			println("totalJailed:", util.ToDecimal(totalJailed, 18).String())
			println("lastWithdrawProfitsBlock:", lastWithdrawProfitsBlock.Uint64())
			println("stakers:")
			for i, staker := range stakers {
				println("\t", i, staker.Hex())
			}
		},
	}

	cmdValidatorDescription = &cobra.Command{
		Use:   "ValidatorDescription [targetValidator]",
		Short: "Print ValidatorDescription",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			targetValidator := common.HexToAddress(args[0])

			moniker, identity, website, email, details, err := validatorInstance.GetValidatorDescription(nil, targetValidator)
			util.CheckErr(err)

			println("moniker:", moniker)
			println("identity:", identity)
			println("website:", website)
			println("email:", email)
			println("details:", details)
		},
	}

	cmdStake = &cobra.Command{
		Use:   "Stake [targetValidator] [amount]",
		Short: "Stake KCS to a targetValidator",
		Long:  "The staked coins of validator must >= 32",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			targetValidator := common.HexToAddress(args[0])
			amount, err := strconv.ParseFloat(args[1], 64)
			util.CheckErr(err)

			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			privateKey, err := crypto.HexToECDSA(viper.GetString("stakePrivateKey"))
			util.CheckErr(err)

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			util.CheckErr(err)

			gasPrice, err := client.SuggestGasPrice(context.Background())
			util.CheckErr(err)

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(viper.GetInt64("chainID")))
			util.CheckErr(err)

			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = util.ToWei(amount, 18)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := validatorInstance.Stake(auth, targetValidator)
			util.CheckErr(err)

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}

	cmdUnStake = &cobra.Command{
		Use:   "UnStake [targetValidator]",
		Short: "UnStake KCS from a targetValidator",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			targetValidator := common.HexToAddress(args[0])

			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			privateKey, err := crypto.HexToECDSA(viper.GetString("stakePrivateKey"))
			util.CheckErr(err)

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			util.CheckErr(err)

			gasPrice, err := client.SuggestGasPrice(context.Background())
			util.CheckErr(err)

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(viper.GetInt64("chainID")))
			util.CheckErr(err)

			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = util.ToWei(0, 18)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := validatorInstance.UnStake(auth, targetValidator)
			util.CheckErr(err)

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}

	cmdWithdrawStaking = &cobra.Command{
		Use:   "WithdrawStaking [targetValidator]",
		Short: "WithdrawStaking KCS from the Validator Contract",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			targetValidator := common.HexToAddress(args[0])

			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			privateKey, err := crypto.HexToECDSA(viper.GetString("stakePrivateKey"))
			util.CheckErr(err)

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			util.CheckErr(err)

			gasPrice, err := client.SuggestGasPrice(context.Background())
			util.CheckErr(err)

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(viper.GetInt64("chainID")))
			util.CheckErr(err)

			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = util.ToWei(0, 18)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := validatorInstance.WithdrawStaking(auth, targetValidator)
			util.CheckErr(err)

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}

	cmdCreateOrEditValidator = &cobra.Command{
		Use:   "CreateOrEditValidator",
		Short: "CreateOrEditValidator",
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			util.CheckErr(err)

			privateKey, err := crypto.HexToECDSA(viper.GetString("validatorPrivateKey"))
			util.CheckErr(err)

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			util.CheckErr(err)

			gasPrice, err := client.SuggestGasPrice(context.Background())
			util.CheckErr(err)

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(viper.GetInt64("chainID")))
			util.CheckErr(err)

			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := validatorInstance.CreateOrEditValidator(
				auth,
				fromAddress,
				viper.GetString("validatorMoniker"),
				viper.GetString("validatorIdentity"),
				viper.GetString("validatorWebsite"),
				viper.GetString("validatorEmail"),
				viper.GetString("validatorDetails"),
			)
			util.CheckErr(err)

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}

	cmdScanMinedBlocks = &cobra.Command{
		Use:   "ScanMinedBlocks",
		Short: "Scan Mined Blocks of Mine",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey, err := crypto.HexToECDSA(viper.GetString("validatorPrivateKey"))
			util.CheckErr(err)

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			validator := crypto.PubkeyToAddress(*publicKeyECDSA)

			db, err := sql.Open("sqlite3", "./validator.sqlite")

			util.CheckErr(err)
			defer db.Close()

			tables, err := db.Query(`SELECT count(name) as num FROM sqlite_master WHERE type = 'table' AND name = 'blocks';
`)
			util.CheckErr(err)
			defer tables.Close()

			var tableExist int
			for tables.Next() {
				err = tables.Scan(&tableExist)
				util.CheckErr(err)
			}

			if tableExist == 0 {
				log.Println("table [blocks] not exist, will create it")

				_, err = db.Exec(`
				create table if not exists blocks (
					id integer not null primary key autoincrement,
					block_number integer not null,
					block_hash varchar(128) not null,
					block_miner varchar(64) not null,
					block_reward varchar(64) not null
				);
				
				create unique index blocks_block_number_uindex on blocks (block_number);
				create unique index blocks_block_hash_uindex on blocks (block_hash);
				`)

				util.CheckErr(err)
			}

			blocks, err := db.Query(fmt.Sprintf(`select id,block_number from blocks where block_miner = "%s" order by block_number desc limit 1`, strings.ToLower(validator.Hex())))
			util.CheckErr(err)
			defer blocks.Close()

			var lastBlock types.Block
			for blocks.Next() {
				err = blocks.Scan(&lastBlock.Id, &lastBlock.BlockNumber)
				util.CheckErr(err)
			}

			log.Printf("last scaned block number is: %d\n", lastBlock.BlockNumber)

			latestBlockNumber, err := client.BlockNumber(context.Background())
			util.CheckErr(err)
			log.Printf("latest block number is: %d\n", latestBlockNumber)

			stmt, err := db.Prepare(`insert into blocks (block_number, block_hash, block_miner, block_reward) values (?, ?, ?, ?)`)
			util.CheckErr(err)

			i := lastBlock.BlockNumber + 1
			for i < latestBlockNumber {
				log.Printf("scanning block %d\n", i)
				header, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(i)))
				util.CheckErr(err)

				if strings.ToLower(header.Coinbase.Hex()) != strings.ToLower(validator.Hex()) {
					i++
					continue
				}

				block, err := client.BlockByHash(context.Background(), header.Hash())
				util.CheckErr(err)

				var blockReward decimal.Decimal
				if block.Transactions().Len() > 0 {
					for _, transaction := range block.Transactions() {
						receipt, err := client.TransactionReceipt(context.Background(), transaction.Hash())
						util.CheckErr(err)

						blockReward = blockReward.Add(util.ToDecimal(util.CalcGasCost(receipt.GasUsed, transaction.GasPrice()), 18))
					}
				}

				result, err := stmt.Exec(block.NumberU64(), block.Hash().Hex(), strings.ToLower(block.Header().Coinbase.Hex()), blockReward.String())
				util.CheckErr(err)

				id, err := result.LastInsertId()
				util.CheckErr(err)
				if id > 0 {
					i++
				}

				log.Printf("block %d mined with reward: %s\n", block.NumberU64(), blockReward.String())
			}
		},
	}

	cmdScanStakers = &cobra.Command{
		Use:   "ScanStakers",
		Short: "Scan Stakers of Mine",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey, err := crypto.HexToECDSA(viper.GetString("validatorPrivateKey"))
			util.CheckErr(err)

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			validatorAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

			db, err := sql.Open("sqlite3", "./validator.sqlite")

			util.CheckErr(err)
			defer db.Close()

			tables, err := db.Query(`SELECT count(name) as num FROM sqlite_master WHERE type = 'table' AND name = 'stakers';
`)
			util.CheckErr(err)
			defer tables.Close()

			var tableExist int
			for tables.Next() {
				err = tables.Scan(&tableExist)
				util.CheckErr(err)
			}

			if tableExist == 0 {
				log.Println("table [stakers] not exist, will create it")

				_, err = db.Exec(`
				create table if not exists stakers (
					id integer not null primary key autoincrement,
					block_number integer not null,
					transaction_hash varchar(128) not null,
					staker varchar(128) not null,
					validator varchar(64) not null,
				    action_type char(10) not null,
					amount varchar(64) not null
				);
				
				create index stakers_block_number_index on stakers (block_number);
				create unique index stakers_transaction_hash_uindex on stakers (transaction_hash);
				`)

				util.CheckErr(err)
			}

			stakers, err := db.Query(fmt.Sprintf(`select id,block_number from stakers where validator = "%s" order by block_number desc limit 1`, strings.ToLower(validatorAddr.Hex())))
			util.CheckErr(err)
			defer stakers.Close()

			var lastStake types.Block
			for stakers.Next() {
				err = stakers.Scan(&lastStake.Id, &lastStake.BlockNumber)
				util.CheckErr(err)
			}

			log.Printf("last scaned block number is: %d\n", lastStake.BlockNumber)

			latestBlockNumber, err := client.BlockNumber(context.Background())
			util.CheckErr(err)
			log.Printf("latest block number is: %d\n", latestBlockNumber)

			stmt, err := db.Prepare(`insert into stakers (block_number, transaction_hash, staker, validator, action_type, amount) values (?, ?, ?, ?, ?, ?)`)
			util.CheckErr(err)

			logStakeEventSignature := []byte("LogStake(address,address,uint256,uint256)")
			logStakeEventSignatureHash := crypto.Keccak256Hash(logStakeEventSignature)
			logUnStakeEventSignature := []byte("LogUnStake(address,address,uint256,uint256)")
			logUnStakeEventSignatureHash := crypto.Keccak256Hash(logUnStakeEventSignature)

			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(lastStake.BlockNumber + 1)),
				ToBlock:   big.NewInt(int64(latestBlockNumber)),
				Addresses: []common.Address{validatorContract},
				Topics:    [][]common.Hash{{logStakeEventSignatureHash, logUnStakeEventSignatureHash}},
			}

			logs, err := client.FilterLogs(context.Background(), query)
			util.CheckErr(err)

			validatorAbi, err := abi.JSON(strings.NewReader(validator.ContractsABI))
			util.CheckErr(err)

			for _, vLog := range logs {

				if strings.ToLower(validatorAddr.Hex()) != strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex()) {
					continue
				}

				LogStakeEvent := struct {
					Staker  common.Address
					Val     common.Address
					Staking *big.Int
					Time    *big.Int
				}{}

				LogUnStakeEvent := struct {
					Staker common.Address
					Val    common.Address
					Amount *big.Int
					Time   *big.Int
				}{}

				if vLog.Topics[0].Hex() == logStakeEventSignatureHash.Hex() {
					err := validatorAbi.UnpackIntoInterface(&LogStakeEvent, "LogStake", vLog.Data)
					util.CheckErr(err)

					log.Println(
						vLog.BlockNumber,
						vLog.TxHash.Hex(),
						strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex()),
						strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex()),
						"stake",
						util.ToDecimal(LogStakeEvent.Staking, 18).String(),
					)

					result, err := stmt.Exec(
						vLog.BlockNumber,
						vLog.TxHash.Hex(),
						strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex()),
						strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex()),
						"stake",
						util.ToDecimal(LogStakeEvent.Staking, 18).String(),
					)
					util.CheckErr(err)

					_, err = result.LastInsertId()
					util.CheckErr(err)
				}

				if vLog.Topics[0].Hex() == logUnStakeEventSignatureHash.Hex() {
					err = validatorAbi.UnpackIntoInterface(&LogUnStakeEvent, "LogUnStake", vLog.Data)
					util.CheckErr(err)

					log.Println(
						vLog.BlockNumber,
						vLog.TxHash.Hex(),
						strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex()),
						strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex()),
						"unstake",
						util.ToDecimal(LogUnStakeEvent.Amount, 18).String(),
					)

					result, err := stmt.Exec(
						vLog.BlockNumber,
						vLog.TxHash.Hex(),
						strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex()),
						strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex()),
						"unstake",
						util.ToDecimal(LogUnStakeEvent.Amount, 18).String(),
					)
					util.CheckErr(err)

					_, err = result.LastInsertId()
					util.CheckErr(err)
				}
			}
		},
	}
)
