/*
 Copyright (c) 2023 <@mdxabu/clusterping@proton.me>

This file is part of the clusterping-cli project.
This source code is licensed under the GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
found in the LICENSE file in the root directory of this source tree.
*/
package cmd

import (
	"fmt"
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
		version :=getVersion()
		
		fmt.Printf("ClusterPing CLI version: %s\n Go version: %s", version, goVersion)
	},
}


func getVersion() string{
	return "v0.0.1"
}



func init() {
	rootCmd.AddCommand(versionCmd)
}
