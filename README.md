# MyGit — Git-Inspired Version Control System (Go)

MyGit is a simplified version control system written in **Go** that demonstrates the core internal concepts behind Git.
The project was built to understand how version control systems manage files, commits, and history using content-addressable storage.

Instead of using Git itself, this project recreates many of Git’s core ideas from scratch using Go and the local filesystem.

---

# Features

• Initialize a repository
• Content-addressable object storage using SHA-1 hashing
• Staging area (index) for tracking file changes
• Commit snapshots with parent commit references
• Commit history traversal
• Branch creation and switching
• Modular CLI architecture

---

# How It Works

MyGit stores repository metadata inside a hidden directory:

```
.mygit
```

This directory contains all version control data including commits, objects, branches, and repository state.

```
.mygit
├── objects
├── commits
├── branches
├── index
└── HEAD
```

---

# Architecture

The project is structured using modular Go packages.

```
mygit
├── main.go
├── go.mod
├── README.md
├── repository
│   └── repo.go
├── objects
│   └── storage.go
├── commits
│   ├── commit.go
│   └── log.go
├── branch
│   └── branch.go
```

Each module handles a specific responsibility:

| Module     | Responsibility                  |
| ---------- | ------------------------------- |
| repository | repository initialization       |
| objects    | file hashing and object storage |
| commits    | commit creation and history     |
| branch     | branch management               |
| main       | CLI command routing             |

---

# Commands

Initialize repository

```
./mygit init
```

Add file to staging area

```
./mygit add hello.txt
```

Create a commit

```
./mygit commit "first commit"
```

View commit history

```
./mygit log
```

Create a branch

```
./mygit branch feature
```

Switch branch

```
./mygit checkout feature
```

---

# Example Usage

Example workflow:

```
./mygit init

echo "hello world" > hello.txt

./mygit add hello.txt

./mygit commit "first commit"

./mygit log
```

Example output:

```
------
commit: b97d65867cd246c0001125c6ddbaf506bf2744c9
message: sixth commit

------
commit: c99529080f01f78901bb51777bcab1ec3c487d04
message: test commit
```

---

# Key Concepts Implemented

This project demonstrates several important systems programming concepts:

• Content-addressable storage
• SHA-1 hashing for object identification
• Commit graphs using parent references
• Branch pointers and HEAD reference
• Filesystem-based version control storage
• CLI tool architecture

---

# Tech Stack

Language: **Go (Golang)**
Interface: **Command Line Interface (CLI)**
Storage: **Local File System**

---

# Future Improvements

Possible improvements include:

• File diff between commits
• Merge functionality
• Remote repository support
• Better CLI argument parsing
• Performance improvements

---

# Author

Kaavya Indhu
Computer Science Student

Project created to explore the internal design of version control systems like Git.
