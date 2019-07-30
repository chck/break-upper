# synr: A script that leaving from your unnecessary chat rooms
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/722a6a380dd44b3f9635fbda8d38fdbe)](https://app.codacy.com/app/chck/synr?utm_source=github.com&utm_medium=referral&utm_content=chck/synr&utm_campaign=Badge_Grade_Dashboard)
[![Build Status](https://travis-ci.com/chck/synr.svg?branch=master)](https://travis-ci.com/chck/synr)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/chck/synr/blob/master/LICENSE)

## Requirements
```bash
go==1.12.*
```

## Installation
```bash
% go get -u github.com/chck/synr
```

## Setup
```bash
# Set tokens in config/secrets.yaml
# slack: https://api.slack.com/docs/oauth-test-tokens
# chatwork: http://developer.chatwork.com/ja/authenticate.html
vi config/secrets.yaml
---
tokens:
  chatwork: YOUR_CHARWORK_TOKEN
  slack: YOUR_SLACK_TOKEN
```

## Usage
```bash
# Show usage with options
% synr help

# Dry-run
% synr -c slack -d

# Run
% synr -c slack

# Run with selected months elapsed from last update (default: 1)
% synr -c slack -m 3

# Supported chats are slack and chatwork
% synr -c chatwork -m 3
```
