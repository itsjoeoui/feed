# feed

This is the server for [my feed](https://www.jyu.dev/feed).

## Architecture

I am adapting to the hexagonal architecture, which is also known as the Ports
and Adapters architecture.

The goal is to create a flexible and decoupled system. Separating the core
business logic from external dependencies and infrastructure details.

## Folder Structure

```bash
.
├── README.md
├── cmd
│   └── main.go
├── config
│   └── viper.go
├── go.mod
├── go.sum
├── internal
│   ├── database
│   │   ├── migrations
│   │   └── repository
│   ├── delivery
│   │   └── http
│   ├── domain
│   │   ├── entity
│   │   └── usecase
│   ├── mock
│   └── ports
│       ├── repository
│       └── usecase
```

- `cmd`: The main entry point of the application. It initializes the application
and wires up all the dependencies.

- `config`: Contains the configuration related code. In this case, Viper is used
to manage the reading and parsing of the configuration.

- `internal`: The core logic of the application and it is not exposed to the
outside world.

Below are all the subfolders in the `internal` folder:

- `database`: Contains all the database related codem such as migrations
for managing schema changes and repository implementations to actually interact
with the database.

- `delivery`: Everything related to delivering the application's functionality
to users or external systems. In our case, we have a HTTP delivery, which will
simply contains the API route handlers and we will wire things up with Chi.

- `domain`: The heart of the application's business logic.
- `domain/entity`: Holds the domain entities or models, representing the main
concept of this application. This is the object that we will be interacting with.
- `domain/usecase`: Contains the use case implementations. It encapsulates the
application-specific logic and define the operations that can be performed on
the domain entities.

- `mock`: Mock implementations for testing/simulation.

- `ports`: Defines the interfaces or contracts that interact with the external
systems or dependencies.
- `ports/repository`: Defines the interfaces that abstract the persistence layer
operations and interactions with the database.
- `ports/usecase`: Defines the interfaces representing the available use cases or
possible application operations, allowing the app to be decoupled from the
specific use case implementations.

## Credit

Thanks [LuigiAzevedo](https://github.com/LuigiAzevedo)'s
[Public-Library](https://github.com/LuigiAzevedo/Public-Library) repo
for the source of inspiration :)
