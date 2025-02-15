This repo is designed to track my steps inside GO's world. I'm studying through the OFFICIAL documentation [Tour of Go](https://go.dev/tour/welcome/1) and also I've created this public [Playlist on YouTube](https://www.youtube.com/watch?v=f6kdp27TYZs&list=PLI12PnJqTSg0Lz2n5-wL_SnWorVbJ7or9) that I'm constantly updating!

### Plan

Phase 1: Core Basics (1-2 Weeks)
- [x] Go Basics
    - [x] History of GO
    - [x] Learn Go syntax: variables, data types (strings, integers, booleans), and basic input/output.
    - [x] Study conditional statements (if, else, switch).
    - [x] Learn about loops (for loop, Go's only loop type), and collections (arrays, slices, and maps).
    - [x] Understand Go’s structure: packages, functions, and modules.
- [x] Goal: Write small programs that test each concept individually, like calculating simple math operations and playing with conditionals.
- [x] Project: Build a Simple Calculator
    - [x] Implement a command-line calculator that supports basic operations: addition, subtraction, multiplication, and division.
    - [x] Use input/output functions to accept numbers and operations from the user.
    - [x] Practice error handling for invalid inputs (e.g., division by zero).
- [x] Goal: Build comfort with Go syntax and learn how to structure a simple project with user interaction.

Phase 2: Understanding Go-Specific Features (1-2 Weeks)
- [x] Concurrency with Goroutines
    - [x] Learn Go’s concurrency model and how goroutines work.
    - [x] Understand how to use channels to communicate between goroutines.
    - [x] Explore when and how to apply concurrency in practical situations.
- [x] Project Idea: Modify your calculator to handle multiple calculations concurrently (e.g., let the user input several calculations at once, and process them in parallel).
- [x] Error Handling
    - [x] Learn how Go deals with errors (explicit error handling via error type).
    - [x] Practice writing functions that return errors and handle them gracefully.
    - [x] Get used to Go’s philosophy of "Errors are values."
- [x] Project Idea: Update your calculator to handle edge cases and invalid inputs (division by zero, non-numeric values) by utilizing Go’s error handling system.

Phase 3: Working with Files and APIs (2-3 Weeks)
- [x] Working with Files
    - [x] Learn how to read from and write to files in Go using os and io packages.
    - [x] Practice opening, writing, and closing files.
- [x] Project Idea: Enhance your calculator to log all calculations to a text file. Implement the ability to read past calculations from the file.
- [x] Build a Simple REST API
    - [x] Learn how to create a basic REST API using Go’s net/http package or the Gin framework.
    - [x] Handle routes, query parameters, and JSON responses.
1. Project Idea: [go-crud](https://github.com/AndreDrummer/go-crud)

Phase 4: Integrating AWS (3-4 Weeks)
- [x] Learn AWS Basics
    - [x] Understand how AWS works, focusing on AWS S3 for file storage and DynamoDB for NoSQL data storage.
    - [ ] Learn to interact with these services using Go's AWS SDK.

Phase 5: Build a Full Backend (2-3 Months)
- [x] Build a Full API Backend
    - [ ] Create a more complete API using Go that handles users, authentication, and dynamic data.
    - [x] Implement full CRUD (Create, Read, Update, Delete) operations in Go.
    - [ ] Use DynamoDB or PostgreSQL to store structured data.
1. Goal: Create a small Go-based backend that could support a Flutter app.
- [x] Deploy the API on AWS
    - [x] Learn how to deploy Go applications on AWS using EC2, Elastic Beanstalk, or AWS Lambda for serverless deployment.
    - [ ] Integrate API Gateway for exposing your Go API.
    - [x] Configure AWS IAM for secure access control.
