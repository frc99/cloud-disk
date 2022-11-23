# CloudDisk
```api
# 创建api服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```