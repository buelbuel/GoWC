# Web Application Scaffolding with Echo Framework
[![Go Reference](https://pkg.go.dev/badge/github.com/buelbuel/gowc.svg)](https://pkg.go.dev/github.com/buelbuel/gowc)
[![codecov](https://codecov.io/github/buelbuel/GoWC/graph/badge.svg?token=0JOT2V6JCO)](https://codecov.io/github/buelbuel/GoWC)
[![Go Report Card](https://goreportcard.com/badge/github.com/buelbuel/gowc)](https://goreportcard.com/report/github.com/buelbuel/gowc)
![GitHub License](https://img.shields.io/github/license/buelbuel/gowc)

This repository contains a boilerplate for a basic MVC style web application using the Echo Framework. The application utilizes native Go templates for rendering views, ensuring efficient server-side HTML generation. ~~It also includes an example model, controller, and view for a user registration and authentication system with JSON Web Tokens for authentication.~~

## Features

⚡️ **Web Components**: Utilizes native Web Components for encapsulated and reusable UI elements.  
💧 **Vanilla JavaScript Controllers**: Implements JavaScript controllers for handling user interactions and extending functionality.  
☁️ **Air**: A lightweight development server that automatically reloads the application when files are modified.  
🌐 **Echo**: A lightweight web framework for Go that provides a fast and efficient way to build web applications.  
📄 **Templates**: Utilizes Go's native templating system to generate HTML content.  
🔒 **Let's Encrypt**: Automatically manages SSL certificates for HTTPS connections.  

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
- **LogOutput**: Specifies where to output logs ("stdout", "stderr", or "file").
- **LogFile**: Path to the log file when LogOutput is set to "file".
- **ColorizeLogger**: Whether to colorize log output.
- **UseTLS**: Whether to use TLS with provided certificate and key files.
- **UseAutoTLS**: Whether to use Let's Encrypt for automatic TLS.
- **CertFile**: Path to the TLS certificate file.
- **KeyFile**: Path to the TLS key file.
- **Domain**: Domain name for Let's Encrypt.
- **CacheDir**: Directory to cache Let's Encrypt certificates.
- **EnableCORS**: Whether to enable CORS.
- **CORSAllowOrigins**: List of allowed origins for CORS.
- **CORSAllowMethods**: List of allowed methods for CORS.
- **RateLimit**: Rate limiting requests per second.
- **RateBurst**: Maximum burst for rate limiter.

### Notes

* CORS is disabled by default. You may need CORS for API endpoints or for local development, depending on your use case. To enable CORS, set `EnableCors`to true.
* The application uses PostgreSQL as the database. You can easily replace the database with any other database of your choice.
* To enable AutoTLS with Let's Encrypt, enable it in config.toml and set the domain as well as cache directory.
* To enable TLS without AutoTLS, set `UseTLS` to `true` and provide paths to the certificate and key files. You can generate a self-signed certificate using the following command:

```bash
go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```

## View Rendering

The application uses Go's native templating system, which is both powerful and flexible. This allows for dynamic HTML content generation based on server-side logic and data. The templates are defined in the `views` directory and are rendered using Echo's built-in renderer.

### JavaScript Controllers

Vanilla JavaScript controllers are implemented to handle user interactions and extend functionality. These controllers are defined in the `public/controllers` directory.

### CSS Preprocessors

Since modern CSS allows for most needed functionality like nesting, variables and more, a preprocessor is not needed in most cases. However, If you want to ensure backwards compatibility and need more extensive features, you can include any preprocessor like SASS or LESS in the air build step. See the comments in air.toml.

### Directory Structure

The default structure of a new project is as follows:

```bash
├── cmd
│   ├── doc.go
│   ├── main.go
│   └── migrate
│       └── main.go
├── config.toml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── app_config.go
│   ├── handlers
│   │   └── web_handlers.go
│   ├── routes
│   │   └── routes.go
│   └── utils
│       └── render_util.go
└── resources
    ├── css
    │   ├── main.css
    │   └── variables.css
    ├── images
    │   ├── favicon.ico
    │   └── screenshot.png
    ├── js
    │   ├── components
    │   │   ├── ButtonComponent.js
    │   │   └── ...
    │   ├── controllers
    │   │   └── ButtonController.js
    │   └── main.js
    └── views
        ├── Base.html
        ├── layouts
        │   └── FrontLayout.html
        └── pages
            └── Start.html
```

## Contribution

Contributions to this project are welcome! Please follow these guidelines:

1. Fork the repository and create a new branch for your contribution.
2. Make your changes and ensure that the code is properly formatted.
3. Write clear and concise commit messages.
4. Test your changes thoroughly.
5. Submit a pull request to the main repository.

Thank you for your contribution!

## License

This project is licensed under the MIT License - see the LICENSE file for details.