# Viper配置管理教学项目

## 🎯 项目概述

这是一个专为教学设计的Viper配置管理项目，旨在帮助学习者理解Viper的核心概念和实际应用。项目简单易懂，专注于教学功能，不涉及复杂技术。

## 📚 教学目标

通过这个项目，学习者将掌握：
1. Viper的基本配置管理
2. 多格式配置文件支持（YAML、JSON、TOML等）
3. 环境变量集成
4. 配置分层和默认值
5. 类型安全的配置结构
6. 配置监控和热更新

## 🏗️ 项目结构

```
viper-teaching-project/
├── main_fixed.go        # 主程序（修复版）
├── config/             # 配置文件目录
│   ├── config.yaml     # 基础配置
│   ├── config.dev.yaml # 开发环境配置
│   └── config.prod.yaml# 生产环境配置
└── README.md          # 项目说明
```

## 🚀 快速开始

### 1. 克隆项目
```bash
git clone https://github.com/yourusername/viper-teaching-project.git
cd viper-teaching-project
```

### 2. 下载依赖
```bash
go get -u github.com/spf13/viper
go get -u github.com/fsnotify/fsnotify
```

### 3. 运行项目
```bash
go run main_fixed.go
```

### 4. 修改配置文件
编辑 `config/config.yaml`，修改端口为8081，观察程序输出变化。

## 🎓 教学流程

### 步骤1: 初始化Viper
```go
viper := viper.New()
viper.SetConfigName("config")
viper.AddConfigPath(".")
viper.AddConfigPath("./config")
viper.SetConfigType("yaml")
```

### 步骤2: 设置默认值
```go
viper.SetDefault("server.port", 8080)
viper.SetDefault("database.user", "guest")
```

### 步骤3: 环境变量集成
```go
viper.SetEnvPrefix("TEACHING")
viper.AutomaticEnv()
```

### 步骤4: 读取配置文件
```go
err := viper.ReadInConfig()
if err != nil {
    if viper.IsConfigFileNotFoundError(err) {
        fmt.Println("⚠️ 配置文件未找到，使用默认值")
    } else {
        log.Fatalf("❌ 读取配置文件失败: %v", err)
    }
} else {
    fmt.Printf("✅ 配置文件加载成功: %s\n", viper.ConfigFileUsed())
}
```

### 步骤5: 解组配置到结构体
```go
type Config struct {
    Server struct {
        Host string `mapstructure:"host"`
        Port int    `mapstructure:"port"`
    } `mapstructure:"server"`
}

viper.Unmarshal(&config)
```

### 步骤6: 配置监控
```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("\n🔄 配置已更改，重新加载...")
    viper.ReadInConfig()
    viper.Unmarshal(&config)
})
```

## 📖 配置文件示例

### 基础配置 (config/config.yaml)
```yaml
server:
  host: "localhost"
  port: 8080
  timeout: 30s
  
database:
  host: "localhost"
  port: 5432
  user: "admin"
  password: "secret"
  
environment:
  name: "development"
  verbose: true
```

### 开发环境配置 (config/config.dev.yaml)
```yaml
environment:
  name: "development"
  verbose: true
  
server:
  port: 3000  # 开发环境使用不同端口
```

### 生产环境配置 (config/config.prod.yaml)
```yaml
environment:
  name: "production"
  verbose: false
  
server:
  port: 80    # 生产环境使用80端口
  
database:
  user: "prod_user"
  password: "prod_secret"
```

## 🎯 教学要点

### 1. 配置分层策略
- **基础配置** → **环境配置** → **默认值**
- 环境特定配置会覆盖基础配置
- 默认值在配置文件不存在时使用

### 2. 多配置源优先级
1. 命令行参数（可选）
2. 环境变量
3. 配置文件
4. 默认值

### 3. 类型安全
- 使用结构体映射配置
- mapstructure标签确保正确映射
- 自动类型转换

### 4. 配置监控
- 实时检测配置文件变化
- 无需重启应用更新配置
- 热更新机制

## 📝 扩展学习

### 1. 添加更多配置格式
```go
// 支持JSON格式
viper.SetConfigType("json")
```

### 2. 增强配置验证
```go
func validateConfig(config *Config) error {
    if config.Server.Port <= 0 || config.Server.Port > 65535 {
        return errors.New("无效的服务器端口")
    }
    return nil
}
```

### 3. 添加命令行参数
```go
flag.Int("port", 8080, "服务器端口")
flag.Parse()
viper.BindFlagValues()
```

### 4. 实现配置分层
```go
// 基础配置
viper.ReadInConfig()

// 环境配置
env := os.Getenv("ENV")
if env != "" {
    viper.SetConfigName(fmt.Sprintf("config.%s", env))
    viper.MergeInConfig()
}
```

## 🛡️ 错误处理

### 常见错误及解决方案
1. **配置文件不存在**
   ```go
   if viper.IsConfigFileNotFoundError(err) {
       fmt.Println("使用默认值")
   }
   ```

2. **配置解析错误**
   ```go
   if err := viper.Unmarshal(&config); err != nil {
       log.Fatalf("配置解析失败: %v", err)
   }
   ```

3. **环境变量映射**
   ```go
   viper.SetEnvPrefix("TEACHING")
   viper.AutomaticEnv()
   ```

## 📊 预期输出

运行程序后，您将看到：
```
🎓 Viper配置管理教学项目
==============================

📝 设置默认配置值...
🌍 集成环境变量...
📄 读取配置文件...
✅ 配置文件加载成功: config.yaml
📋 解组配置到结构体...
📊 当前配置信息:
  环境: development
  详细模式: false
  服务器: localhost:8080
  超时: 30s
  数据库: localhost:5432
  数据库用户: guest
  
🔄 配置监控演示...
   (按Ctrl+C退出，修改config.yaml查看变化)

🎯 程序已进入监控模式，等待配置变化...
   (按Ctrl+C退出，或修改config.yaml测试)
```

## 🎓 教学优势

1. **简单易懂**：代码结构清晰，注释详细
2. **循序渐进**：从基础到高级，逐步教学
3. **实时反馈**：配置变化实时显示
4. **多种配置源**：文件、环境变量、默认值
5. **类型安全**：使用结构体确保类型正确
6. **可扩展性**：容易添加新功能进行扩展学习

## 📚 相关资源

- [Viper官方文档](https://github.com/spf13/viper)
- [Go配置管理最佳实践](https://blog.golang.org/viper)
- [配置文件格式对比](https://yaml.org/spec/1.2/spec.html)

## 📝 贡献

欢迎提交Issue和Pull Request来改进这个教学项目！

## 📜 许可证

本项目采用MIT许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。