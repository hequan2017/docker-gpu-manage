# AGENTS.md

## Cursor Cloud specific instructions

### Project Overview

This is the **天启算力管理平台 (Tianqi GPU Computing Platform)** — a GPU computing resource management platform built on gin-vue-admin (Go + Vue 3). It has two main services:

| Service | Path | Port | Command |
|---------|------|------|---------|
| Go Backend | `server/` | 8890 | `cd server && go run main.go` |
| Vue Frontend | `web/` | 8080 | `cd web && npm run dev` |

### Prerequisites

- **Go 1.25+** installed at `/usr/local/go/bin/go` (ensure `PATH` includes `/usr/local/go/bin`)
- **Node.js 20+** (available via nvm)
- **MariaDB/MySQL** running locally on port 3306

### Starting Services

1. **Start MariaDB**: `sudo service mariadb start`
2. **Start Backend**: `cd server && go run main.go` (runs on port 8890)
3. **Start Frontend**: `cd web && npm run dev` (runs on port 8080, proxies `/api` to backend)

### Database Initialization

- The system uses a web-based DB initialization flow. If `server/config.yaml` has empty MySQL fields (as in `config.yaml.bak`), the backend starts in "init mode" and the frontend shows a DB setup wizard at `http://localhost:8080`.
- To initialize: call `POST /init/initdb` with `dbType: "mysql"`, `host: "127.0.0.1"`, `port: "3306"`, `userName: "gva"`, `password: "123456"`, `dbName: "docker-gpu"`, `adminPassword: "123456"`.
- Default admin credentials after init: `admin` / `123456`.
- **Important**: If `config.yaml` already has valid DB settings, the server auto-migrates tables but does NOT seed data (admin user, menus, roles). You must start with the clean `config.yaml.bak` template for a full init.

### Lint & Test Commands

- **Frontend lint**: `cd web && npx eslint .` (pre-existing lint errors exist in the codebase)
- **Backend build check**: `cd server && go build ./...`
- **Backend vet**: `cd server && go vet ./...` (some pre-existing vet warnings)
- **Backend tests**: `cd server && go test ./...` (some pre-existing test failures due to missing fixture files)

### Non-obvious Gotchas

- The frontend Vite proxy strips the `/api` prefix when forwarding to the backend. The backend routes do NOT have `/api` prefix (e.g., `/base/captcha`, not `/api/base/captcha`).
- `server/config.yaml` is rewritten by the init process. Do not commit it with environment-specific values.
- MariaDB is used in place of MySQL in the Cloud Agent environment; they are wire-compatible. The Ubuntu 24.04 MySQL package has installation issues in Docker-in-Docker environments.
- The SSH jumpbox service (port 2026) starts automatically if `jumpbox.enabled: true` in config.yaml. It requires actual Docker containers to be useful.
- Redis is optional and disabled by default (`use-redis: false`).
- No lockfile exists for the frontend; `npm install` is used.
