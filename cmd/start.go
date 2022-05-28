/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/aoyouer/go-ngx-reqlimiter/internal"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var (
	rate  float64
	burst int
	ip    string
	port  string
)
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Nginx request rate limiter",
	Long:  `An nginx request rate limiter depends on ip.`,
	Run: func(cmd *cobra.Command, args []string) {
		limiter := internal.NewReqLimiter(ip+":"+port, rate, burst)
		limiter.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rate = *startCmd.Flags().Float64P("rate", "r", 50, "rate limit")
	burst = *startCmd.Flags().IntP("burst", "b", 100, "rate burst")
	ip = *startCmd.Flags().StringP("ip", "i", "127.0.0.1", "bind ip")
	port = *startCmd.Flags().StringP("port", "p", "514", "bind port")
}
