/*
Copyright © 2023 Caner Emeç caner.emec@gmail.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package metadata

import (
	"fmt"

	"github.com/spf13/cobra"
)

// MetadataCmd represents the metadata command
var MetadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Get chaincode metadata info for selected chaincode.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("metadata called")
	},
}

func init() {
	MetadataCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	MetadataCmd.PersistentFlags().String("channel", "mychannel", "Hyperledger fabric channel.")

	MetadataCmd.PersistentFlags().String("chaincode", "basic", "Chaincode name.")
}
