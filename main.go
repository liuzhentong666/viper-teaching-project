package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host    string        `mapstructure:"host"`
		Port    int           `mapstructure:"port"`
		Timeout time.Duration `mapstructure:"timeout"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`

	Environment struct {
		Name    string `mapstructure:"name"`
		Verbose bool   `mapstructure:"verbose"`
	} `mapstructure:"environment"`
}

var config Config

func main() {
	fmt.Println("🎓 Viper配置管理教学项目")
	fmt.Println("==============================")

	// 初始化Viper
	viper := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	// 设置默认值
	fmt.Println("\n📝 设置默认配置值...")
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.timeout", "30s")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "guest")
	viper.SetDefault("database.password", "default_password")
	viper.SetDefault("environment.name", "development")
	viper.SetDefault("environment.verbose", false)

	// 集成环境变量
	fmt.Println("\n🌍 集成环境变量...")
	viper.SetEnvPrefix("TEACHING")
	viper.AutomaticEnv()

	// 读取配置文件
	fmt.Println("\n📄 读取配置文件...")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("⚠️ 配置文件未找到或读取失败，使用默认值")
		// 或者直接使用默认值，不区分具体错误类型
	} else {
		fmt.Printf("✅ 配置文件加载成功: %s\n", viper.ConfigFileUsed())
	}

	// 解组配置
	fmt.Println("\n📋 解组配置到结构体...")
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("❌ 配置解组失败: %v", err)
	}

	// 显示配置信息
	fmt.Println("\n📊 当前配置信息:")
	fmt.Printf("  环境: %s\n", config.Environment.Name)
	fmt.Printf("  详细模式: %v\n", config.Environment.Verbose)
	fmt.Printf("  服务器: %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("  超时: %v\n", config.Server.Timeout)
	fmt.Printf("  数据库: %s:%d\n", config.Database.Host, config.Database.Port)
	fmt.Printf("  数据库用户: %s\n", config.Database.User)

	// 配置监控
	fmt.Println("\n🔄 配置监控演示...")
	fmt.Println("   (按Ctrl+C退出，修改config.yaml查看变化)")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("\n🔄 配置已更改，重新加载...")

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("❌ 重新加载配置失败: %v", err)
			return
		}

		if err := viper.Unmarshal(&config); err != nil {
			log.Printf("❌ 重新解组配置失败: %v", err)
			return
		}

		fmt.Println("✅ 配置重新加载成功")
		fmt.Printf("  新环境: %s\n", config.Environment.Name)
		fmt.Printf("  新服务器端口: %d\n", config.Server.Port)
	})

	fmt.Println("\n🎯 程序已进入监控模式，等待配置变化...")
	fmt.Println("   (按Ctrl+C退出，或修改config.yaml测试)")

	// 保持程序运行
	select {}
}
