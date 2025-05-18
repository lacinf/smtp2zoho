# Installation Guide

This document explains how to install and run `smtp2zoho` locally or using Docker.

---

## 🔧 Prerequisites

Before running this project, make sure you have the following tools installed:

- [Go (≥ 1.21)](https://golang.org/dl/) — for local builds
- [Docker](https://www.docker.com/) — for containerized execution
- `make` — optional, if you use Makefile automation

> 🔑 See [zoho-auth.md](zoho-auth.md) for instructions to obtain your refresh token from Zoho.

---

## 🖥️ Option 1: Local Build and Run (via Makefile)

This is the recommended approach for development or debugging locally.


```bash
git clone https://github.com/lacinf/smtp2zoho.git
cd smtp2zoho

make build
./smtp2zoho --version
```

> 🧱 The build system uses the current Git tag or commit hash and embeds it in the binary.
> Make sure your environment variables are properly configured (e.g. via `.env`).

---

## 🐳 Option 2: Run via Docker (using Makefile)

The recommended way to run in production is via Docker. A `release` command is available to simplify this:

```bash
git clone https://github.com/lacinf/smtp2zoho.git
cd smtp2zoho

make release
docker run --rm --env-file .env smtp2zoho:"$(git describe --tags --always)"
```

> 🐋 The `release` target builds a versioned Docker image using the current Git tag or commit.

---

## 📥 Option 3: Run using prebuilt Docker image

If you prefer not to build the image yourself, you can use the official prebuilt image:

```bash
docker pull lacinf/smtp2zoho:latest
docker run --rm --env-file .env lacinf/smtp2zoho:latest
```

> 🐋 This is the fastest way to get started. Make sure to create a `.env` file with all required variables.

---

## 🧪 Testing the binary

To verify the binary was built correctly:

```bash
./smtp2zoho --version
```

> ⚙️ This command should print the current version (tag or commit hash).

---

## 📦 Building for production

To create a minimal binary optimized for containers:


```bash
go build -ldflags="-s -w" -o smtp2zoho .
```

> ⚙️ This reduces binary size by stripping debug symbols.


