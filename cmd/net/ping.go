/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var ( //our variable from our flags adding, we add it here
	urlPath string

	client = http.Client{ //we are creating our client instead of using http.DefaultClient, which have all the timeout and shiit but we are customising our own here
		Transport: &http.Transport{
			Dial: (&net.Dialer{Timeout: 2 * time.Second}).Dial, //here is all our custom timeout and et all
		},
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "These pings network and returns a reponse",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		//What we want it to do here
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
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
