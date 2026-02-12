package agent_types

import (
	"time"

	"github.com/windlant/protocol/types/skill_types"
)

// AgentCard 表示 Agent 的基本信息和能力
type AgentCard struct {
	AgentID     string                        `json:"agentId"`     // Agent 唯一标识符
	Name        string                        `json:"name"`        // Agent 名称
	Description string                        `json:"description"` // Agent 描述
	Skills      []skill_types.SkillDefinition `json:"skills"`      // Agent 支持的技能列表
	URL         string                        `json:"url"`         // Agent 服务地址
}

// Task 表示发送给 Agent 的任务
type Task struct {
	TaskID    string                     `json:"taskId"`
	AgentID   string                     `json:"agentId"` // 目标 Agent ID
	SkillID   string                     `json:"skillId"` // 要调用的技能 ID
	Input     skill_types.SkillArguments `json:"input"`   // 任务输入数据
	CreatedAt time.Time                  `json:"createdAt"`
}

// TaskResult 表示任务执行结果
type TaskResult struct {
	TaskID  string      `json:"taskId"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// MessageType 定义消息类型
type MessageType string

const (
	MessageTypeTaskRequest  MessageType = "task_request"  // 任务请求
	MessageTypeTaskResponse MessageType = "task_response" // 任务响应
	MessageTypeChat         MessageType = "chat"          // 普通聊天消息
	MessageTypeStatusUpdate MessageType = "status_update" // 状态更新
)

// Message 表示 Agent 之间的通信消息
type Message struct {
	MessageID   string                 `json:"messageId"`          // 消息唯一标识符
	FromAgentID string                 `json:"fromAgentId"`        // 发送方 Agent ID
	ToAgentID   string                 `json:"toAgentId"`          // 接收方 Agent ID
	Type        MessageType            `json:"type"`               // 消息类型
	Content     interface{}            `json:"content"`            // 消息内容
	Timestamp   time.Time              `json:"timestamp"`          // 时间戳
	Metadata    map[string]interface{} `json:"metadata,omitempty"` // 自定义元数据
}
