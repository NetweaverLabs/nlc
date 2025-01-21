/*
Copyright © 2025 toastsandwich
*/
package cmd

import (
	"fmt"

	"github.com/NetweaverLab/nlc/client"
	"github.com/NetweaverLab/nlc/request"
	"github.com/NetweaverLab/nlc/response"
	"github.com/spf13/cobra"
)

func Echo(payload ...string) error {
	req := &request.Request{
		Cmd:  "echo",
		Args: payload,
	}
	resp := &response.Response{}
	dc, err := client.NewDaemonClient()
	if err != nil {
		return err
	}
	if err := dc.Send(req); err != nil {
		return err
	}
	if err := dc.Recieve(resp); err != nil {
		return err
	}
	if resp.Status != "OK" {
		return fmt.Errorf("error by daemon for echo, check /tmp/nld.log")
	}
	str := ""
	for _, p := range resp.Payload {
		str += p + " "
	}
	fmt.Println(str)
	return nil
}

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "basic echo command to test if daemon is responding",
	Long: `with echo you will send a payload to echo and 
it will respond back with the same payload. For example:
nlc echo <payload>
daemon will respond as <payload>
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Echo(args...)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
