package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kucoin-community-chain/kcc-toolkit/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "kcc-toolkit",
		Short: "Toolkit for KuCoin Community Chain",
	}

	client *ethclient.Client
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .kcc-toolkit.yaml)")

	rootCmd.AddCommand(cmdActiveValidators)
	rootCmd.AddCommand(cmdTopValidators)
	rootCmd.AddCommand(cmdStake)
	rootCmd.AddCommand(cmdUnStake)
	rootCmd.AddCommand(cmdWithdrawStaking)
	rootCmd.AddCommand(cmdStakingInfo)
	rootCmd.AddCommand(cmdValidatorInfo)
	rootCmd.AddCommand(cmdValidatorDescription)
	rootCmd.AddCommand(cmdCreateOrEditValidator)
	rootCmd.AddCommand(cmdScanMinedBlocks)
	rootCmd.AddCommand(cmdScanStakers)

	rootCmd.AddCommand(cmdCreateProposal)
	rootCmd.AddCommand(cmdVoteProposal)
	rootCmd.AddCommand(cmdProposals)
	rootCmd.AddCommand(cmdProposalInfo)

	rootCmd.AddCommand(cmdMonitorValidators)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".kcc-toolkit")
	}

	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	var err error
	client, err = ethclient.Dial(viper.GetString("rpc"))
	util.CheckErr(err)
}
