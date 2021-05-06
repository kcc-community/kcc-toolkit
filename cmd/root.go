package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
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
	rootCmd.AddCommand(cmdStakingInfo)
	rootCmd.AddCommand(cmdValidatorInfo)
	rootCmd.AddCommand(cmdValidatorDescription)
	rootCmd.AddCommand(cmdCreateOrEditValidator)

	rootCmd.AddCommand(cmdCreateProposal)
	rootCmd.AddCommand(cmdVoteProposal)
	rootCmd.AddCommand(cmdGetProposalInfo)
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
	if err != nil {
		panic(err)
	}
}
