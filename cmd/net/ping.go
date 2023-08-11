/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	urlPath string
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "These pings network and returns a reponse",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	
		var client = &http.Client{
			Transport: &http.Transport {
				Dial: net.Dialer{timeout: 2 * time.Second}.Dial,
				
			}
		}


	},
}

func init() {

	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "URL to ping")
	pingCmd.MarkFlagRequired("url")

	NetCmd.AddCommand(pingCmd) //we adding the command here because we using the net package

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
