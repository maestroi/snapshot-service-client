# Go Application

This is a Go application that downloads and extracts a snapshot from a specified URL.

## Prerequisites

- Go installed
- Docker (optional)

## Getting Started

### Running Locally

1. Clone the repository:

   ```bash
   git clone https://github.com/maestroi/snapshot-service-client
   ```

2. Change to the project directory:

   ```bash
   cd snapshot-service-client
   ```

3. Build the application:

   ```bash
   make build
   ```

4. Run the application:

   ```bash
   make run
   ```

   This will run the application locally.

### Running with Docker

1. Build the Docker image:

   ```bash
   make docker-build
   ```

2. Run the Docker container:

   ```bash
   make docker-run
   ```

   This will run the application inside a Docker container.

3. The application will download and extract the snapshot files to the specified target directory on your host machine.

## Makefile Commands

- `make build`: Build the application.
- `make run`: Run the application locally.
- `make docker-build`: Build the Docker image.
- `make docker-run`: Run the Docker container.
- `make clean`: Clean up the generated files.
- `make help`: Show a list of available commands.


## Run binary
This will force download a latest snapshot for the nimiq non albatross testnet node in current directory.
```bash
snapshot-downloader -f nimiq-v1 testnet . 
``` 

## Options

- `-f`: Overwrite existing files without prompt.

## License

This project is licensed under the [MIT License](LICENSE).
