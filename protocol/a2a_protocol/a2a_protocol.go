package a2a_protocol

import (
	"errors"

	"github.com/windlant/protocol/types/agent_types"
	"github.com/windlant/protocol/types/skill_types"
)

// A2AProtocol 定义 A2A 协议的核心常量和接口
type A2AProtocol struct{}

// HTTP 路径常量
const (
	// Agent 发现和元数据
	AgentCardPath = "/.well-known/agent.json"

	// 任务管理
	TasksPath = "/tasks"
	TaskPath  = "/tasks/{taskId}"

	// 消息通信
	MessagesPath = "/messages"
	MessagePath  = "/messages/{messageId}"
)

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

// SendMessageRequest 发送消息的请求格式
type SendMessageRequest struct {
	ToAgentID string                  `json:"toAgentId"`
	Type      agent_types.MessageType `json:"type"`
	Content   interface{}             `json:"content"`
}

// SendMessageResponse 发送消息的响应格式
type SendMessageResponse struct {
	MessageID string `json:"messageId"`
}

// NewA2AProtocol 创建 A2A 协议实例
func NewA2AProtocol() *A2AProtocol {
	return &A2AProtocol{}
}

// ValidateAgentCard 验证 AgentCard 的基本有效性
func (p *A2AProtocol) ValidateAgentCard(card *agent_types.AgentCard) error {
	if card.AgentID == "" {
		return errors.New("agentId is required")
	}
	if card.Name == "" {
		return errors.New("name is required")
	}
	if card.URL == "" {
		return errors.New("url is required")
	}

	// 验证技能（如果有）
	for _, skill := range card.Skills {
		if skill.Name == "" {
			return errors.New("skill name is required")
		}
		if skill.Description == "" {
			return errors.New("skill description is required")
		}
	}

	return nil
}

// ValidateMessage 验证消息的基本有效性
func (p *A2AProtocol) ValidateMessage(msg *agent_types.Message) error {
	if msg.MessageID == "" {
		return errors.New("messageId is required")
	}
	if msg.FromAgentID == "" {
		return errors.New("fromAgentId is required")
	}
	if msg.ToAgentID == "" {
		return errors.New("toAgentId is required")
	}
	if msg.Type == "" {
		return errors.New("message type is required")
	}
	if msg.Content == nil {
		return errors.New("message content is required")
	}

	return nil
}
