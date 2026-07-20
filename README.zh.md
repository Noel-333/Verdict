# Verdict
**「轻量、易用、便捷的文件完整性验证CLI工具」**

## 简介

  **Verdict是一个基于Golang开发的CLI工具，使用Cobra框架，能够方便的随包部署、随包验证文件哈希。目前支持的算法：**
- *SHA-256*

Verdict不能替代密码学工具，仅作为软件发布、验证文件损坏的便捷辅助工具

## 使用说明
### 基础命令
#### 查看版本信息

```powershell
vec --version
```
返回 `Verdict` 的版本信息。

#### 验证文件完整性
```powershell
vec check
```
自动读取命令行运行目录下的`verdict.json`配置文件，并验证文件完整性。

---

### 配置命令
```powershell
vec init
```
在命令行运行目录处初始化 `Verdict` 。将生成 `verdict.json`。

```
vec snapshot
```
在命令行运行目录处，对根目录文件及子目录文件递归进行哈希，并将`Verdict`版本、哈希算法以及文件哈希储存在 `verdict.json` 中。
使用标志`-a`可指定算法，例如
```
 vec snapshot -a sha256
```
注意，如果不指定算法，默认值为`sha256`。
此外，没有权限访问的文件将被略过。