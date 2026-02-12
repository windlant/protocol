package registry_protocol

import (
	"time"

	"github.com/windlant/protocol/types/agent_types"
)

// RegistryProtocol 定义 Agent Registry 的核心常量和类型
type RegistryProtocol struct{}

// HTTP 路径常量
const (
	// Agent 注册和管理
	RegisterPath    = "/register"
	DeregisterPath  = "/deregister"
	AgentsListPath  = "/agents"
	AgentDetailPath = "/agents/{agentId}"
)

// RegisterRequest Agent 注册请求
type RegisterRequest struct {
	AgentCard  agent_types.AgentCard `json:"agentCard"`
	WebhookURL string                `json:"webhookUrl"` // 用于接收更新通知的回调 URL
}

// RegisterResponse Agent 注册响应
type RegisterResponse struct {
	AgentID   string    `json:"agentId"`
	ExpiresAt time.Time `json:"expiresAt"` // 注册过期时间（用于心跳续期）
}

// DeregisterRequest Agent 注销请求
type DeregisterRequest struct {
	AgentID string `json:"agentId"`
}

// AgentsListResponse 获取所有已注册 Agent 的响应
type AgentsListResponse struct {
	Agents []agent_types.AgentCard `json:"agents"`
}

// AgentHealthCheckRequest Registry Server 发送给 Agent 的健康检查请求
type AgentHealthCheckRequest struct {
	Timestamp time.Time `json:"timestamp"`         // 健康检查发起时间
	AgentID   string    `json:"agentId,omitempty"` // Agent 标识（可选）
}

// AgentHealthCheckResponse Agent 返回给 Registry Server 的健康检查响应
type AgentHealthCheckResponse struct {
	Status    string                 `json:"status"`              // "healthy" 或 "unhealthy"
	Timestamp time.Time              `json:"timestamp"`           // 响应时间
	AgentID   string                 `json:"agentId"`             // Agent 标识
	Message   string                 `json:"message,omitempty"`   // 额外状态信息（如错误详情）
	AgentCard *agent_types.AgentCard `json:"agentCard,omitempty"` // 可选：更新后的 AgentCard（如果技能列表有变化）
}

// AgentUpdateNotification Agent 更新通知（通过 webhook 发送）
type AgentUpdateNotification struct {
	Type      NotificationType       `json:"type"`                // 通知类型
	Timestamp time.Time              `json:"timestamp"`           // 通知时间
	AgentID   string                 `json:"agentId"`             // 相关 Agent ID
	AgentCard *agent_types.AgentCard `json:"agentCard,omitempty"` // 更新后的 AgentCard（仅在 AGENT_UPDATED 或 AGENT_REGISTERED 时提供）
}

// NotificationType 通知类型枚举
type NotificationType string

const (
	NotificationTypeAgentRegistered   NotificationType = "AGENT_REGISTERED"   // 新 Agent 注册
	NotificationTypeAgentUpdated      NotificationType = "AGENT_UPDATED"      // Agent 信息更新
	NotificationTypeAgentDeregistered NotificationType = "AGENT_DEREGISTERED" // Agent 注销
)
