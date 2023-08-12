/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

// diskCheckCmd represents the diskCheck command
var diskCheckCmd = &cobra.Command{
	Use:   "diskCheck",
	Short: "Print memory usage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diskCheck called")
	},
}

func init() {
	InfoCmd.AddCommand(diskCheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskCheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskCheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
