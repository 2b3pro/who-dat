package main

import (
	"context"
	"encoding/json"

	"github.com/lissy93/who-dat/lib"
	"github.com/modelcontextprotocol/go-sdk/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// WhoisMultiArgs defines the arguments for the get_whois_multi tool.
type WhoisMultiArgs struct {
	Domains []string `json:"domains" mcp:"the list of domains to lookup"`
}

// WhoisMultiOutput defines the expected structure of the get_whois_multi tool's output.
type WhoisMultiOutput map[string]interface{}

// GetWhoisMultiTool is the handler for the get_whois_multi tool.
func GetWhoisMultiTool(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[WhoisMultiArgs]) (*mcp.CallToolResultFor[json.RawMessage], error) {
	results := make(map[string]interface{})
	for _, domain := range params.Arguments.Domains {
		whoisInfo, err := lib.GetWhois(domain)
		if err != nil {
			results[domain] = map[string]string{"error": err.Error()}
		} else {
			results[domain] = whoisInfo
		}
	}

	resultJSON, err := json.Marshal(results)
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
		Name:        "get_whois_multi",
		Description: "Get WHOIS information for a list of domains",
		OutputSchema: &jsonschema.Schema{ // Directly initialize jsonschema.Schema
			Type:                 "object",
			Title:                "WHOIS Multi Output",
			Description:          "A map of domain names to their WHOIS information or error messages.",
			AdditionalProperties: nil, // Set AdditionalProperties to allow any additional properties
		},
	}, GetWhoisMultiTool)

	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		// Log messages are omitted when using stdio transport to prevent interference with protocol messages.
		// Any errors will be communicated via the MCP protocol itself.
	}
}
