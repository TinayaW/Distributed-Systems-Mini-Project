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

The decision to split the application into multiple microservices is driven by several key factors as follows.

**Modularity -** Due to the division of the application into several services with defined roles, modularity is attained. This clean division of responsibilities enables anyone working on either part of the application to solve, build, or understand either part without worrying about the other.

**Scalability -** Microservices can be scaled individually of one another. The architecture allows for adding more capacity to a specific service if the load, for example, in the Submission Service, is very high while the other services will be unaffected.

**Fault Isolation -** In monolithic architecture, a problem affecting one part of the application has the potential to result in the collapse of the entire system. Thus, by relying on microservices, we effectively limit failures of services, which increases the general reliability of the system.

**Technology Diversity -** It is possible that one service may use the best of one set of technologies while the other service may use the best of another set of technologies. It assists the teams to be flexible and optimize each service without the interference of the others.

**Continuous Deployment -** Microservices make it easier to deploy more often and more rapidly. Thus, since all services are independent, incorporating changes and bug-fixes can be done for one service only, without involving all other services of the application.

### 2.2.1 Contributions of Each Service to Overall Functionality

1. User Service

Oversees functions concerning the user, including creating, searching, modifying, or destroying the user. Serves as the authority for users' actual data and guarantees the proper storing and usage of the user data to be provided for other services' usage. It exposes contracts for basic operations on user data, namely CRUD, which is always needed to manage users in an application.

2. Challenge Service

Responsible for all the operations concerning challenges, including creating new ones, getting details of the challenges, editing, and even deleting the challenges. Thus, the application developed based on managing challenges is in the focus of the described work’s key functionality. It comprehensively secretes challenge data, thereby enabling the user to engage in the particular challenge as well as submitting the solutions. It also ensures that issues of challenge are well instantiated and ready for user exploits.

3. Submission Service

Handles the submissions for the challenges and it has functions to upload new submissions and to search the submissions by the user or by the challenge. Additionally, it has functions to view the submission details. Delivers all the features that are needed by users for submitting their solutions to the challenges as well as by administrators for viewing and moderating these submissions. This service is crucial to the implementation of the application’s objective of creating an environment where users can participate and respond to challenges.

### 2.2.2 Detailed Explanation of Design Decisions

1. User Service Design
- **Reasoning -** User management is a core process that cuts across the chains of the application. To separate them out into their own service means that the user-related operations are compartmentalized and can be further scaled if necessary.
- **Impact -** Provides higher security and better administration for user data, ordering all the data in a singular, specific service for the application and providing uniformity for the user data within the program.

2. Challenge Service Design
- **Reasoning -** Challenges are the main asset of the users as they involve real-life contents. This way, challenge management can be developed as its own unique service that we can fine-tune and scale the challenge-related operations separately.
- **Impact -** Enhances how challenge-related procedures are managed with specific regard made to challenge data as well as challenge logic, making challenges easier to maintain as well as scale.

3. Submission Service Design
- **Reasoning -** Submissions are the outcomes of challenges, which is a crucial aspect of the user’s engagement with the application. Processing submissions as a different service helps to organize data storing, processing, and submissions separately.
- **Impact -** Submission handling gets improved, offering users reliability when it comes to submission as well as the retrieval of challenge responses.

### 2.2.3 Overall Impact on Functionality

Each service contributes to the overall functionality by focusing on a specific aspect of the application.

- **User Service -** Provides maximum security checks in the handling of users’ information, which forms the basis of all relations in the application.
- **Challenge Service -** Handles the hosting of the game’s content and operations that allow the creation and management of challenges that the users can solve.
- **Submission Service -** Enables user contribution as the application allows users to post their solutions to the challenges posed by an organization, which is the main engagement aspect of the application.

By employing three distinct services, we achieve a well-organized, scalable, and maintainable architecture which supports the application's goals while ensuring a smooth user experience.

# 3. Microservices

## 3.1 Services

