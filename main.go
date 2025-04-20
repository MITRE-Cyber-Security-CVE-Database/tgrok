package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

const version = "1.0.0"

var bold = color.New(color.Bold)
var blue = color.New(color.FgBlue)

func main() {
	// Command-line flags
	provider := flag.String("provider", os.Getenv("AI_PROVIDER"), "AI provider (default: grok)")
	apiKey := flag.String("key", os.Getenv("AI_API_KEY"), "API key for Grok 3")
	isQuiet := flag.Bool("q", false, "Quiet mode (no loading animation)")
	isVersion := flag.Bool("v", false, "Show version")
	isHelp := flag.Bool("h", false, "Show help")
	flag.Parse()

	// Handle version and help flags
	if *isVersion {
		fmt.Println("grok-light", version)
		return
	}
	if *isHelp {
		showHelp()
		return
	}

	// Default provider to Grok if not specified
	if *provider == "" {
		*provider = "grok"
	}

	// Signal handling for clean exit
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-terminate
		os.Exit(0)
	}()

	// Initialize parameters
	params := Params{
		ApiKey:   *apiKey,
		Provider: *provider,
		Url:      "https://api.x.ai/v1/grok", // Placeholder; actual endpoint requires API access
	}

	// Handle command-line prompt or piped input
	prompt := flag.Arg(0)
	pipedInput := ""

	// Check for piped input
	stat, err := os.Stdin.Stat()
	if err != nil {
		PrintError(fmt.Sprintf("Error accessing standard input: %v", err))
		return
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			pipedInput += scanner.Text() + "\n"
		}
		if err := scanner.Err(); err != nil {
			PrintError(fmt.Sprintf("Error reading standard input: %v", err))
			return
		}
		pipedInput = strings.TrimSpace(pipedInput)
	}

	// Process input
	if len(prompt) > 0 {
		// Handle command-line prompt
		trimmedPrompt := strings.TrimSpace(prompt)
		if len(trimmedPrompt) < 1 {
			PrintError("You need to provide some text")
			fmt.Println(`Example: grok-light "What is the capital of France?"`)
			return
		}
		input := trimmedPrompt
		if len(pipedInput) > 0 {
			input += "\n\nContext:\n" + pipedInput
		}
		GetResponse(input, params, *isQuiet)
	} else if len(pipedInput) > 0 {
		// Handle piped input only
		GetResponse(pipedInput, params, *isQuiet)
	} else {
		// Interactive mode
		bold.Print("Interactive mode started. Press Ctrl+C to quit.\n\n")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			blue.Print("You> ")
			if !scanner.Scan() {
				break
			}
			input := strings.TrimSpace(scanner.Text())
			if input == "exit" {
				bold.Println("Exiting...")
				os.Exit(0)
			}
			if len(input) > 0 {
				GetResponse(input, params, *isQuiet)
			}
		}
		if err := scanner.Err(); err != nil {
			PrintError(fmt.Sprintf("Error reading input: %v", err))
		}
	}
}

func showHelp() {
	fmt.Println("grok-light - Lightweight terminal AI client for Grok 3")
	fmt.Println("\nUsage:")
	fmt.Println("  grok-light [flags] [prompt]")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help          Show this help message")
	fmt.Println("  -v, --version       Show version")
	fmt.Println("  -q, --quiet         Quiet mode (no loading animation)")
	fmt.Println("  -key string         API key for Grok 3 (or set AI_API_KEY)")
	fmt.Println("  -provider string    AI provider (default: grok, or set AI_PROVIDER)")
	fmt.Println("\nExamples:")
	fmt.Println(`  grok-light "What is the capital of France?"`)
	fmt.Println(`  echo "Explain gravity" | grok-light`)
	fmt.Println(`  grok-light -q "Define AI"`)
	fmt.Println("\nFor API access, visit https://x.ai/api")
}