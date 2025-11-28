package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var loginEmail string
var loginPassword string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Giriş yap ve token kaydet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Giriş yapılıyor...")
		data := map[string]string{
			"email":    loginEmail,
			"password": loginPassword,
		}
		jsonData, _ := json.Marshal(data)
		res, err := http.Post(apiURL+"/login", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		var result map[string]interface{}
		json.Unmarshal(body, &result)

		if token, ok := result["token"].(string); ok {
			err = saveToken(token)
			if err != nil {
				fmt.Println("Token kaydedilemedi:", err)
				return
			}
			fmt.Println("✅ Giriş başarılı, token kaydedildi")
		} else {
			fmt.Println("❌ Giriş başarısız:", string(body))
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&loginEmail, "email", "e", "", "E-posta adresi (zorunlu)")
	loginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "Parola (zorunlu)")
	loginCmd.MarkFlagRequired("email")
	loginCmd.MarkFlagRequired("password")
}