# Project Title

Golang Boilerplate for **T** Project based on **T** Standarization

## Setting Up Environment Configuration

To configure your environment variables for this project, please follow these steps:

1. **Create `api.env`, `grpc.env` and `graphql.env` Files**: Inside the `configs/env` directory, create a new file named `.env`. This file will hold your environment-specific configuration variables.

2. **Refer to `example.api.env`, `grpc.example.env` and `graphql.example.env` respectively**: The contents of your `.env` file should be based on the template provided in `configs/env/*.example.env`. This file contains example configurations and serves as a guide for setting up your environment variables.

3. **Configure Environment Variables**: Open eachs `*.env` file you created in the previous step and populate it with the necessary environment variables for your project. Replace the placeholder values with your actual configuration settings.

4. **Copy Configuration Settings**: Copy the contents of `example.env` into your newly created `.env` file. Ensure that you only modify the values in the `.env` file and not in `example.env`.

5. **Customize Configuration Values**: Modify the values in the `.env` file according to your specific environment requirements. Ensure that sensitive information such as API keys, database credentials, etc., are kept secure.

By following these steps, you'll be able to set up your environment configuration efficiently and ensure that your application functions correctly across different environments. If you have any questions or need further assistance, feel free to reach out. Happy coding!
