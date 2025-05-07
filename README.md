# Telegram Bot with Message Reversal and Story Generation

This is a Telegram bot written in Go that provides two main functionalities:
1. Reverses any incoming message
2. Generates short sci-fi stories using OpenAI's GPT-3.5 API

## Features

- **Message Reversal**: The bot will respond to any message by sending back its reversed version
- **Story Generation**: Use the `/story` command to receive a randomly generated sci-fi story (max 400 characters)
- **Environment Variables**: Secure configuration using .env file
- **Docker Support**: Easy deployment using Docker and Docker Compose

## Prerequisites

- Docker and Docker Compose
- Telegram Bot Token (from [@BotFather](https://t.me/BotFather))
- OpenAI API Key

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd GenAI-basics-HW2-3
```

2. Create a `.env` file based on `.env.example`:
```bash
cp .env.example .env
```

3. Edit the `.env` file and add your credentials:
```
TELEGRAM_BOT_TOKEN=your_telegram_bot_token_here
OPENAI_API_KEY=your_openai_api_key_here
```

## Running the Bot

### Using Docker Compose (Recommended)

1. Build and run the bot in detached mode:
```bash
docker-compose up -d
```

2. Check the logs:
```bash
docker-compose logs -f
```

3. Stop the bot:
```bash
docker-compose down
```

### Local Development (Alternative)

1. Install Go 1.21 or higher
2. Install dependencies:
```bash
go mod download
```

3. Run the bot:
```bash
go run main.go
```

## Testing

Run the unit tests:
```bash
go test ./utils
```

## Project Structure

```
.
├── main.go              # Main bot implementation
├── utils/
│   ├── string_utils.go  # String reversal function
│   └── string_utils_test.go  # Unit tests
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Contributing

Feel free to submit issues and enhancement requests!