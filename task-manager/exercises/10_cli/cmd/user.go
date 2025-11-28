package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Kullanıcı Manager CLI",
	Long:  `Kullanıcı oluştur, listele, güncelle, sil`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kullanıcılar listeleniyor...")
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}