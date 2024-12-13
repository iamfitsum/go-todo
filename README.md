# Go Todo

Welcome to the **Go Todo App**! This is a simple project built in Go (Golang) as part of a learning exercise. The app offers both a **Command-Line Interface (CLI)** version and a lightweight **HTTP server** for managing your tasks.

## 📖 Table of Contents
- [✨ Features](#-features)
- [💻 Technologies Used](#-technologies-used)
- [🚀 Getting Started](#-getting-started)
  - [🔧 Prerequisites](#-prerequisites)
  - [📥 Installation](#-installation)
  - [▶️ Running the App](#-running-the-app)
    - [🔹 Command-Line Version](#-command-line-version)
    - [🔹 HTTP Server Version](#-http-server-version)
- [🔍 Code Overview](#-code-overview)
  - [📝 Example Tasks](#-example-tasks)
- [📂 Directory Structure](#-directory-structure)
- [🛡️ License](#-license)

## ✨ Features
- Display a list of tasks
- Add a new task (in the CLI version)
- Serve tasks via a web interface

## 💻 Technologies Used
- Go (Golang)
- HTTP package for serving web pages

## 🚀 Getting Started
Follow these instructions to set up and run the app locally on your machine.

### 🔧 Prerequisites
Make sure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.19 or higher)

### 📥 Installation
1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/iamfitsum/go-todo.git
   cd go-todo
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

### ▶️ Running the App
The app can be run in two modes:

#### 🔹 Command-Line Version
To use the CLI version, uncomment the relevant sections in the `main.go` file under `// Command Line Version` and comment out the HTTP-related code.

#### 🔹 HTTP Server Version
By default, the HTTP version is enabled. The app will start a server on `localhost:8080` with the following endpoints:

- **`/`**: Welcome message
- **`/show-tasks`**: Display the list of tasks

To test it:
1. Open a browser and go to `http://localhost:8080`.
2. Visit `http://localhost:8080/show-tasks` to view the tasks for today.

## 🔍 Code Overview
The project consists of the following:

- **Main Function**: Initializes tasks and sets up the HTTP server.
- **HTTP Handlers**:
  - `helloUser`: Displays a welcome message.
  - `showTasks`: Displays the list of tasks.
- **Command-Line Functions** (commented out by default):
  - `printTasks`: Prints the task list in the terminal.
  - `addTask`: Adds a new task to the list.

### 📝 Example Tasks
The app comes preloaded with these sample tasks:
- Watch Go Crash Course
- Watch Golang Full Course
- Build a TodoList App
- Reward myself with a cup of tea

## 📂 Directory Structure
```
.
├── main.go    # Main application file
├── go.mod     # Go module file
└── README.md  # Documentation file
```

## 🛡️ License
This project is licensed under the MIT License. See the LICENSE file for details.
