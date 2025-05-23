package llm

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmc/langchaingo/llms"
)

func TestNewLLM(t *testing.T) {
	tests := []struct {
		name    string
		model   string
		wantErr bool
	}{
		{
			name:    "创建有效的 LLM 实例",
			model:   "llama3.1:8b",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			llm, err := NewLLM(tt.model)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, llm)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, llm)
				assert.Equal(t, tt.model, llm.model)
				assert.NotNil(t, llm.client)
			}
		})
	}
}

func TestLLM_GetStreamAnswer(t *testing.T) {
	llm, err := NewLLM("llama3.1:8b")
	if err != nil {
		t.Fatalf("创建 LLM 实例失败: %v", err)
	}

	tests := []struct {
		name     string
		ctx      context.Context
		history  []Message
		question string
		wantErr  bool
	}{
		{
			name:     "简单问题测试",
			ctx:      context.Background(),
			history:  []Message{},
			question: "你好",
			wantErr:  false,
		},
		{
			name: "带历史消息的对话测试",
			ctx:  context.Background(),
			history: []Message{
				{
					Role:    llms.ChatMessageTypeHuman,
					Content: "你好, 我是小明, 从现在开始你就是小王",
				},
				{
					Role:    llms.ChatMessageTypeAI,
					Content: "你好！从现在开始我就是小王, 有什么我可以帮你的吗？",
				},
			},
			question: "请介绍一下你自己",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch, errCh := llm.GetStreamAnswer(tt.ctx, tt.history, tt.question)
			select {
			case err := <-errCh:
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			case <-ch:
				assert.NotNil(t, ch)

				// 读取响应
				var response string
				for msg := range ch {
					response += msg
					fmt.Print(msg)
				}
				fmt.Println()
				assert.NotEmpty(t, response)
			}
		})
	}
}
