// Package handlers provides HTTP request handlers for the application.
//
// It includes handlers for user operations such as creating, retrieving,
// updating, and deleting users. It also contains handlers for authentication
// operations like registration, login, and logout.
//
// The package is structured around three main types:
//   - [UserHandlers]: Handles CRUD operations for users.
//   - [AuthHandlers]: Handles authentication-related operations.
//   - Web Handlers: Handles web page rendering. See [AuthPageHandler], [LoginFormHandler], etc.
//
// Each handler type is designed to work with the [echo.Echo] web framework and
// follows RESTful principles in its API design.
//
// Key handlers include:
//   - [UserHandlers.CreateUser]: Creates a new user.
//   - [AuthHandlers.LoginHandler]: Authenticates a user.
//   - [AuthHandlers.RegisterHandler]: Registers a new user.
//   - [AuthPageHandler]: Renders the authentication page.
//   - [DashboardPageHandler]: Renders the dashboard page.
//
// For more information on specific handlers, refer to the respective files:
//   - user_handlers.go: Contains all user-related handlers.
//   - auth_handlers.go: Contains all authentication-related handlers.
//   - web_handlers.go: Contains all web page rendering handlers.
package handlers
