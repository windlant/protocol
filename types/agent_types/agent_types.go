package agent_types

import (
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
