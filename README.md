# slack-blaster

## Install

1. Download the release from our [releases page](https://github.com/armory-io/slack-blaster/releases).
2. Copy the binary into you `$PATH` - `cp ~/Downloads/slack-blaster /usr/local/bin`
3. Make the binary executable - `chmod +x /usr/local/bin/slack-blaster`

## Configure Slack Token

In your web browser, go to [https://cloud-armory.slack.com](https://cloud-armory.slack.com)
and then view page source. You will be looking for `api_token`. Use that token
as your `SLACK_TOKEN`.

Set your `SLACK_TOKEN` environment variable:

```bash
export SLACK_TOKEN={slack-api-token}
```

## Getting List of Channels

To get a list of all the channels:

```bash
slack-blaster channels
```

To get a list of channels matching a pattern:

```bash
slack-blaster channels --pattern {pattern}
```

## Sending a message to channels

1. Create a file with a list of channels to send your message to
2. Create a Markdown file which contains your message content
3. Run slack blaster as specified below

```bash
slack-blaster blast --channels-list channels.txt --message-file message.md --for-real
```

(the `--for-real` flag will actually send the message)

