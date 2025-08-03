<p align="center">
<img src="https://i.ibb.co/J5r1zCP/who-dat-square.png" width="128" /><br />
<i>Free & Open Source WHOIS Lookup Service - MCP Version</i>
<br />
<i>An MCP-enabled version of the no-CORS, no-auth API that's publicly available or easily self-hostable</i>
<br />
<b>🌐 <a href="https://who-dat.as93.net/">who-dat.as93.net</a></b><br />
</p>

---

<details>
  <summary>Contents</summary>
  
- [About this Fork](#about-this-fork)
- [MCP Server Usage](#mcp-server-usage)
  - [Building the Server](#building-the-server)
  - [Running the Server](#running-the-server)
  - [Implementation in Cline, Cursor, and Claude Apps](#implementation-in-cline-cursor-and-claude-apps)
- [Original Project](#original-project)
- [Contributing](#contributing)
- [License](#license)

</details>

## About this Fork

This is a fork of the original [Who-Dat](https://github.com/Lissy93/who-dat) project, modified to run as an MCP (Model Context Protocol) server. This allows AI assistants like Cline, Cursor, and Claude to use the WHOIS lookup functionality as a tool.

## MCP Server Usage

### Building the Server

To build the MCP server, run the following command from the root of the project directory:

```bash
go build -o build/mcp-server mcp/main.go
```

This will create an executable file at `build/mcp-server`.

### Running the Server

You can run the server directly from your terminal:

```bash
./build/mcp-server
```

When run, the server will listen for requests on standard input and send responses to standard output, as per the MCP specification for stdio transport.

### Implementation in Cline, Cursor, and Claude Apps

To use this MCP server with your AI assistant, you need to add it to your MCP configuration file. This is typically a JSON file that lists the available MCP servers.

Here is an example configuration for adding the `who-dat-mcp` server. You will need to replace `/path/to/who-dat-mcp` with the actual absolute path to this project's directory on your system.

**Example `mcp_servers.json` configuration:**

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
```

**Instructions for specific applications:**

*   **Cline:** Add the above JSON object to your `cline_mcp_settings.json` file.
*   **Cursor:** Add the above JSON object to your MCP configuration within Cursor's settings.
*   **Claude Apps:** If you are using a Claude-compatible application that supports MCP, add the server configuration to the appropriate settings file.

Once configured, your AI assistant will be able to use the `get_whois` tool provided by this server.

---

## Original Project

This project is based on the excellent [Who-Dat](https://github.com/Lissy93/who-dat) by Alicia Sykes. For information about the original API, deployment options (like Docker and Vercel), and other details, please refer to the [original README.md](https://github.com/Lissy93/who-dat/blob/main/README.md).

---

## Contributing

Contributions to this MCP version are welcome! Please follow the [Code of Conduct](https://github.com/Lissy93/who-dat/blob/main/.github/CODE_OF_CONDUCT.md).

If you wish to contribute to the original project, please see the [contribution guidelines](https://github.com/Lissy93/who-dat#contributing) in the original repository.

---

## License

This project is licensed under the MIT License. The original copyright and license notice are included in the `LICENSE` file.

> _**[2b3pro/who-dat-mcp](https://github.com/2b3pro/who-dat-mcp)** is a fork of **[Lissy93/Who-Dat](https://github.com/Lissy93/who-dat)**, licensed under [MIT](https://github.com/Lissy93/who-dat/blob/HEAD/LICENSE) © [Alicia Sykes](https://aliciasykes.com) 2024._