- **User Management** - Register and authenticate users to access platform features.
- **Challenge Service** - Create and manage challenges for users to solve.
- **Submission Service** - Submit code solutions, evaluated against predefined test cases, and receive scores.
- **Discovery Service** - Enable dynamic discovery of microservice instances.
- **API Gateway** - Serve as a single entry point for client interactions with microservices.

## 3.2 Implementation Methods

### Netflix Software Stack

While this project decided to use Go as the backend language and Consul as the service discovery tool, there are several of the Netflix software stack’s ideas and instruments that defined the architecture and the approach to implementation. The Netflix stack provides tools that help in creating fault-tolerant, highly available, and easy-to-manage microservices such as service discovery, load balancing, and routing done at the gateway level.

### Implementation with Go and Consul

#### Service Discovery with Consul

Consul is used for the discovery of services and the checking of the health of the services. Every service in the model runs a registration with pull data to Consul; data include service name, address, port, and the health checks’ port. Consul monitors these services and enables other services to— discovering the health state and communicating with the healthy instances.

- **Service Registration** - What happens is that when some service begins, it intends to join a particular Consul server. This registration includes relevant information that any service provider should provide such as the name of the service, and the place where the service can be accessed from.

- **Health Checks** - Usually, Consul takes a special check up on the registered services to know if they are healthy and running okay. If a service does not pass a health check, the service is then rejected from the pool of running services.

#### API Gateway with Go and Consul

The API Gateway acts as a reverse proxy, that is as the entry point into the application since it reverses the client requests and forwards them to the corresponding microservices. It communicates with the consul to get real time information of the available service instance to forward call accordingly.

- **Dynamic Routing** -The gateway utilizes the information from the Consul to actively write client requests to working instances of a service. This means that requests are always made to healthy services only.
- **Load Balancing** -  Load balancing is achieved through the fact that the gateway will forward requests coming in for a certain service to multiple instances of the service for efficiency.

### Core Services in Go

All the services (User Service, Challenge Service, Submission Service) are written in Go and adhere to REST conventions. Each service defines multiple APIs and communicates with other services using HTTP calls.

### Summary
As to the implementation, it has been done with Go and Consul, which replicates the Netflix software stack’s principles, including service discovery, dynamic routing, and load balancing. Such an approach guarantees the anti-frail, scalable and maintainable structure of microservices.

## 3.3 Core Services

### 3.3.1 Functionality

#### User Service 

##### Functionality
This service is organized to give overall charge of all operations affecting a user account or information whether it entails Creating New User Account, Retrieval, Updation or Deletion of User Account. It becomes the place where all the user related data is held and managed within the context of the application.

##### REST API Endpoints

###### Create User
- **URL:** `http://localhost:8083/user/create`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Description:** Creates a new user in the system.

###### Get Users
- **URL:** `http://localhost:8083/user/users`
- **Method:** `GET`
- **Description:** Retrieves a list of all users.

###### Get User by ID
- **URL:** `http://localhost:8083/user/{userid}`
- **Method:** `GET`
- **Description:** Retrieves detailed information for a specific user by ID.

###### Get User by name
- **URL:** `http://localhost:8083/user/usernamae{usernamae}`
- **Method:** `GET`
- **Description:** Retrieves detailed information for a specific user by usernamae.

###### Update User
- **URL:** `http://localhost:8083/user/update/{userid}`
- **Method:** `PUT`
- **Content-Type:** `application/json`
- **Description:** Updates the information of a specific user by ID.

###### Delete User
- **URL:** `http://localhost:8083/user/delete/{userid}`
- **Method:** `DELETE`
- **Description:**  Deletes a specific user by ID.


#### Challenge Service 

##### Functionality
The Challenge Service manages all operations related to challenges, including creation, retrieval, updating, and deletion of challenges. It is central to the application’s core functionality, facilitating the management and participation in challenges.

