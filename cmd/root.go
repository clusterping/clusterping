package cmd

import (
	"github.com/spf13/cobra"	
)

var (
	
	rootCmd = &cobra.Command{
		Use:   "clusterping",
		Short: "ClusterPing is a lightweight network monitoring tool for Docker clusters, providing real-time insights with low resource usage.",
		Long:  `The tool enables real-time analysis of **network reachability** between clusters, 
		ensuring that geographically dispersed services can communicate seamlessly. By continuously testing network connections, 
		**ClusterPing** identifies issues such as **latency spikes**, 
		**packet loss**, and **bandwidth bottlenecks**. It provides a set of performance metrics to help administrators ensure that the distributed Docker infrastructure is functioning optimally.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}