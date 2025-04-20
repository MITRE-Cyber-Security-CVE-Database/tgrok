# tgrok

![Go](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## Overview

`tgrok` is a lightweight, terminal-based AI client for interacting with the Grok 3 API, designed for minimal resource usage. It supports command-line prompts, piped input, and a simple interactive mode, making it ideal for cybersecurity professionals and developers who need quick access to AI-driven insights. Developed under the MITRE-Cyber-Security-CVE-Database organization, `tgrok` is optimized for environments where low memory usage is critical, such as in vulnerability analysis workflows.

## Features

- **Lightweight**: Minimal dependencies and no heavy UI frameworks, ensuring low RAM usage.
- **Flexible Input**: Supports command-line prompts, piped input, and interactive mode.
- **Quiet Mode**: Optional flag to disable loading animations for scripting.
- **Error Handling**: Robust error messages for API, input, and network issues.
- **MIT License**: Open-source and free to use, modify, and distribute.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or later
- A valid Grok 3 API key (obtain from [xAI API](https://x.ai/api))

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/MITRE-Cyber-Security-CVE-Database/tgrok.git
   cd tgrok
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the binary:
   ```bash
   go build -o tgrok
   ```

4. (Optional) Install globally:
   ```bash
   sudo mv tgrok /usr/local/bin/
   ```

## Usage

### Command-Line Prompt
```bash
tgrok "What is the capital of France?"
```

### Piped Input
```bash
echo "Explain gravity" | tgrok
```

### Interactive Mode
```bash
tgrok
```
Enter prompts at the `You>` prompt, and press `Ctrl+C` or type `exit` to quit.

### Flags
- `-h, --help`: Show help message
- `-v, --version`: Show version (1.0.0)
- `-q, --quiet`: Disable loading animation
- `-key string`: Grok 3 API key (or set `AI_API_KEY`)
- `-provider string`: AI provider (default: grok, or set `AI_PROVIDER`)

### Example with Quiet Mode
```bash
tgrok -q "Define AI"
```

### Environment Variables
Set these for convenience:
```bash
export AI_API_KEY="your-api-key"
export AI_PROVIDER="grok"
```

## API Configuration

The tool uses a placeholder API endpoint (`https://api.x.ai/v1/grok`). To use `tgrok`, obtain an API key and endpoint from [xAI API](https://x.ai/api). Update the `Url` field in `main.go` or pass the correct endpoint via a configuration if needed.

**Note**: The response parsing in `helper.go` assumes a JSON response with a `response` field. Adjust the parsing logic based on the actual API specification.

## Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please follow the [Code of Conduct](CODE_OF_CONDUCT.md) and ensure your code adheres to the project's style guidelines.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For issues or questions, open an issue on the [GitHub repository](https://github.com/MITRE-Cyber-Security-CVE-Database/tgrok) or contact the MITRE-Cyber-Security-CVE-Database team.

## Acknowledgments

- Built with [Go](https://golang.org/) and [fatih/color](https://github.com/fatih/color).
- Inspired by the need for lightweight tools in cybersecurity workflows, aligned with MITRE’s mission to advance vulnerability management.[](https://github.com/MITRE-Cyber-Security-CVE-Database/mitre-cve-database)

---
Copyright © 2025 GulfOfAmerica
