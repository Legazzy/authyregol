# Authyre API

A service for authenticating users

## Description

The Authyre API is a REST service for providing authentication and generic authorization for users.
It is capable of managing user accounts, permissions and tokens with fine-grained control over
access to resources and permissions. The goal of the project is to provide a standalone service
that can be accessed by other services as a single point of authentication.

## Installation

### General

The whole application is written in the go language and can be directly compiled from source,
without any further configuration. The general build command would look like this:

```$ go build -o authyreapi cmd/```

This will create a binary called `authyreapi` in the current directory. This binary can be
executed directly. The application can also be compiled to different platforms by specifying
the `GOOS` and `GOARCH` environment variables.

### Persistence

The application itself does not directly store anything on the disk. Even the configuration
is loaded from the environment. However, the application requires a MongoDB database to store
the user data. To establish a connection to the database, the application needs to be
configured with the right connection details.

### Docker

The application can also be built as a docker image. A matching [Dockerfile](run/Dockerfile) is
provided in combination with a sample [docker-compose](run/docker-compose.yml) file. The
docker-compose file also bundles the required MongoDB database.

## Configuration

The configuration is done by using environment variables and kept really slim. All the content of the
application gets populated into the Database on the first run with standard values. Because of the nature of
the API, those values can be modified on runtime to the own needs.

### Environment Variables

| Variable                          | Description                                                 | Default   |
|-----------------------------------|-------------------------------------------------------------|-----------|
| API_HOSTNAME                      | The hostname the API is listening on                        | localhost |
| API_INETPORT                      | The port the API is listening on                            | 8080      |
| MONGO_HOSTNAME                    | The hostname of the MongoDB server                          | localhost |
| MONGO_INETPORT                    | The port of the MongoDB server                              | 27017     |
| MONGO_DATABASE                    | The name of the MongoDB database                            | authyre   |
| MONGO_USERNAME                    | The username for the MongoDB database                       | authyre   |
| MONGO_PASSWORD                    | The password for the MongoDB database                       | authyre   |
| POPULATION_REPOPULATE_PERMISSIONS | Whether to repopulate the permissions collection on startup | true      |
| POPULATION_REPOPULATE_SERVICES    | Whether to repopulate the services collection on startup    | true      |
| POPULATION_REPOPULATE_USERS       | Whether to repopulate the users collection on startup       | true      |

### Defaults

The application comes with a set of default values for the permissions, services and users.
The following table lists the default values:

#### Permissions

Authyre uses a permission system to control access to resources. The permissions are stored
in the database and can be modified at runtime. Each one consists at minimum of a keyword,
that the application itself and also external services can use to identify the permission,
a reference to a user and one to the service that the permission belongs to.

The following table lists the default permissions for the `authyre` service:

| Keyword              | Description                                            |
|----------------------|--------------------------------------------------------|
| personal_changes     | The permission to change personal data                 |
| personal_infos       | The permission to view personal data                   |
| personal_permissions | The permission to view and manage personal permissions |
| personal_tokens      | The permission to view and manage personal tokens      |
| service_permissions  | The permission to manage services and permissions      |
| service_services     | The permission to manage services                      |
| service_users        | The permission to manage all users                     |

### Services

Services are used to group permissions and users to the actual services that use the API.
Because Authyre itself has to manage permissions in regard to users and services as well,
it is also just a service. At startup the application will create a service called `authyre`
as a reference to itself. The service has the Identifier `00000000-0000-0000-0000-000000000000`.

### Users

Users are the main entity of the application. They are used to authenticate against the API
and to manage permissions and tokens. The application will create a default
user called `authyre` with the password `authyre` on startup.
It has the default Identifier`00000000-0000-0000-0000-000000000000`.

## Routes

The API provides a set of routes to manage users, permissions, services and tokens. The
following table lists all routes and their purpose:

| Method | Route                                                | Description                                                          |
|--------|------------------------------------------------------|----------------------------------------------------------------------|
| POST   | /token                                               | Authenticate a user and create a new token with provided permissions |
| GET    | /token                                               | Get all tokens for the authenticated user                            |
| GET    | /token/:credential_access                            | Get a specific token for the authenticated user                      |
| DELETE | /token/:credential_access                            | Delete a specific token for the authenticated user                   |
| ---    | ---                                                  | ---                                                                  |
| POST   | /user                                                | Create a new user                                                    |
| GET    | /user                                                | Get all users                                                        |
| GET    | /user/:credential_username                           | Get a specific user                                                  |
| PATCH  | /user/:credential_username                           | Update a specific user                                               |
| DELETE | /user/:credential_username                           | Delete a specific user                                               |
| ---    | ---                                                  | ---                                                                  |
| POST   | /service                                             | Create a new service                                                 |
| GET    | /service                                             | Get all services                                                     |
| GET    | /service/:description_name                           | Get a specific service                                               |
| DELETE | /service/:description_name                           | Delete a specific service                                            |
| ---    | ---                                                  | ---                                                                  |
| POST   | /permission/:identifier_permission:/:identifier_user | Create a new permission                                              |
| GET    | /permission/:identifier_permission:/:identifier_user | Get all permissions                                                  |
| DELETE | /permission/:identifier_permission:/:identifier_user | Delete a permission                                                  |

Both request and response bodies are encoded in JSON. The individual request DTOs are
not provided in this documentation, but can be found in the [api/transfer](api/transfer) directory.

The API uses a custom authentication method. The authentication is done by providing a
Bearer token in the `Authorization` header of the request. The token can be obtained by
calling the `/token` route with a valid username and password provided trough Basic authentication.