##### REST API Endpoints

   ###### Create Challenge
   - **URL:** `http://localhost:8083/challenge/create`
   - **Method:** `POST`
   - **Description:** Creates a new challenge in the system.

   ###### Get Challenges
   - **URL:** `http://localhost:8083/challenge/challenges`
   - **Method:** `GET`
   - **Description:** Retrieves a list of all challenges.

   ###### Get Challenge by ID
   - **URL:** `http://localhost:8083/challenge/{challengeid}`
   - **Method:** `GET`
   - **Description:** Retrieves detailed information for a specific challenge by ID.

   ###### Get Challenge by Author ID
   - **URL:** `http://localhost:8083/challenge/challenges/user{userid}`
   - **Method:** `GET`
   - **Description:** Retrieves detailed information for a specific challenge by Author ID.

   ###### Get Challenge by Difficulty
   - **URL:** `http://localhost:8083/challenge/difficulty{difficulty}`
   - **Method:** `GET`
   - **Description:** Retrieves detailed information for a specific challenge by Difficulty.

   ###### Update Challenge
   - **URL:** `http://localhost:8083/challenge/update/{challengeid}`
   - **Method:** `PUT`
   - **Description:** Updates the information of a specific challenge by ID.

   ###### Delete Challenge
   - **URL:** `http://localhost:8083/challenge/delete/{challengeid}`
   - **Method:** `DELETE`
   - **Description:** Deletes a specific challenge by ID.


#### Submission Service 

##### Functionality
The Submission Service manages all operations related to submissions, including uploading new submissions, retrieving submissions, and viewing submission details. It handles user responses to challenges.

##### REST API Endpoints

   ###### Upload Submission
   - **URL:** `http://localhost:8083/submission/upload`
   - **Method:** `POST`
   - **Description:** Uploads a new submission for a challenge.

   ###### Get Submissions by User ID
   - **URL:** `http://localhost:8083/submission/user/{userId}`
   - **Method:** `GET`
   - **Description:** Retrieves all submissions made by a specific user.

   ###### Get Submissions by Challenge ID
   - **URL:** `http://localhost:8083/submission/challenge/{challengeId}`
   - **Method:** `GET`
   - **Description:** Retrieves all submissions for a specific challenge.

   ###### Get Submission by ID
   - **URL:** `http://localhost:8083/submission/{submissionid}`
   - **Method:** `GET`
   - **Description:** Retrieves detailed information for a specific submission by ID.


### 3.3.2 REST endpoints

