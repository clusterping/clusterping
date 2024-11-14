/*
 Copyright (c) 2023 <@mdxabu/clusterping@proton.me>
 
This file is part of the clusterping-cli project.
This source code is licensed under the GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007 
found in the LICENSE file in the root directory of this source tree.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file for ClusterPing",
	Long: `The init command creates a configuration file for ClusterPing,`,
	Run: createConfigFile,
}


func createConfigFile(cmd *cobra.Command, args []string) {
	fmt.Println("Creating config file...")

	/*
	Need to execute the file creation logic here
	*/
	
}

func init() {
	rootCmd.AddCommand(initCmd)
}
