package cmd

import (
	"os"

	"github.com/boydmeyer/goshare/share"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goshare",
	Short: "Share files locally",
	Long: `Goshare helps you share local files from your device over the network. The files could be retrieved by scanning a QR-code or visiting the local URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		directory, _ := cmd.Flags().GetString("directory")
		hideqr, _ := cmd.Flags().GetBool("hide-qr")
		sh, err := share.New(port, directory, hideqr)
		if err != nil {
			panic(err)
		}

		// Start Sharing
		sh.StartServer()
	},
}

func init() {
	rootCmd.Flags().StringP("directory", "d", ".", "Directory to share")
	rootCmd.Flags().StringP("port", "p", "1387", "Port to share files on")
	rootCmd.Flags().Bool("hide-qr", false, "Hide QR Code when sharing")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
