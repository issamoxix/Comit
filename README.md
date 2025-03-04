# Comit

## Overview
Comit is a tool that automatically generates commit messages based on the staged diff files in a Git repository.
It analyzes the changes and provides meaningful commit messages, 
making your commit history more structured and informative.

## Features
- Analyzes staged diff files to generate commit messages
- Provides structured and meaningful commit messages
- Helps maintain consistency in commit history
- Generate Branch Names
- Chat with an Agent

## Installation
Download [Here](https://github.com/issamoxix/Comit/releases)  

| System | Excutable | 
| :---         |     :---:      | 
| MacOS  | comit     | 
| Windows | comit.exe | 
| Linux  | comitl     |


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