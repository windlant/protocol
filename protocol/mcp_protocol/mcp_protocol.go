package mcp_protocol

import "github.com/windlant/protocol/types/tools_types"

// MCP 方法常量
const (
	MCPMethodListTools = "list_tools"
	MCPMethodCallTool  = "call_tool"
)

// MCP 请求结构

type MCPListToolsRequest struct {
	Method  string `json:"method"`             // 必须为 "list_tools"
	AgentID string `json:"agent_id,omitempty"` // 可选的 Agent ID，用于权限控制
}

type MCPToolCallRequest struct {
	Method  string                 `json:"method"` // 必须为 "call_tool"
	Name    string                 `json:"name"`
	Args    map[string]interface{} `json:"arguments"`
	AgentID string                 `json:"agent_id,omitempty"` // 可选的 Agent ID，用于权限控制
}

// MCP 响应结构

type MCPListToolsResponse struct {
	Tools []tools_types.ToolDefinition `json:"tools"`
}

type MCPToolCallResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}
