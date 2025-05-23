package llm

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Role = llms.ChatMessageType

type Message struct {
	Role    Role
	Content string
}

type LLM struct {
	model  string
	client llms.Model
}

func NewLLM(model string) (*LLM, error) {
	llm, err := ollama.New(ollama.WithModel(model), ollama.WithServerURL("http://localhost:11434"))
	if err != nil {
		return nil, fmt.Errorf("创建 LLM 实例失败: %w", err)
	}

	return &LLM{model: model, client: llm}, nil
}

// 获取流式回答
func (l *LLM) GetStreamAnswer(ctx context.Context, history []Message, question string) (chan string, chan error) {
	ch := make(chan string)
	errCh := make(chan error)
	go func() {
		defer close(ch)
		defer close(errCh)
		// 将历史消息转换为 langchaingo 的消息格式
		messages := make([]llms.MessageContent, len(history))
		for i, msg := range history {
			messages[i] = llms.TextParts(msg.Role, msg.Content)
		}

		// 添加当前问题
		messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, question))

		// 使用 langchaingo 的流式生成
		_, err := l.client.GenerateContent(ctx, messages,
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				ch <- string(chunk)
				return nil
			}),
		)
		if err != nil {
			errCh <- err
		}
	}()
	return ch, errCh
}
