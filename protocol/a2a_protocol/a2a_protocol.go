package a2a_protocol

import (
	"time"

	"github.com/windlant/protocol/types/skill_types"
)

// A2AProtocol 定义 A2A 协议的核心常量和接口
type A2AProtocol struct{}

// HTTP 路径常量
const (
	// Agent 发现和元数据，目前使用注册中心，暂时用不到
	AgentCardPath = "/.well-known/agent.json"

	// 任务管理
	TasksPath = "/tasks"
	TaskPath  = "/tasks/{taskId}"

	// 消息通信
	MessagesPath = "/messages"
	MessagePath  = "/messages/{messageId}"
)

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
	TaskID   string      `json:"taskId"`
	Success  bool        `json:"success"`
	Artifact interface{} `json:"artifact,omitempty"`
	Error    string      `json:"error,omitempty"`
}

// CreateTaskRequest 创建任务的请求格式
type CreateTaskRequest struct {
	AgentID string                     `json:"agentId"`
	SkillID string                     `json:"skillId"`
	Input   skill_types.SkillArguments `json:"input"`
}

// CreateTaskResponse 创建任务的响应格式
type CreateTaskResponse struct {
	TaskID string `json:"taskId"`
}

type Part struct {
	Type    string      `json:"type"` // "text" 或 "file"或 "json"
	Content interface{} `json:"content"`
}

type Message struct {
	MessageID string `json:"messageId"`
	Role      string `json:"role"`
	Parts     []Part `json:"parts"`
}

// SendMessageRequest 发送消息的请求格式
type SendMessageRequest struct {
	ToAgentID string `json:"toAgentId"`
	Message   string `json:"message"`
}

// SendMessageResponse 发送消息的响应格式
type SendMessageResponse struct {
	Message string `json:"message"`
}
