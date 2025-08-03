# Who-Dat MCP Server

This branch contains an MCP server implementation of the Who-Dat WHOIS lookup service.

## Building the Server

To build the MCP server, run the following command:

```bash
go build -o build/mcp-server mcp/main.go
```

## Running the Server

The server can be run directly from the command line:

```bash
./build/mcp-server
```

Alternatively, you can configure it to run as an MCP server in your IDE. Add the following to your MCP settings file:

```json
{
  "mcpServers": {
    "who-dat-mcp": {
      "type": "stdio",
      "command": "/path/to/who-dat-mcp/build/mcp-server",
      "args": [],
      "disabled": false,
      "autoApprove": []
    }
  }
}
