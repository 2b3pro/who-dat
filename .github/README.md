<p align="center">
<img src="https://i.ibb.co/J5r1zCP/who-dat-square.png" width="128" /><br />
<b>who-dat-mcp</b><br />
<i>Free & Open Source WHOIS Lookup Service — MCP Server</i>
<br />
<i>An MCP-enabled version of the no-CORS, no-auth WHOIS API that's publicly available or easily self-hostable</i>
<br />
<b>🌐 <a href="https://who-dat.as93.net/">who-dat.as93.net</a></b><br />
</p>

---

<details>
  <summary>Table of Contents</summary>

- [About](#about)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
  - [1. Clone and Build](#1-clone-and-build)
  - [2. Configure Your MCP Client](#2-configure-your-mcp-client)
    - [Claude Code](#claude-code)
    - [Claude Desktop](#claude-desktop)
    - [Cline (VS Code)](#cline-vs-code)
    - [Cursor](#cursor)
- [Available Tools](#available-tools)
- [Running as an HTTP Server](#running-as-an-http-server)
- [Docker](#docker)
- [Environment Variables](#environment-variables)
- [Original Project](#original-project)
- [Contributing](#contributing)
- [License](#license)

</details>

## About

This is a fork of [Who-Dat](https://github.com/Lissy93/who-dat) that adds an **MCP (Model Context Protocol) server** interface. It lets AI assistants — Claude Code, Claude Desktop, Cline, Cursor, and others — perform WHOIS domain lookups as a native tool call.

The server communicates over **stdio** (standard input/output), following the MCP specification.

## Prerequisites

- **Go 1.23+** — [Install Go](https://go.dev/doc/install)
- **Git**

## Setup

### 1. Clone and Build

```bash
git clone https://github.com/2b3pro/who-dat-mcp.git
cd who-dat-mcp
go build -o build/who-dat-mcp mcp/main.go
```

This produces an executable at `build/who-dat-mcp`. Verify it built correctly:

```bash
ls -la build/who-dat-mcp
```

### 2. Configure Your MCP Client

Add the server to your client's MCP configuration. Replace `/absolute/path/to/who-dat-mcp` with the actual path on your machine.

#### Claude Code

Add to `~/.claude/settings.json` (global) or `.claude/settings.json` (project-level):

```json
{
  "mcpServers": {
    "who-dat": {
      "type": "stdio",
      "command": "/absolute/path/to/who-dat-mcp/build/who-dat-mcp"
    }
  }
}
```

#### Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or `%APPDATA%\Claude\claude_desktop_config.json` (Windows):

```json
{
  "mcpServers": {
    "who-dat": {
      "command": "/absolute/path/to/who-dat-mcp/build/who-dat-mcp"
    }
  }
}
```

#### Cline (VS Code)

Add to your `cline_mcp_settings.json`:

```json
{
  "mcpServers": {
    "who-dat": {
      "type": "stdio",
      "command": "/absolute/path/to/who-dat-mcp/build/who-dat-mcp",
      "args": [],
      "disabled": false,
      "autoApprove": []
    }
  }
}
```

#### Cursor

Add to your Cursor MCP settings (Settings > MCP Servers):

```json
{
  "mcpServers": {
    "who-dat": {
      "type": "stdio",
      "command": "/absolute/path/to/who-dat-mcp/build/who-dat-mcp"
    }
  }
}
```

After adding the configuration, restart your client. The `get_whois_multi` tool should appear in the available tools list.

## Available Tools

### `get_whois_multi`

Look up WHOIS information for one or more domains.

**Input:**

```json
{
  "domains": ["example.com", "github.com"]
}
```

**Output:** A JSON object mapping each domain to its parsed WHOIS record (registrar, expiration date, name servers, etc.) or an error message if the lookup failed.

**Example output (abbreviated):**

```json
{
  "example.com": {
    "domain": {
      "domain": "example.com",
      "name_servers": ["a.iana-servers.net", "b.iana-servers.net"],
      "status": ["ACTIVE"],
      "expiration_date": "2025-08-13"
    },
    "registrar": {
      "name": "RESERVED-Internet Assigned Numbers Authority"
    }
  },
  "invalid.tld": {
    "error": "whois: query failed"
  }
}
```

## Running as an HTTP Server

The project also includes a standalone HTTP/REST server with an embedded web UI:

```bash
# Build the frontend first (requires Node.js)
npm install && npm run build

# Build and run the HTTP server
go build -o who-dat .
PORT=8080 ./who-dat
```

**REST API endpoints:**
- `GET /{domain}` — WHOIS lookup for a single domain
- `GET /multi?domains=example.com,github.com` — WHOIS lookup for multiple domains

## Docker

The HTTP server (not the MCP server) can be run via Docker:

```bash
docker build -t who-dat .
docker run -p 8080:8080 who-dat
```

Or pull the pre-built image:

```bash
docker run -p 8080:8080 lissy93/who-dat
```

## Environment Variables

| Variable   | Default | Description                                                                 |
|------------|---------|-----------------------------------------------------------------------------|
| `PORT`     | `8080`  | Port for the HTTP server (not applicable to MCP stdio mode)                 |
| `AUTH_KEY` | —       | Optional API key for the HTTP server. When set, requests must include a `Authorization: Bearer <key>` header. |

## Original Project

Based on [Who-Dat](https://github.com/Lissy93/who-dat) by Alicia Sykes. See the [original README](https://github.com/Lissy93/who-dat/blob/main/README.md) for more about the upstream API and deployment options.

---

## Contributing

Contributions welcome! Please follow the [Code of Conduct](https://github.com/Lissy93/who-dat/blob/main/.github/CODE_OF_CONDUCT.md).

## License

MIT License. See the `LICENSE` file for details.

> _**[2b3pro/who-dat-mcp](https://github.com/2b3pro/who-dat-mcp)** is a fork of **[Lissy93/Who-Dat](https://github.com/Lissy93/who-dat)**, licensed under [MIT](https://github.com/Lissy93/who-dat/blob/HEAD/LICENSE) © [Alicia Sykes](https://aliciasykes.com) 2024._
