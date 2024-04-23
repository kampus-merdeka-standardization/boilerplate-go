# Golang Boilerplate

Golang Boilerplate for **T** Project based on **T** Standarization

## Development Requirement

To Develop the application make sure these app are installed on your machine

1. **Docker**: Install docker to simplify development.

2. **Make**: Makes are used to simplify command typed in the command line, list of command can be found in Makefile file.

3. **Golang Migrate CLI**: used for database migration, refer to this link for installation instructions : https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Setting Up Environment Configuration

To configure your environment variables for this project, please follow these steps:

1. **Create `.env` File**: Inside the `configs/env` directory, create a new file named `.env`. This file will hold your environment-specific configuration variables.

2. **Refer to `example.env`**: The contents of your `.env` file should be based on the template provided in `configs/env/*.example.env`. This file contains example configurations and serves as a guide for setting up your environment variables.

3. **Copy Configuration Settings**: Copy the contents of `example.env` into your newly created `.env` file. Ensure that you only modify the values in the `.env` file and not in `example.env`.

4. **Customize Configuration Values**: Modify the values in the `.env` file according to your specific environment requirements. Ensure that sensitive information such as API keys, database credentials, etc., are kept secure.

By following these steps, you'll be able to set up your environment configuration efficiently and ensure that your application functions correctly across different environments. If you have any questions or need further assistance, feel free to reach out. Happy coding!