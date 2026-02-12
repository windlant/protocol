package skill_types

// SkillArguments 表示技能调用的输入参数
type SkillArguments map[string]interface{}

// SkillFunc 是所有技能实现必须遵循的函数签名
type SkillFunc func(SkillArguments) (string, error)

// SkillParameter 描述技能的一个参数
type SkillParameter struct {
	Type        string `json:"type"`        // 例如 "string"、"number"、"object"
	Description string `json:"description"` // 参数用途说明
	Required    bool   `json:"required"`    // 该参数是否必需（仅用于文档，实际 required 列表在 Schema 中）
}

// SkillSchema 描述技能期望的输入结构（遵循 JSON Schema 规范）
type SkillSchema struct {
	Type       string                    `json:"type"`       // 通常为 "object"
	Properties map[string]SkillParameter `json:"properties"` // 参数定义
	Required   []string                  `json:"required"`   // 必需参数的名称列表
}

// SkillDefinition 包含 LLM 使用技能所需的全部元数据
type SkillDefinition struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  SkillSchema `json:"parameters"`
	Function    SkillFunc   `json:"-"` // 不参与 JSON 序列化，仅在本地执行时使用
}
