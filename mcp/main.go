package main

import (
	"context"
	"encoding/json"

	"github.com/lissy93/who-dat/lib"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// WhoisArgs defines the arguments for the get_whois tool.
type WhoisArgs struct {
	Domain string `json:"domain" mcp:"the domain to lookup"`
}

// GetWhoisTool is the handler for the get_whois tool.
func GetWhoisTool(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[WhoisArgs]) (*mcp.CallToolResultFor[json.RawMessage], error) {
	whoisInfo, err := lib.GetWhois(params.Arguments.Domain)
	if err != nil {
		return nil, err
	}

	resultJSON, err := json.Marshal(whoisInfo)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[json.RawMessage]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(resultJSON)},
		},
	}, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "who-dat"}, nil)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_whois",
		Description: "Get WHOIS information for a domain",
	}, GetWhoisTool)

	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		// Log messages are omitted when using stdio transport to prevent interference with protocol messages.
		// Any errors will be communicated via the MCP protocol itself.
	}
}
