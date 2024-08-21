package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/zalando/go-keyring"
)

const (
	serviceName = "ANTHROPIC_API_KEY"
	userName    = "system"
)

func main() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alfred",
		Short: "AI-powered Kubernetes assistant",
		Long: `Alfred is an AI-powered assistant for Kubernetes, allowing you to ask questions using natural language.

Commands:
  ask      Ask a question about Kubernetes
  set-key  Set the Anthropic API key in the system keyring
  del-key  Delete the Anthropic API key from the system keyring`,
	}

	cmd.AddCommand(NewAskCmd())
	cmd.AddCommand(NewSetKeyCmd())
	cmd.AddCommand(NewDelKeyCmd())

	return cmd
}

func NewAskCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ask <question>",
		Short: "Ask a question about Kubernetes",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runAIQuery,
	}
}

func NewSetKeyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-key <api-key>",
		Short: "Set the Anthropic API key in the system keyring",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return keyring.Set(serviceName, userName, args[0])
		},
	}
}

func NewDelKeyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "del-key",
		Short: "Delete the Anthropic API key from the system keyring",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return keyring.Delete(serviceName, userName)
		},
	}
}

func runAIQuery(cmd *cobra.Command, args []string) error {
	question := strings.Join(args, " ")

	apiKey, err := keyring.Get(serviceName, userName)
	if err != nil {
		return fmt.Errorf("failed to get API key from keyring: %w", err)
	}

	err = os.Setenv("ANTHROPIC_API_KEY", apiKey)
	if err != nil {
		return fmt.Errorf("failed to set API key environment variable: %w", err)
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
