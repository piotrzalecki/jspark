# jspark

> ⚠️ **Warning: This is a Proof of Concept (POC) project**
>
> This project is currently in POC stage and requires further development before being production-ready. Use at your own risk.

CLI tool for creating Jira tickets in cooperation with LLMs.

## Description

jspark is a command-line tool that helps you create Jira tickets using the power of Large Language Models (LLMs). It provides an interactive way to generate and manage Jira tickets with AI assistance.

## Prerequisites

- [Ollama](https://ollama.ai/) installed on your system
- OpenAI API key (if using OpenAI models)

## Model Setup

### Local Model (Ollama)

To use the local model, you need to build it first using Ollama. The model configuration is provided in the `Modefile.llama3` file.

```bash
ollama create jspark -f Modefile.llama3
```

### OpenAI Integration

To use OpenAI models, you need to set up your API key as an environment variable:

```bash
export OPENAI_API_KEY='your-api-key-here'
```

## Usage

jspark provides two main commands for generating Jira tickets: `ollama` and `gpt`. Both commands can read input either directly from the command line or from your clipboard.

### Using Ollama (Local Model)

```bash
# Using direct input
jspark ollama "Your ticket description here"

# Using clipboard content
jspark ollama --clipboard

# Specify a different model
jspark ollama --model custom-model "Your ticket description here"
```

### Using GPT (OpenAI)

```bash
# Using direct input
jspark gpt "Your ticket description here"

# Using clipboard content
jspark gpt --clipboard
```

### Interactive Mode

After generating a ticket, jspark will:

1. Display the proposed ticket title and summary
2. Ask if you want to modify the title
3. Ask if you want to add additional content to the summary

Example output:

```
----------------------
TICKET PROPOSAL

TITLE:        Example Ticket Title
SUMMARY:      Example ticket summary describing the issue

----------------------

Do you want to change title?(Y/N)
Do you want to add something to summary? (Y/N)
```

## License

MIT License - See [LICENSE](LICENSE) file for details
