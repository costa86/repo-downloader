# Repo Cloner
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

Downloads remote repositories (GitHub, GitLab, BitBucket) based on a JSON file. 

- [Providing the repositories](#providing-the-repositories)
- [Installation](#installation)
- [Usage](#usage)
- [Concurrent cloning](#concurrent-cloning)
- [Links](#links)


# Providing the repositories
The program takes a JSON file that defines the repositories to be cloned. Use [repos.json](repos.json) as a template.

⚠️ Valid repositories must start with "http", "https", "git" or "ssh", and end with ".git".

Here's a template example:

```json
[
    {
        "folder": "autogit",
        "url": "https://github.com/costa86/autogit.git"
    },
    {
        "folder": "coppy",
        "url": "https://github.com/costa86/coppy.git"
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

Run it with -help

    ./repo-download-linux -help

Output

    Usage of ./repo-download-linux:
    -file string
            JSON file with repositories (default "repos.json")
    -password string
            password (case authentication is required)
    -username string
            username (case authentication is required)


**Option 1: file**

This is an option when no credentials are required

    ./repo-download-linux -file <json_file>

**Option 2: file + username**

In this case, the user will be prompted for a password (***)

    ./repo-download-linux -file <json_file> -username <username>

**Option 3: file + username + password**

This option is more suitable for CI/CD pipelines, where the password can be passed via some mechanism that manages sensitive variables, such as Jenkins.

    ./repo-download-linux -file <json_file> -usename <username> -password <password>


It clones each repository into the specified local folder.

# Concurrent cloning

The program was built in Go, which takes advantage of the Goroutines to clone repositories concurrently. This maximizes throughput by cloning them simultaneously.

# Build from source
Will generate the executables for Windows, Linux and MacOS

Requirements:

* Go (check version in [go.mod](go.mod))
* Make

Command:

    make build

# Links
[![portfolio](https://img.shields.io/badge/my_portfolio-030?style=for-the-badge&logo=ko-fi&logoColor=yellow)](https://costa86.tech/)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/costa86/)

