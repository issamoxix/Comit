# Comit

## Overview
Comit is a tool that automatically generates commit messages based on the staged diff files in a Git repository.
It analyzes the changes and provides meaningful commit messages, 
making your commit history more structured and informative.
[API](https://github.com/issamoxix/comitApi)

## Features
- Analyzes staged diff files to generate commit messages
- Provides structured and meaningful commit messages
- Helps maintain consistency in commit history
- Generate Branch Names

## Installation
Download the comit.exe or comit (for mac from Realeases)

## Usage
To generate Commit message, stage the changes and the run the command below :
```bash
comit
# output
# feat: commit message here
# fix: commit message here
```
To Generate branch name :
```bash
comit -b "branch context here"
# output
# feat/branch-name-here
# bugfix/branch-name-here
```

## Example Output

```
feat: Add user authentication flow
fix: Resolve issue with payment processing
refactor: Optimize database queries
```
