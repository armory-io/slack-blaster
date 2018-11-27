# slack-blaster

## Install

1. Download the release from our [releases page](https://github.com/armory-io/slack-blaster/releases).
2. Copy the binary into you `$PATH` - `cp ~/Downloads/slack-blaster /usr/local/bin`
3. Make the binary executable - `chmod +x /usr/local/bin/slack-blaster`

## Usage

1. Set your `SLACK_TOKEN` environment variable - `export SLACK_TOKEN={slack-api-token}`
2. Create a file with a list of channels to send your message to
3. Create a Markdown file which contains your message content
4. Run `slack-blaster blast --channels-list channels.txt --message-file message.md --for-real` (the `--for-real` flag will actually send the message)
