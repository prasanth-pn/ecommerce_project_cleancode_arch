# MObdex - Ecommerce API for Selling Mobile and Accessories

MObdex is an ecommerce API built in Golang using a clean architecture. It provides functionality for selling mobile phones and accessories. The project incorporates various features such as cart management, user management, order management, product management, payment integration using Razor Pay, JWT authentication, SMTP email authentication, and Docker.

## Features

- **Cart Management**: Users can add products to their cart, view the cart, and update or remove items from the cart.

- **User Management**: The API allows for user registration, login, and authentication using JWT tokens. Users can also update their profile information.

- **Order Management**: Users can place orders for products, view their order history, and track the status of their orders.

- **Product Management**: The API provides endpoints for managing products, including adding new products, updating product details, and deleting products.

- **Payment Integration**: MObdex integrates with Razor Pay for secure payment processing. Users can make payments for their orders using various payment methods supported by Razor Pay.

- **JWT Authentication**: The API uses JWT tokens for user authentication and authorization, ensuring secure access to protected endpoints.

- **SMTP Email Authentication**: MObdex supports SMTP email authentication for user registration and password reset functionality, ensuring reliable email communication with users.

- **Docker and Docker Compose**: The project is containerized using Docker, allowing for easy deployment and scalability. Docker Compose is used to define and manage the multi-container application.

## Swagger Documentation

The API's Swagger documentation is available at the following link: [Swagger Documentation](https://prasanthpn.online/docs/index.html)

## Dependencies

MObdex relies on the following dependencies:

- **Viper**: Viper is used for configuration management, allowing easy handling of environment-specific settings.

- **Wire**: Wire is used for dependency injection, ensuring modular and testable code.

## Getting Started

To get started with MObdex, follow these steps:

1. Clone the repository: `git clone https://github.com/prasanth-pn/ecommerce_project_cleancode_arch.git`

2. Install the required dependencies: `make deps`

3. Configure the application by updating the necessary settings in the configuration file.

4. Build and run the application: `make run`

5. Access the API endpoints using a tool like cURL or Postman.

## Docker Deployment

To deploy MObdex using Docker, follow these steps:

1. Ensure Docker is installed and running on your system.

2. Docker multistage build is created for ease of use

3. Run the Docker image: `docker compose up -d `

## Contributing

Contributions to MObdex are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on the project's GitHub repository.

## License

MObdex is released under the [MIT License](https://opensource.org/licenses/MIT). Please see the `LICENSE` file for more details.

## Contact

For any inquiries or questions, please contact the project maintainer at [prasanthpn68@gmail.com](mailto:prasanthpn68@gmail.com).
