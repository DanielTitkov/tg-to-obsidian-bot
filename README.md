# tg-to-obsidian-bot

Bot that converts to messages from tg channel to obsidian notes

## Requirements

- docker
- make

## Usage

Provide environment variables within `.env` file it the project root.

```
TG_TOKEN=9999999:AAAAAAVVVVVVVVVGGGGGG
NOTES_PATH=/home/username/Vault/_inbox/
TEMPLATE_PATH=/home/username/Vault/templates/
TEMPLATE_FILE=telegram.md
TIMEOUT=180
```

After this is done just run `make up` to set up the container with bot binary. It will have access to your host directories you provided. Template path is read only. 

## Environment args

- TG_TOKEN - your bot's token
- NOTES_PATH - path where you want bot to save your notes
- TEMPLATE_PATH - path where you have your template
- TEMPLATE_FILE - template file name
- TIMEOUT - time in seconds after which bot will delete processed messages from chat

## Template

### Example

```md
# {{title}}
Created: **{{date}} {{time}}**
Tags: #seed 
Themes: 

{{content}}

## Refs
```

### Placeholders

- {{content}} - message from chat with the bot
- {{time}} - time 15:04
- {{date}} - date 2006-01-02
- {{title}} - post title Telegram Note 2006-01-02 15:04:05