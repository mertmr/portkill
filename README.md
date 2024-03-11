# PortKill

PortKill is a command-line application designed to kill a process running on a specified port. It provides a convenient way to free up a port that may be occupied by an application.

## Features

- Kill a process running on a specified port.
- Supports both TCP and UDP protocols.
- Cross-platform compatibility (Windows, macOS, Linux).

## Installation

PortKill can be installed via releases:

https://github.com/mertmr/portkill/releases/

For Windows - with scoop

```
scoop bucket add scoop https://github.com/mertmr/scoop
scoop install scoop/portkill
```

## Usage

To use PortKill, simply execute it from the command line with the desired port number:

```bash
portkill <port_number>
```

Replace `<port_number>` with the port you want to free up.

### Example

To kill the process running on port 3000:

```bash
portkill 3000
```


## License

PortKill is licensed under the MIT License. See [LICENSE](LICENSE) for more information.

## Disclaimer

This application is provided as-is without any guarantees. Use it at your own risk.
