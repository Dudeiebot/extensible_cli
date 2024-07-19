/*
Copyright © dudeiebot
*/
package info

import (
	"github.com/spf13/cobra"
)

// diskCheckCmd represents the diskCheck command
var diskCheckCmd = &cobra.Command{
	Use:   "diskCheck",
	Short: "Print memory usage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
    cmd.Help()
  },
}

  //testing out some shiit here to see how far
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
