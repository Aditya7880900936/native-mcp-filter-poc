# Native Envoy MCP Filter Evaluation PoC

A prototype environment for evaluating native Envoy-based MCP request processing, dynamic tool-aware routing, and authorization as a potential replacement for ext_proc-based MCP parsing architectures.

---

## Overview

This project explores how Envoy-native HTTP filters can be used to process MCP-style JSON-RPC requests **without relying on external processing hops** (`ext_proc`). The prototype demonstrates:

- **MCP-style request parsing**
- **Dynamic tool extraction**
- **Tool-aware backend routing**
- **Metadata/header propagation**
- **Authorization via `ext_authz`**
- **Multi-backend MCP-style architecture**
- **Dockerized local development environment**

---

## Architecture

```text
Client
  |
  v
Envoy Gateway
  |
  |-- Lua Filter
  |     └── Extract MCP tool name
  |     └── Inject routing metadata/header
  |
  |-- ext_authz Filter
  |     └── Validate access policies
  |
  |-- Dynamic Route Matching
  |
  +--> Weather Backend
  |
  +--> Github Backend
```

---

## Features Implemented

### MCP Request Processing

Supports MCP-style JSON-RPC requests such as:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "github.search_repo"
  }
}
```

### Dynamic Tool Extraction

The Envoy Lua filter extracts tool names directly from request bodies.

**Example extracted values:**
- `weather.get`
- `github.search_repo`

### Dynamic Backend Routing

Requests are dynamically routed based on the extracted tool names.

| Tool              | Backend         |
|-------------------|----------------|
| weather.get       | weather_service |
| github.search_repo| github_service  |

### Authorization Layer

Integrated with Envoy's `ext_authz` filter.

**Current demo policy:**
| Backend | Policy  |
|---------|---------|
| weather | allowed |
| github  | denied  |

---

## Tech Stack

- Golang
- Envoy Proxy
- Docker Compose
- Envoy Lua Filters
- Envoy ext_authz

---

## Project Structure

```
native-mcp-filter-poc/
│
├── envoy/
│   ├── envoy.yaml
│   └── Dockerfile
│
├── backend-weather/
│   ├── main.go
│   └── Dockerfile
│
├── backend-github/
│   ├── main.go
│   └── Dockerfile
│
├── auth-service/
│   ├── main.go
│   └── Dockerfile
│
├── test-client/
│   └── request.json
│
├── docker-compose.yml
│
└── README.md
```

---

## Running the Prototype

### Start Services

```sh
docker compose up --build
```

### Testing Requests

**Allowed Request:**

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "weather.get"
  }
}
```

**Blocked Request:**

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "github.search_repo"
  }
}
```

---

## Key Findings

- Envoy-native request inspection can support MCP-style request processing.
- Dynamic routing can be achieved without ext_proc.
- Metadata-driven authorization flows integrate naturally with ext_authz.
- Native MCP filter support exists in current Envoy builds, though configuration/runtime compatibility still requires deeper evaluation.
- Lua filters provide a lightweight intermediate solution for experimentation and migration prototyping.

---

## Future Work

- Native MCP filter configuration evaluation
- Backend aggregation/federation
- SSE / streamable HTTP support
- Session management
- Benchmarking against ext_proc architecture
- Integration with Authorino / Kuadrant policy flows

---

## Goal

This prototype is part of an exploration into reducing latency, simplifying MCP parsing architecture, and evaluating the feasibility of migrating toward Envoy-native MCP request handling.
