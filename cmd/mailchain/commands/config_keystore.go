// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"github.com/mailchain/mailchain"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/config"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/config/defaults"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/setup"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func cfgKeystore(viper *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "setup keystore",
		// Long:  ``,
		PreRunE:  prerunInitConfig(viper),
		PostRunE: config.WriteConfig,
		RunE: func(cmd *cobra.Command, args []string) error {
			keystoreType, err := setup.DefaultKeystore().Select(cmd, mailchain.RequiresValue)
			if err != nil {
				return err
			}
			cmd.Printf("Key store %q configured\n", keystoreType)
			return cmd.Usage()
		},
	}
	cmd.AddCommand(cfgKeystoreNaclFilestore(viper))
	return cmd
}

func cfgKeystoreNaclFilestore(viper *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:      "nacl-filestore",
		Short:    "setup nacl filestore",
		PreRunE:  prerunInitConfig(viper),
		PostRunE: config.WriteConfig,
		RunE: func(cmd *cobra.Command, args []string) error {
			keystoreType, err := setup.DefaultKeystore().Select(cmd, mailchain.StoreNACLFilestore)
			if err != nil {
				return err
			}
			cmd.Printf("Key store %q configured\n", keystoreType)
			return nil
		},
	}
	cmd.Flags().String("keystore-path", defaults.KeystorePath, "Path where keys will be stored.")
	return cmd
}
