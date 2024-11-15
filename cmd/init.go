/*
 Copyright (c) 2023 <@mdxabu/clusterping@proton.me>

This file is part of the clusterping-cli project.
This source code is licensed under the GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
found in the LICENSE file in the root directory of this source tree.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file for ClusterPing",
	Long: `The init command creates a configuration file for ClusterPing,`,
	Run: createConfigFile,
}


func createConfigFile(cmd *cobra.Command, args []string) {
	fmt.Println("Creating clusterpingconfig.yml file...")
	if checkfileExists() {
		fmt.Println("Config file already exists") 
	} else {
		file, err := os.Create("clusterpingConfig.yml") // create the file
		if err != nil {
			fmt.Println("Error creating config file")
			return
		}
		defer file.Close()
	}
	fmt.Println("initialized the clusterping-cli configuration files")
	
	
}

func checkfileExists() bool {
	if _, err := os.Stat("clusterpingConfig.yml"); os.IsNotExist(err) {
		return false // file does not exist
	}
	return true // file exists
}

func init() {
	rootCmd.AddCommand(initCmd)
}
