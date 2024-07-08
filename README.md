# Web Application Scaffolding with Echo Framework

This repository contains a boilerplate for a basic MVC style web application using the Echo Framework. The application utilizes native Go templates for rendering views, ensuring efficient server-side HTML generation. It also includes an example model, controller, and view for a user registration and authentication system with JSON Web Tokens for authentication.

![Screenshot1](https://github.com/buelbuel/gowc/blob/main/public/images/screenshot.png?raw=true)

## Features

- **Web Components**: Utilizes native Web Components for encapsulated and reusable UI elements.
- **Vanilla JavaScript Controllers**: Implements JavaScript controllers for handling user interactions and extending functionality.
- **Air**: A lightweight development server that automatically reloads the application when files are modified.
- **Echo**: A lightweight web framework for Go that provides a fast and efficient way to build web applications.
- **Templates**: Utilizes Go's native templating system to generate HTML content.
- **JSON Web Tokens**: Utilizes JSON Web Tokens for authentication and authorization.
- **PostgreSQL**: Easily replace with any other database of your choice.
- **Let's Encrypt**: Automatically manages SSL certificates for HTTPS connections.

## Getting Started

1. Clone the repository:
    ```bash
    git clone https://github.com/buelbuel/gowc.git
    cd gowc
    ```
2. Install dependencies:
    ```bash
    go mod download
    go install github.com/air-verse/air@latest
    ```
3. Run the application:
    ```bash
    air
    ```
4. Access the application at http://localhost:4000

## Configuration

Below are the available configuration options in config.toml:

- **ServerAddress**: The address the server listens on.
- **StaticPaths**: A map of routes to static file directories.
- **UseLogger**: Whether to use the logger middleware.
- **UseTLS**: Whether to use TLS with provided certificate and key files.
- **UseAutoTLS**: Whether to use Let's Encrypt for automatic TLS.
- **CertFile**: Path to the TLS certificate file.
- **KeyFile**: Path to the TLS key file.
- **Domain**: Domain name for Let's Encrypt.
- **CacheDir**: Directory to cache Let's Encrypt certificates.
- **EnableCORS**: Whether to enable CORS.
- **CORSAllowOrigins**: List of allowed origins for CORS.
- **CORSAllowMethods**: List of allowed methods for CORS.
- **JWTSecret**: Secret for signing and verifying JSON Web Tokens.
- **DatabaseURL**: URL for the database.
- **DatabaseMaxConns**: Maximum number of database connections.
- **DatabaseMaxIdleConns**: Maximum number of idle database connections.
- **DatabaseConnMaxLifetime**: Maximum lifetime of a database connection.

### Notes

* CORS is disabled by default. You may need CORS for API endpoints or for local development, depending on your use case. To enable CORS, set `EnableCors`to true.
* The application uses PostgreSQL as the database. You can easily replace the database with any other database of your choice.
* To enable AutoTLS with Let's Encrypt, enable it in config.toml and set the domain as well as cache directory.
* To enable TLS without AutoTLS, set `UseTLS` to `true` and provide paths to the certificate and key files. You can generate a self-signed certificate using the following command:

```bash
go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```

## Authentication

The application uses JWT (JSON Web Tokens) for authentication. Below are the steps to set up and use authentication in the application.

### JWT Configuration

* The JWT configuration is managed through the `JwtConfig` struct in `internal/config/jwt_config.go`.
* The `RequireAuth` middleware in `internal/layers/auth_layer.go` is used to protect routes that require authentication. It uses the JWT configuration to validate tokens.
* In `internal/routes/web_routes.go`, the JWT middleware is applied to the routes that need protection. Public routes remain accessible without a token.
* The `AuthHandlers` struct in `internal/handlers/auth_handlers.go` contains the handlers for login and logout. The `LoginHandler` generates a JWT token upon successful login.

## View Rendering

The application uses Go's native templating system, which is both powerful and flexible. This allows for dynamic HTML content generation based on server-side logic and data. The templates are defined in the `views` directory and are rendered using Echo's built-in renderer.

### Template Structure

The templates are organized into different directories based on their purpose:

- **components**: Reusable components
- **layouts**: Layout templates
- **pages**: Page-specific templates
- **partials**: Reusable template components

### JavaScript Controllers

Vanilla JavaScript controllers are implemented to handle user interactions and extend functionality. These controllers are defined in the `public/controllers` directory.

### CSS Preprocessors

Since modern CSS allows for most needed functionality like nesting, variables and more, a preprocessor is not needed in most cases. However, If you want to ensure backwards compatibility and need more extensive features, you can include any preprocessor like SASS or LESS in the air build step. See the comments in air.toml.

### Directory Structure

The default structure of a new project is as follows:

```bash
├── cmd
│   └── main.go
├── config.toml
├── internal
│   ├── config
│   │   ├── app_config.go
│   │   ├── db_config.go
│   │   ├── jwt_config.go
│   │   ├── state_config.go
│   │   └── tls_config.go
│   ├── handlers
│   │   ├── api_handlers.go
│   │   ├── auth_handlers.go
│   │   └── web_handlers.go
│   ├── layers
│   │   └── auth_layer.go
│   ├── migrations
│   │   └── ...
│   ├── models
│   │   └── user.go
│   ├── routes
│   │   ├── api_routes.go
│   │   └── web_routes.go
│   └── utils
│       ├── context_util.go
│       ├── render_util.go
│       └── state_util.go
├── public
│   ├── css
│   │   └── ...
│   ├── images
│   │   └── ...
│   └── js
│       ├── components
│       │   └── ...
│       ├── controllers
│       │   └── ...
│       └── main.js
├── README.md
└── views
    ├── Base.html
    ├── layouts
    │   └── ...
    └── pages
        ├── app
        │   └── ...
        └── ...
```

## Contribution

Contributions to this project are welcome. Please refer to the contributing guidelines before making a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
