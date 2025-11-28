package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var taskTitle string
var taskDescription string

var apiURL = "http://localhost:3004"

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task Manager CLI",
	Long:  `Task oluştur, listele, güncelle, sil`,
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "Task listesi",
	Long:  `Tum taskları listele`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tasklar listeleniyor...")
		token := loadToken()
		if token == "" {
			fmt.Println("❌ Önce login olun: taskctl login -e email -p password")
			return
		}

		req, _ := http.NewRequest("GET", apiURL+"/tasks", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err:= http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}
		defer resp.Body.Close()

		body,_ := io.ReadAll(resp.Body)

		fmt.Println(string(body))
		fmt.Println("Status code: ", resp.StatusCode)
	},
}

var taskCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Yeni task oluştur",
	Long:  `Yeni task oluştur`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Yeni task oluşturuluyor: %s\n", taskTitle)
		token := loadToken()
		if token == "" {
			fmt.Println("❌ Önce login olun: taskctl login -e email -p password")
			return
		}

		data := map[string]string{
			"title":        taskTitle,
			"description":  taskDescription,
		}
		jsonData, _ := json.Marshal(data)

		req, _ := http.NewRequest("POST", apiURL+"/tasks", bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")


		resp, err:= http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}
		defer resp.Body.Close()


		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		fmt.Println("Status code: ", resp.StatusCode)
	},
}

func init () {
	rootCmd.AddCommand(taskCmd)
	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskCreateCmd)


	taskCreateCmd.Flags().StringVarP(&taskTitle, "title", "t", "", "Task başlığı (zorunlu)")
	taskCreateCmd.Flags().StringVarP(&taskDescription, "description", "d", "", "Task açıklaması")
	taskCreateCmd.MarkFlagRequired("title")
}