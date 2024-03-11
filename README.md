# PortKiller

PortKiller is a command-line application designed to kill a process running on a specified port. It provides a convenient way to free up a port that may be occupied by an application.

## Features

- Kill a process running on a specified port.
- Supports both TCP and UDP protocols.
- Cross-platform compatibility (Windows, macOS, Linux).

## Installation

PortKiller can be installed via pip:

```bash
pip install portkiller
```

## Usage

To use PortKiller, simply execute it from the command line with the desired port number:

```bash
portkiller <port_number>
```

Replace `<port_number>` with the port you want to free up.

### Example

To kill the process running on port 3000:

```bash
portkiller 3000
```

## Contributing

Contributions are welcome! If you'd like to contribute to PortKiller, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/improvement`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature/improvement`).
6. Create a new Pull Request.

## License

PortKiller is licensed under the MIT License. See [LICENSE](LICENSE) for more information.

## Disclaimer

This application is provided as-is without any guarantees. Use it at your own risk.
