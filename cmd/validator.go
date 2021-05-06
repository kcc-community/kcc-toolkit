package cmd

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kucoin-community-chain/kcc-toolkit/contracts/validator"
	"github.com/kucoin-community-chain/kcc-toolkit/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"strconv"
)

var (
	cmdActiveValidators = &cobra.Command{
		Use:   "ActiveValidators",
		Short: "Print ActiveValidators",
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			if err != nil {
				log.Fatal(err)
			}

			validators, err := validatorInstance.GetActiveValidators(nil)
			if err != nil {
				log.Fatal(err)
			}

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
			if err != nil {
				log.Fatal(err)
			}

			validators, err := validatorInstance.GetTopValidators(nil)
			if err != nil {
				log.Fatal(err)
			}

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
			if err != nil {
				log.Fatal(err)
			}

			staker := common.HexToAddress(args[0])
			targetValidator := common.HexToAddress(args[1])

			coins, unStakeBlock, index, err := validatorInstance.GetStakingInfo(nil, staker, targetValidator)
			if err != nil {
				log.Fatal(err)
			}

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
			if err != nil {
				log.Fatal(err)
			}

			targetValidator := common.HexToAddress(args[0])

			feeAddr, status, coins, Income, totalJailed, lastWithdrawProfitsBlock, stakers, err := validatorInstance.GetValidatorInfo(nil, targetValidator)
			if err != nil {
				log.Fatal(err)
			}

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
			if err != nil {
				log.Fatal(err)
			}

			targetValidator := common.HexToAddress(args[0])

			moniker, identity, website, email, details, err := validatorInstance.GetValidatorDescription(nil, targetValidator)
			if err != nil {
				log.Fatal(err)
			}

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
			amount, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
			}

			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			if err != nil {
				log.Fatal(err)
			}

			privateKey, err := crypto.HexToECDSA(viper.GetString("stakePrivateKey"))
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
			auth.Value = util.ToWei(amount, 18)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := validatorInstance.Stake(auth, targetValidator)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}

	cmdCreateOrEditValidator = &cobra.Command{
		Use:   "CreateOrEditValidator",
		Short: "CreateOrEditValidator",
		Run: func(cmd *cobra.Command, args []string) {
			validatorContract := common.HexToAddress(viper.GetString("validatorContractAddr"))
			validatorInstance, err := validator.NewContracts(validatorContract, client)
			if err != nil {
				log.Fatal(err)
			}

			privateKey, err := crypto.HexToECDSA(viper.GetString("validatorPrivateKey"))
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

			tx, err := validatorInstance.CreateOrEditValidator(
				auth,
				fromAddress,
				viper.GetString("validatorMoniker"),
				viper.GetString("validatorIdentity"),
				viper.GetString("validatorWebsite"),
				viper.GetString("validatorEmail"),
				viper.GetString("validatorDetails"),
			)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("tx sent: %s", tx.Hash().Hex())
		},
	}
)
