# Problem Solving Platform 

**CodeForge: Crafting Solutions - Building Coders**

## Overview

CodeForge is a dynamic platform designed for coders to tackle coding challenges and foster collaboration. Built using microservices architecture, CodeForge provides a range of core services to facilitate learning and skill development.

## Features

- **User Management**: Register and authenticate users to access platform features.
- **Challenge Service**: Create and manage challenges for users to solve.
- **Submission Service**: Submit code solutions, evaluated against predefined test cases, and receive scores.
- **Discovery Service**: Enable dynamic discovery of microservice instances.
- **API Gateway**: Serve as a single entry point for client interactions with microservices.

## Architecture

![Overview](https://github.com/ManujaDewmina/Distributed-Systems-Mini-Project/assets/92631934/21872aa2-1a78-49c0-b58e-d409380c1a73)

## Usage
- **Register and log in** to the platform.
- **Create new challenges** using the Challenge Service.
- **Submit solutions** and get immediate feedback through the Solution Service.
- **Explore and participate** in the coding community.

## Functionalities

### User Management

- **Register Users**: New users can register by providing their details.
- **Authenticate Users**: Handle user login and authentication.
- **Manage User Profiles**: Users can update their profile information and continue coding.

### Challenge Management

- **Create Challenges**: Allow users to create new challenges.
- **View Challenges**: Provide functionality to view available challenges.
  - View all challenges
  - View challenges by difficulty
  - Owned challenges
- **Update Challenges**: Enable modification of existing challenges by the creator.

### Submission Handling

- **Submit Solutions**: Users can submit their solutions to challenges.
- **View Submissions**: Provide a way to view submissions for a particular challenge.
- **Evaluate Submissions**: Facilitate automatic evaluation of submitted solutions.
- **Scoreboard**: Display scoreboard for every challenge and user separately.

## Getting Started

To get started with CodeForge, follow these steps:

1. Clone the project repository
2. Update GO mod and Go sum files
3. Docker compose up backend

## Tech Stack

![Code Forge - Mid Evaluation](https://github.com/ManujaDewmina/Distributed-Systems-Mini-Project/assets/92631934/12415e26-078b-4750-a22d-661123b4fbfb)

## User service component

![Userservice](https://github.com/ManujaDewmina/Distributed-Systems-Mini-Project/assets/92631934/d01b34d3-a4b1-4701-9a75-29f4ff5c4be0)

## Challenge service component

![Challenge Service](https://github.com/ManujaDewmina/Distributed-Systems-Mini-Project/assets/92631934/e2f49d13-c6e9-4e4c-8a2f-795222817596)

## Submission service component

![Submission Service](https://github.com/ManujaDewmina/Distributed-Systems-Mini-Project/assets/92631934/a6c2c251-d053-4162-b696-988b79c1fbc1)
