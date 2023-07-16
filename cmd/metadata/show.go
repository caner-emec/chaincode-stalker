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

	conf "github.com/caner-emec/chaincode-stalker/configs"
	conn "github.com/caner-emec/chaincode-stalker/internal/fabric/connection"
	"github.com/spf13/cobra"
)

var (
	chaincode = "basic"
	channel   = "mychannel"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")

		channel, err := cmd.Flags().GetString("channel")
		if err != nil {
			panic(err)
		}

		chaincode, err := cmd.Flags().GetString("chaincode")
		if err != nil {
			panic(err)
		}

		gw, cl, err := conn.NewConnection(&conf.Conf)
		if err != nil {
			panic(err)
		}
		defer cl.Close()
		defer gw.Close()

		network := gw.GetNetwork(channel)
		contract := network.GetContract(chaincode)

		result, err := contract.EvaluateTransaction("org.hyperledger.fabric:GetMetadata")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(result))

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
