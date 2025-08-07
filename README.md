## Overview

Reno1 is a hashing-based system configured specifically for devs using templating engines.

Renoi helps with restarting the dev server anytime changes are made to the the 'views' and 'locales' folder and inside sub-apps. 
All other files like .js .ts are all monitored by nodemon so there's no  overhead complexity.

Without having to restart your server ever again after changes to ```.njk ``` or ```.yaml``` files or any templating engine at all
, renoi does that for you like your new autosave 
with no configuration or complexity involved. Your server still works fine, so no worries

## Features:

- Automatically watches for changes in files CASA does not automatically look for. ie nunjucks, locales and other kind of files 
- Restarts the server automatically 


## File Structure

- **`main.go`**: The entry point of the application. It initializes the program and manages the execution flow.
- **`reno.go`**: Contains the core logic and functionality of the Reno application.

## Getting Started

### Prerequisites

Ensure you have the following installed on your system:
- [Go](https://golang.org/dl/) (version 1.18 or later)

### Installation

## Manual Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/persona-mp3/reno.git
    cd reno
    ```


2. After cloning the repository execute the setup script


```bash
# changes file permissions to be executable
chmod u+x setup.sh
./setup.sh
```

## Usage
```bash
renoi # starts watching files
renoi help
```


## Acknowledgments

- Inspired by Git 
- Inspired by Nodemon