![Code Forge - Mid Evaluation](https://github.com/user-attachments/assets/505cd53f-7095-4a56-a78d-fd92c42dac36)

![14](https://github.com/user-attachments/assets/8ef3babc-711b-45a4-aaef-973485ed02bd)

![15](https://github.com/user-attachments/assets/d1c52867-3282-4cbb-a429-b5a048a7e3c4)

### 3.3.3 Inter service Interaction

The Challenge Service communicates the Submission Service to request and perform operations on submissions concerning certain challenges. That way, the Submission Service communicates communicates with the Challenge Service to obtain data related to the challenges for every submission. This helps in guaranteeing that the submissions are well associated with the appropriate users and challenges.

![Interconnection](https://github.com/user-attachments/assets/c98eba21-e405-4934-815a-cb0ca1de72b7)

## 3.4 Discovery Server - Consul

### Service Registration and Monitoring with Consul

#### Service Registration

Most of the services have the ability to register themselves with the Consul server when the service starts. During the registration process, the metadata like the service name, the address, the port number, and health check endpoints, are required. This information makes it possible for Consul to have a current register of all the available services.

#### Monitoring

Consul runs health checks periodically for the registered services. Every service has a designated endpoint for health check, such as `/health`, which Consul uses to check for the health of the service in question. The health check can fail when Service B is unavailable and unhealthy, then Consul will exclude it from the list of available services, besides services B will not be discovered by other services and the API Gateway since they are unhealthy and it will Automatically re-adds previously unhealthy services to the pool of available services once they become healthy again

- Service Registration: Microservices register themselves with discovery service, providing necessary metadata and health check information.

- Service Discovery: Services query discovery service to discover other services they need to communicate with, ensuring dynamic and up-to-date connections.

- Service Health Checks: Regular health checks ensure that only healthy service instances are available for handling requests.

- Automatic Recovery: Automatically re-adds previously unhealthy services to the pool of available services once they become healthy again.

## 3.5 API Gateway

The API Gateway dynamically routes client requests to the appropriate service based on URL paths and metadata retrieved from Consul. It employs reverse proxying to properly direct the requests to the concerned service instances while making sure that each request is handled by the right service. API Gateway takes incoming requests and divides them among multiple instances of a given service to increase the efficiency of the system. Consul is used to discover the services available in the system and this integration is done with the API Gateway. It gets the Consul to fetch the addresses of the healthy instances of the services and this informs how the requests are forwarded. This makes certain that the API Gateway always forwards requests to services that are active and responsive in making sure the total system availability.

- Single External Endpoint: Acts as the sole entry point for all external client requests, ensuring that microservices endpoints remain internal and secure.

- Routing Requests: Directs incoming client requests to the appropriate microservice based on the URL and other request parameters.

- Service Discovery Integration: Integrates with a service registry to dynamically discover and keep track of available service instances.

- Logging and Monitoring: Provides centralized logging and monitoring of all incoming and outgoing traffic for better observability.

- Request and Response Transformation: Modifies incoming requests and outgoing responses as needed.

# 4 User Interface

## 4.1 Implementation Details

React-Typescript(Discuss)

## 4.2 API Testing Tools

Postman is used here and it is a tool that is widely used and utilized in API testing that makes the implementation of API testing rather easy and manageable. It is a convenient tool to send the HTTP requests and study the response as it is widely used by developers and testers. 

For this particular project, there was the development of a Postman collection to contain the different API requests of the services under development. The collection consisted of such requests – User Service, Challenge Service, and Submission Service. Every request in the collection was created to target the certain endpoints and the possibilities that are connected with them.

### Testing Workflow

1. Setting Up Requests - All endpoints of the services were tested by creating a request in postman corresponding to its end points. For instance, the User Service contained the requests to create, get, modify, or remove users. The requests were assigned with correct HTTP verbs (GET, POST, PUT, DELETE) and headers and mostly the Content-Type header that defines the type of data that is being transferred.
2. Organizing Tests - The requests were grouped in to folders bearing the nature of the services they contained so as to be easily managed and methodologically ran through tests. This organization assisted in categorizing similar requests and executing them in a serial manner to evaluate different automations.
3. Running Collections - Collection runner in postman permitted the running of the whole collection, as well as individual folders within the collection. This feature allows for each endpoint to be tested as a group and it was also autographed to guarantee that changing one section of the API would not impact the other sections. The collection runner also supplied me with comprehensive reports of the test runs of the failed requests and the cause of the failures which we had to address in order to fix the API.

### Benefits of Using Postman
   - Reduced the amount of time it takes to perform testing by making the sending of requests and analysis of responses easier.
   - The collection could be made available to all members of a particular team so that the way various tests are conducted remains standardized, and people can work together.

In general, Postman was an essential help in the phases of API development, because using these tools, it is possible to make sure that the APIs are stable, effective and created with desired functionality before all APIs would be put into production.

# 5 Deployment

Pre-request instrallations : Docker(v24.0.6), Go(V1.22.0), Node(v18.18.2), PostgreSQl
Environment : Linux

## 5.1 Frontend
1. Go to the project directory and run `npm install` to install the necessary dependencies.
2. Run `npm start` to start the application.

## 5.2 Backend

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

# 6 Source Code

## 6.1 Github Link

https://github.com/TinayaW/Distributed-Systems-Mini-Project

## 6.2 Development Challenges

(Discuss)

# 7 References

1. Go documentation - https://go.dev/doc/
2. Docker documentation - https://docs.docker.com/
3. consul documentation - https://developer.hashicorp.com/consul/docs
4. React documentation - https://react.dev/learn   ,   https://react.dev/blog/2023/03/16/introducing-react-dev

