# Standup Bot

A simple bot to request for:

1. What was done yesterday
2. What is intended to be done today
3. Any blockers to getting things done

# Getting started

## Setting up a Slack App

1. Go to https://api.slack.com/apps
   1. Click on **Create New App**
   2. Select **From scratch**
   3. Give your app a name
   4. Select the workspace for the app to reside in
   5. Note the app details:
      1. App ID => `SLACK_APP_ID`
      2. Client ID => `SLACK_CLIENT_ID`
      3. Client Secret => `SLACK_CLIENT_SECRET`
      4. Signing Secret => `SLACK_SIGNING_SECRET`
2. Navigate to **App Home** from the left navigation menu
   1. Scroll down to **Show Tabs**
   2. Check the checkbox beside **Allow users to send Slash commands and messages from the messages tab**
3. Navigate to **Event Subscriptions** from the left navigation menu
   1. Toggle **Enable Events** so that it is enabled
   2. Subscribe to the following events under **Subscribe to events on behalf of users**:
      1. `message.im`
4. Navigate to **OAuth & Permissions** from the left navigation menu
   1. Scroll down to **Scopes**
   2. Add the following **Bot Token Scopes**:
      1. `chat:write`
   3. Scroll up to **OAuth Tokens > Workspace Installation**
   4. Click on **Install to ${WORKSPACE}**
5. You should be redirected to a page asking for permissions, hit **Allow**
6. Save the **Bot User OAuth Token** => `SLACK_BOT_USER_OAUTH_TOKEN`
7. Navigate to **Socket Mode** from the left navigation menu
   1. Toggle Socket Mode so that it is enabled
   2. In the popup modal, give the token a name, `dev` if you're using this for development
   3. Save the generated token => `SLACK_BOT_SOCKET_TOKEN`

## Running locally

1. Ensure that the following environment variables are populated in your `.envrc`
   1. `SLACK_APP_ID`
   2. `SLACK_CLIENT_ID`
   3. `SLACK_CLIENT_SECRET`
   4. `SLACK_SIGNING_SECRET`
   5. `SLACK_BOT_USER_OAUTH_TOKEN`
   5. `SLACK_BOT_SOCKET_TOKEN`
2. Run `go mod vendor` to bring in dependencies
3. Run `direnv allow .` to allowlist the `.envrc` file
4. Run `go run .` to start the bot

# Feature Roadmap

## V1

### Functional requirements

- [ ] Ask a user the 3 questions at a daily timing
- [ ] Schedule for asking user the 3 questions is configurable by each user
- [ ] Automatically post a user's response into a channel at a scheduled daily timing
- [ ] Schedule for posting a user's response is configurable by each user
- [ ] Channel to post a user's update to is configurable by each user
- [ ] User can ask bot about mentions of a technology and be pointed to other users' standup notes

### Non-functional requirements

- [ ] Bot should be deployable on an EC2 instance with Docker Compose
- [ ] Bot should be deployable in a Kubernetes cluster with a Helm chart
