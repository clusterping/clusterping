/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of the clusterping-cli",
	Long: `Version of the clusterping-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		goVersion := runtime.Version()
		version, err :=getVersion()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("ClusterPing CLI version: %s\n Go version: %s", version, goVersion)
	},
}


func getVersion() (string, error) {
	version, err := os.ReadFile("VERSION")
	if err != nil {
		return "", fmt.Errorf("failed to read version: %w", err)
	}

	return string(version), nil
}



func init() {
	rootCmd.AddCommand(versionCmd)
}
