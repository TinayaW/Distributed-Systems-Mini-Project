# Problem Solving Platform 

**CodeForge: Crafting Solutions - Building Coders**

# 1. Introduction

## 1.1. Overview

CodeForge is a dynamic platform designed for coders to tackle coding challenges and collaboration with other members of the community. CodeForge, designed with a microservices architecture, provides an array of services that aims at enhancing the learning process, developing necessary skills, and encouraging active community presence. The idea of the platform is to help corders at various skill levels to code more effectively and in doing so, they are provided with real time responses and a rich set of features to track and improve their coding abilities.

- **Register and log in** to the platform.
- **Create new challenges** using the Challenge Service.
- **Submit solutions** and get immediate feedback through the Solution Service.
- **Explore and participate** in the coding community.

## 1.2. Use cases and Functionalities

1. User Management
- Register Users: Allows new users to register by providing their details.
- Authenticate Users: Handles user login and authentication.
- Manage User Profiles: Enables users to update their profile information and continue coding profile.

2. Challenge Management
- Create Challenges: Allows users to create new challenges.
- View Challenges: Provides functionality for users to view available challenges. (View all challenges, View challenges by difficulty, Owned challenges)
- Update Challenges: Enables modification of existing challenges by the creator.

3. Submission Handling
- Submit Solutions: Users can submit their solutions to challenges.
- View Submissions: Provides a way to view submissions for a particular
challenge.
- Evaluate Submissions: Facilitates the automatic evaluation of submitted solutions.
- Score board: Give score board for every challenge and every user seperately.

## 1.3. Overall goal

The overall goal of CodeForge is to build a friendly, vibrant, and instructional learning community where programmers can develop their abilities continuously. Thus, CodeForge’s goals are aligned with its users’ interests – to provide a vast number of coding challenges that include live feedback and collaboration, allowing users to enhance their problem-solving skills and learn new algorithms and techniques in order be more efficient programmers. The layout of the codeforge platform is built to encourage users to promotes knowledge sharing as well as competition in the challenges. Thus, being an all-in-one solution, CodeForge has the potential to become the ultimate guide and the popular platform for the majority of coders aspiring to develop their coding abilities and connect with like-minded people.

## 1.4. Tech Stack

![Code Forge - Mid Evaluation (1)](https://github.com/user-attachments/assets/71302e10-2ed7-41ba-b2d8-a54e05d05825)

# 2. Architecture

## 2.1. Architectural Diagram

### 1. Solution overview architecture

![Architecture](https://github.com/user-attachments/assets/10a548f5-12fa-41b8-a69a-153d8b215cc1)

### 2. User Service

![Userservice](https://github.com/user-attachments/assets/ed1ee0f4-0c7d-4669-b3e0-86141e4ac8f7)

### 3. Challenge Service

![Challenge Service](https://github.com/user-attachments/assets/60a6b585-87a7-486a-b460-31bb6969e49f)

### 4. Submission Service

![Submission Service](https://github.com/user-attachments/assets/3899d4e6-6c14-48a3-b132-4c61e7584110)

## 2.2. Design Decisions

(Discuss)

# 3. Microservices

## 3.1.Services

- **User Management**: Register and authenticate users to access platform features.
- **Challenge Service**: Create and manage challenges for users to solve.
- **Submission Service**: Submit code solutions, evaluated against predefined test cases, and receive scores.
- **Discovery Service**: Enable dynamic discovery of microservice instances.
- **API Gateway**: Serve as a single entry point for client interactions with microservices.

## 3.2. Implementation Methods

(Discuss)

## 3.3. Core Services

### 3.3.1. Functionality

(Discuss)

### 3.3.2. REST endpoints

![Code Forge - Mid Evaluation](https://github.com/user-attachments/assets/505cd53f-7095-4a56-a78d-fd92c42dac36)

![14](https://github.com/user-attachments/assets/8ef3babc-711b-45a4-aaef-973485ed02bd)

![15](https://github.com/user-attachments/assets/d1c52867-3282-4cbb-a429-b5a048a7e3c4)

### 3.3.3. Inter service Interaction

(Discuss)

![Interconnection](https://github.com/user-attachments/assets/c98eba21-e405-4934-815a-cb0ca1de72b7)

## 3.4. Discovery Server - Consul

(Discuss Config)

- Service Registration: Microservices register themselves with discovery service, providing necessary metadata and health check information.

- Service Discovery: Services query discovery service to discover other services they need to communicate with, ensuring dynamic and up-to-date connections.

- Service Health Checks: Regular health checks ensure that only healthy service instances are available for handling requests.

- Automatic Recovery: Automatically re-adds previously unhealthy services to the pool of available services once they become healthy again.

## 3.5. API Gateway

(Discuss Config)

- Single External Endpoint: Acts as the sole entry point for all external client requests, ensuring that microservices endpoints remain internal and secure.

- Routing Requests: Directs incoming client requests to the appropriate microservice based on the URL and other request parameters.

- Service Discovery Integration: Integrates with a service registry to dynamically discover and keep track of available service instances.

- Logging and Monitoring: Provides centralized logging and monitoring of all incoming and outgoing traffic for better observability.

- Request and Response Transformation: Modifies incoming requests and outgoing responses as needed.

# 4. User Interface

## 4.1.Implementation Details

React-Typescript(Discuss)

## 4.2. API Testing Tools

Postman(discuss)

# 5. Deployment

Pre-request instrallations : Docker(v24.0.6), Go(V1.22.0), Node(v18.18.2), PostgreSQl
Environment : Linux

## 5.1. Frontend
1. Go to the project directory and run `npm install` to install the necessary dependencies.
2. Run `npm start` to start the application.

## 5.2. Backend

1. Open Terminal in root directory

2. `cd api-gateway`
   `go mod init backend/api-gateway`
   `go mod tidy`
   `cd ..`

3. `cd challenge-service`
   `go mod init backend/challenge-service`
   `go mod tidy`
   `cd ..`

4. `cd submission-service`
   `go mod init backend/submission-service`
   `go mod tidy`
   `cd ..`

5. `cd user-service`
   `go mod init backend/user-service`
   `go mod tidy`
   `cd ..`

6. `docker-compose up --build`

# 6. Source Code

## 6.1. Github Link

https://github.com/TinayaW/Distributed-Systems-Mini-Project

## 6.2. Development Challenges

(Discuss)

# 7. References

1. Go documentation - https://go.dev/doc/
2. Docker documentation - https://docs.docker.com/
3. consul documentation - https://developer.hashicorp.com/consul/docs
4. React documentation - https://react.dev/learn   ,   https://react.dev/blog/2023/03/16/introducing-react-dev

