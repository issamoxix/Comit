# Comit

## Overview
Comit is a powerful tool that automatically generates commit messages based on staged changes in a Git repository. 
It analyzes diffs to create structured and meaningful commit messages, helping maintain consistency in your commit history. <br>
Additional features include automatic branch name generation and an AI-powered chat assistant for coding-related queries. <br>
Simplify your Git workflow with Comit!

## Features
- Analyzes staged diff files to generate commit messages
- Provides structured and meaningful commit messages
- Helps maintain consistency in commit history
- Generate Branch Names
- Chat with an Agent

## Installation

| System | Excutable | 
| :---         |     :---:      | 
| MacOS  | comit     | 
| Windows | comit.exe | 
| Linux  | comitl     |

Download [Here](https://github.com/issamoxix/Comit/releases)  

## Usage
To generate Commit message, stage the changes and the run the command below :
```bash
comit
# output
# feat: allow provided config object to extend other configs
# fix: handle out of range errors
```
To Generate branch name :
```bash
comit -b "Add a New Feature to the User Profile Page"
# output
# feature/add-new-feature-user-profile-page
# bugfix/save-button-not-working-settings-page
```
To Chat with an Agent (for now we using gpt-4o):
```bash
comit -c "how to run postgres container in docker"    
# output
# 1. Pull the Postgres image:
#    ===
#    docker pull postgres
#    ===

# 2. Run the Postgres container:
#    ===
#    docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
#    ===

# Replace `mysecretpassword` with your desired password.
```
