# gitlab-tools

一个用于高效管理 GitLab 项目和用户的命令行工具。

## 功能特性

- 项目管理
  - 创建新项目，支持自定义设置
  - 灵活搜索和列出项目
  - 查看项目详情，包括 SSH URL 和描述
  - 列出项目用户及其权限

- 用户管理
  - 搜索并列出用户详细信息
  - 查看用户最后登录时间和账户状态
  - 列出用户可访问的项目
  - 添加用户到项目并设置特定权限级别

- 命名空间操作
  - 列出和搜索命名空间
  - 查看命名空间详情和统计信息

## 配置

在环境变量中设置您的 GitLab 令牌和 API URL：

```bash
export GITLAB_TOKEN="your-gitlab-token"
export GITLAB_API="your-gitlab-api-url"
```

## 使用方法

### 基本命令

```bash
# 获取项目信息
app get project [项目名称] --namespace [命名空间]

# 创建新项目
app create project [项目名称] --namespace [命名空间] --desc [描述]

# 列出用户
app get users [用户名]

# 添加用户到项目
app create invite [项目名称] --users [用户名1,用户名2] --access [rep|dev|main|owner]
```

### 权限级别

- `rep`: 报告者权限
- `dev`: 开发者权限
- `main`: 维护者权限
- `owner`: 所有者权限

## 贡献

欢迎提交贡献！请随时提交 Pull Request。

## 许可证

本项目采用 MIT 许可证 - 详见 LICENSE 文件。

## 作者

Abel
