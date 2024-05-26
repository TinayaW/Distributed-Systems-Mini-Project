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

<div align='center'>
    <img src="https://github.com/ManujaDewmina/Problem-Solving-Platform/assets/92631934/a454961b-6801-4762-b9b5-71da617534f6" width="700" align="center">
</div>

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
