package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// ============================================
// EGZERSİZ 2: Viper ile Config Yönetimi
// ============================================
//
// VIPER NEDİR?
// ------------
// Config dosyalarını okur (yaml, json, env)
// TypeScript'te dotenv + config paketleri gibi
//
// NEDEN KULLANILIR?
// -----------------
// - Ayarları koddan ayırır (port, db bilgileri vs)
// - Farklı ortamlar için farklı config (dev, prod)
// - Environment variable'ları otomatik okur
//
// TEMEL KULLANIM:
// ---------------
// viper.SetConfigName("config")  → config.yaml dosyasını ara
// viper.SetConfigType("yaml")    → Dosya tipi
// viper.AddConfigPath(".")       → Hangi klasörde ara
// viper.ReadInConfig()           → Dosyayı oku
// viper.GetString("app.name")    → Değer al
// viper.SetDefault("key", val)   → Default değer
//
// ============================================

// GÖREV 1: Viper'ı import et

func main() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("app.name", "my-task-manager")
	viper.SetDefault("server.port", "8080")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config okunamadı: %w", err))
	}

	appname := viper.GetString("app.name")
	port := viper.GetString("server.port")
	debug := viper.GetString("app.debug")
	// GÖREV 2: Config dosyası ayarlarını yap

	// GÖREV 3: Default değerler tanımla

	// GÖREV 4: Config dosyasını oku

	// GÖREV 5: Değerleri oku ve yazdır

	fmt.Println("Config başarıyla okundu!\n", "app.name:", appname, "server.port:", port, "app.debug:", debug)
}

// ============================================
// TEST ETMEK İÇİN:
// ============================================
// 1. Önce config.yaml dosyasını oluştur (aşağıda)
// 2. go run main.go
// ============================================
