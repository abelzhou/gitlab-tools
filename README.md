# gitlab-tools

A command-line tool for managing GitLab projects and users efficiently.

## Features

- Project Management
  - Create new projects with customizable settings
  - Search and list projects with flexible filtering
  - View project details including SSH URLs and descriptions
  - List project users and their permissions

- User Management
  - Search and list users with detailed information
  - View user's last sign-in time and account status
  - List projects a user has access to
  - Add users to projects with specific access levels

- Namespace Operations
  - List and search namespaces
  - View namespace details and statistics

## Configuration

Set your GitLab token and API URL in your environment:

```bash
export GITLAB_TOKEN="your-gitlab-token"
export GITLAB_API="your-gitlab-api-url"
```

## Usage

### Basic Commands

```bash
# Get project information
app get project [project-name] --namespace [namespace]

# Create new project
app create project [project-name] --namespace [namespace] --desc [description]

# List users
app get users [username]

# Add users to project
app create invite [project-name] --users [username1,username2] --access [rep|dev|main|owner]
```

### Access Levels

- `rep`: Reporter permissions
- `dev`: Developer permissions
- `main`: Maintainer permissions
- `owner`: Owner permissions

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

Abel
