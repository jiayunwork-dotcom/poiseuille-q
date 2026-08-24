Go 圆管层流核算命令行：flow 读 JSON 算 Hagen-Poiseuille 流量、轴心/平均速度、壁剪和雷诺数，Re≥2300 直接报 not_laminar；delta-p 按目标流量反解压差。默认也可 serve 在 :8080 提供同一套 HTTP。

## 构建 / 运行 / 测试

```text
go build ./...
go run . flow --table example/capillary.json
go run . serve --http :8080
go test ./...
```

## 评测镜像

本目录评测专用文件（勿覆盖项目自带 Dockerfile/README）：

- `benzhi.Dockerfile`
- `build_benzhi_docker.sh`
- `BENZHI_README.md`（本文件）

两种架构都要构建并进容器验证：

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh <image-name> linux/arm64
./build_benzhi_docker.sh <image-name> linux/amd64
docker run -it <image-name>:latest
```
