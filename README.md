
# SSE-on-Golang

Demo implementation of Server-Sent Events (SSE) in Golang.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Installation](#installation)
5. [Usage](#usage)
6. [Architecture](#architecture)
7. [Examples](#examples)
8. [Contributing](#contributing)
9. [License](#license)

## Introduction

Server-Sent Events (SSE) is a server push technology enabling a client to receive automatic updates from a server via an HTTP connection. This project is a demo implementation of SSE in Golang, showcasing how to build a simple SSE server and client.

## Features

- **Real-time Updates**: Push notifications to clients in real-time.
- **Simple Implementation**: Easy to understand and extend.
- **Golang**: Utilizes the powerful and efficient Go programming language.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your machine. You can download it from [golang.org](https://golang.org/dl/).
- Basic understanding of Golang and SSE concepts.

## Installation

Follow these steps to set up and run the project:

1. Clone the repository:

    ```sh
    git clone https://github.com/maen08/sse-on-golang.git
    cd sse-on-golang
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Build the project:

    ```sh
    go build -o sse-server main.go
    ```

4. Run the server:

    ```sh
    ./sse-server
    ```

## Usage

To use the SSE server, follow these steps:

1. Start the SSE server (as shown in the installation section).
2. Open a web browser or an HTTP client and navigate to `http://localhost:8080/events`.
3. The client will automatically receive updates from the server.

## Architecture

The project consists of the following components:

- **Server**: Handles client connections and sends events.
- **Client**: Listens for events from the server and processes them.
- **Event Handler**: Manages the creation and broadcasting of events to connected clients.

## Examples

Here is an example of how to use the SSE server:

1. Start the server:

    ```sh
    ./sse-server
    ```

2. Open a web browser and navigate to `http://localhost:8080/events`.

3. The server will send real-time updates to the client, which will be displayed in the browser.

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a pull request.

Please ensure your code adheres to the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more information.
