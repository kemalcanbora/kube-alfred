package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func main() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alfred <question>",
		Short: "AI-powered Kubernetes assistant",
		Long:  `Alfred is an AI-powered assistant for Kubernetes, allowing you to ask questions using natural language.`,
		RunE:  runAIQuery,
	}

	return cmd
}

func runAIQuery(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a question about Kubernetes")
	}

	question := strings.Join(args, " ")
	err := os.Setenv("ANTHROPIC_API_KEY", "ANTHROPIC_API_KEY")
	if err != nil {
		return fmt.Errorf("failed to set API key: %w", err)
	}

	llm, err := anthropic.New(
		anthropic.WithModel("claude-3-5-sonnet-20240620"),
	)
	if err != nil {
		return fmt.Errorf("failed to create Anthropic client: %w", err)
	}

	ctx := context.Background()
	prompt := fmt.Sprintf("You are a Kubernetes expert. Please answer the following question about Kubernetes: %s", question)

	_, err = llms.GenerateFromSinglePrompt(ctx, llm, prompt,
		llms.WithTemperature(0.7),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to generate response: %w", err)
	}

	return nil
}
