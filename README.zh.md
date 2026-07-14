# Verdict
**「轻量、易用、便捷的文件完整性验证CLI工具」**

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
在命令行运行目录处，对根目录文件及子目录处的文件进行哈希，并将`Verdict`版本、哈希算法以及文件哈希储存在 `verdict.json` 中。