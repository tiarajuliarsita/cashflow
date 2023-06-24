# Project Name

Simple ATM API

## Description

This project is an implementation of a simple ATM API using the Gin framework and a SQL database. The API provides endpoints that allow users to perform basic ATM operations such as checking balance, cash withdrawal, and fund transfers.

## Features

- User balance inquiry
- Cash withdrawal
- Account-to-account transfers
- User management (add, update, delete)

## Technology Used

- Programming Language: Go
- Framework: [Gin](https://github.com/gin-gonic/gin)
- Database: SQL (e.g., MySQL)
- ORM: [GORM](https://gorm.io)

## Installation

1. Make sure Go is installed on your system. Installation instructions can be found in the [official Go documentation](https://golang.org/doc/install).

2. Clone this repository to your local directory:
   git clone https://github.com/username/repository.git

3. Navigate to the project directory:
  cd repository

4. Install the required dependencies using the command:
  go mod download

5. Database Configuration:
  Create a database on your SQL server.
  
6. Run the application using the command:
  go run main.go
  The application will run at http://localhost:8080.
