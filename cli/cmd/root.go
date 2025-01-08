package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seo-analyzer",
	Short: "SEO Analyzer CLI for keyword analysis",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
