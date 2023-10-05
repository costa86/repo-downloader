# Repo Cloner
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

This program clones multiple GitHub repositories. 

- [Providing the repositories](#providing-the-repositories)
- [Installation](#installation)
- [Usage](#usage)
- [Concurrent cloning](#concurrent-cloning)
- [Links](#links)


# Providing the repositories
The program takes a JSON file that defines the repositories to be cloned. Use [repos.json](repos.json) as a template.

```json
[
    {
        "folder": "autogit",
        "url": "https://github.com/costa86/autogit.git"
    },
    {
        "folder": "coppy",
        "url": "https://github.com/costa86/coppy"
    }
]
```

# Installation
No installation required. Just run the executable providing the required arguments.

|OS|Executable|
|--|--|
|Windows|[repo-download-windows.exe](repo-download-windows.exe)|
|Linux|[repo-download-linux](repo-download-linux)|
|MacOS|[repo-download-darwin](repo-download-darwin)|

# Usage

**Option 1: file**

This is an option when no credentials are required

    ./repo-download-linux <json_file>

**Option 2: file + username**

In this case, the user will be prompted for a password (***)

    ./repo-download-linux <json_file> <username>

**Option 3: file + username + password**

This option is more suitable for CI/CD pipelines, where the password can be passed via some mechanism that manages sensitive variables, such as Jenkins.

    ./repo-download-linux <json_file> <username> <password>


It clones each repository into the specified local folder.

# Concurrent cloning

The program was built in Go, which takes advantage of the Goroutines to clone repositories concurrently. This maximizes throughput by cloning them simultaneously.


# Links
[![portfolio](https://img.shields.io/badge/my_portfolio-030?style=for-the-badge&logo=ko-fi&logoColor=yellow)](https://costa86.tech/)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/costa86/)

