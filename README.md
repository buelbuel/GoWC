# Web Application Scaffolding with Echo Framework

This repository contains the scaffolding for a basic MVC style web application using the Echo Framework. The application utilizes native Go templates for rendering views, ensuring efficient server-side HTML generation.

## Features

- **Turbo**: Integrates Hotwire's Turbo to speed up page transitions without requiring full page reloads. This approach enhances the responsiveness and speed of the application.
- **Stimulus**: Uses Stimulus for modest JavaScript enhancements, allowing for extended functionality with minimal overhead.
- **Air**: A lightweight development server that automatically reloads the application when files are modified.
- **Echo**: A lightweight web framework for Go that provides a fast and efficient way to build web applications.
- **Templates**: Utilizes Go's native templating system to generate HTML content.

## Getting Started

1. Clone the repository:
    ```bash
    git clone https://github.com/buelbuel/gowired.git
    cd gowired
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

If you run this application in a container, you can use the devcontainer extension to open the project in a remote container. After the post create command has been executed, you can use the devcontainer extension to open the project in a remote container and simply run `air` to start the application.

## View Rendering

The application uses Go's native templating system, which is both powerful and flexible. This allows for dynamic HTML content generation based on server-side logic and data.
Since modern CSS allows for most needed functionality like nesting, variables and more, a preprocessor is not needed in most cases. However, If you want to ensure backwards compatibality and need more extensive features, you can include any preprocessor in the air build step. See the comments in air.toml.
```bash
gowired/
├── cmd/
│ └── main.go # Application entry point
├── internal/
│ ├── api/
│ │ ├── handlers/ # HTTP request handlers
│ │ └── routes/ # Route definitions
│ ├── config/ # Application configuration
│ ├── middleware/ # Custom middleware
│ └── utils/ # Utility functions
├── public/
│ ├── css/ # CSS files
│ └── js/ # JavaScript files
├── views/
│ ├── components/ # Reusable components
│ ├── layouts/ # Layout templates
│ ├── pages/ # Page-specific templates
│ └── partials/ # Reusable template components
└── go.mod # Go module file
```

## Contribution

Contributions to this project are welcome. Please refer to the contributing guidelines before making a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
