# gin-app

## Introduction

This Go Gin application provides two main functionalities:
1. **Exchange Rate Retrieval**: Fetches the current exchange rate of USD to UAH.
2. **Subscription Management**: Allows users to subscribe to daily email updates with the current exchange rate.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Endpoints](#endpoints)
- [Configuration](#configuration)
- [Documentation](#documentation)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/hryhorskyi/gin-app
    cd gin-app
    ```

2. Create and configure the `.env` file(check .env.example file):
    ```plaintext
    DB_HOST=your-db-host
    DB_PORT=your-db-port
    DB_USER=your-db-user
    DB_NAME=your-db-name
    DB_PASSWORD=your-db-password

    EMAIL_FROM=your-email@gmail.com
    EMAIL_PASSWORD=your-email-password
    SMTP_HOST=smtp.gmail.com
    SMTP_PORT=587
    ```

3. Run the application with Docker Compose:
    ```bash
    docker-compose up --build
    ```

## Usage

- **Get Exchange Rate**: Retrieve the current USD to UAH exchange rate.
- **Subscribe to Updates**: Subscribe to receive daily email updates with the current exchange rate.

## Features

- Fetch current USD to UAH exchange rate via a third-party API.
- Manage subscriptions for daily email updates with the current exchange rate.
- Daily job scheduler to send emails at 08:00 UTC.

## Endpoints

- **Exchange Rate Retrieval Endpoint**
  - **GET** `/api/rate`: Fetches the current USD to UAH exchange rate and returns it as a JSON response.
  
- **Subscription Endpoint**
  - **POST** `/api/subscribe`: Allows users to subscribe to daily exchange rate updates by providing their email address. The email is stored in the database, and the user will receive daily emails with the current exchange rate.

## Configuration

Ensure the `.env` file in the root directory is populated with your environment variables as specified in the installation section.

## Documentation

Swagger documentation is available once the application is running:
- Access it at `/swagger/index.html`.
