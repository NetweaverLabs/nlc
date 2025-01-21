/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/NetweaverLabs/nlc/client"
	"github.com/NetweaverLabs/nlc/request"
	"github.com/NetweaverLabs/nlc/response"
	"github.com/NetweaverLabs/types"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "used to create account for the forge",
	Long: `The create command will send username and password to daemon which will communicate with our server to create your account. 
For example:

nlc create -u admin -p passwd
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}
		passwd, err := cmd.Flags().GetString("passwd")
		if err != nil {
			return err
		}
		user := &types.User{
			Username: username,
			Password: passwd,
		}
		dc, err := client.NewDaemonClient()
		if err != nil {
			return err
		}
		err = dc.Send(
			&request.Request{
				Cmd:  "create",
				Args: user,
			},
		)
		if err != nil {
			return err
		}
		resp := &response.Response{}
		if err := dc.Recieve(resp); err != nil {
			return err
		}
		if resp.Status != "OK" {
			return errors.New("check daemon logs, something went wrong")
		}
		fmt.Println(resp.Payload)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("username", "u", "", "set username for your account")
	createCmd.Flags().StringP("passwd", "p", "", "set password for your account")
	createCmd.MarkFlagRequired("username")
	createCmd.MarkFlagRequired("passwd")
}
