# GitPulse - Github issues Tracker

![GitHub](https://img.shields.io/github/license/yourusername/github-issue-tracker)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yourusername/github-issue-tracker)

GitPulse Tracker is an open-source application that allows you to fetch and list open issues in repositories that have been active in the last 3 months. It provides a convenient way to keep track of recent issues across various GitHub projects.

## Features

- Fetch open issues from active GitHub repositories.
- Filter repositories based on their activity within the last 3 months. (TODO)
- Minimalistic and user-friendly interface.

## Configuration

The application requires a GitHub API key for authentication.
You can set this API key as an environment variable such as :
GITHUB_API_KEY=your_api_key_here

## Usage

Access the application in your web browser by navigating to client folder and runinng npm run dev.

Start the backend server by running : go run main.go in the server directory .

Explore the open issues from the active repositories.

## License

This project is licensed under the MIT License
