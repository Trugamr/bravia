# Bravia TV CLI

A command-line interface tool for controlling Sony Bravia TVs. Control your TV's power, volume, inputs, and apps directly from your terminal.

## Features

- Power control (on/off/status)
- Volume adjustment
- Input switching
- App management
- Secure PSK authentication

## Installation

Download the latest binary for your platform from the [releases page](https://github.com/trugamr/bravia/releases).

## Usage as a Library

The API package can be used independently in your Go projects:

```go
import "github.com/trugamr/bravia/api"

client := api.NewClient(baseURL).WithAuthPSK(psk)

// Control power
client.System.SetPowerStatus(true)

// Control volume
client.Audio.SetAudioVolume("25", "speaker")
```

## Configuration

Create a `config.yaml` file in one of the following locations:
- Current directory
- `$HOME/.config/bravia/`
- `/etc/bravia/`

```yaml
base-url: "http://your-tv-ip"
psk: "your-pre-shared-key"
```

Alternatively, you can provide these values via command-line flags:
```bash
bravia --base-url="http://your-tv-ip" --psk="your-pre-shared-key" [command]
```

## Usage

```bash
bravia [command]

Available Commands:
  apps        List and open apps on your TV
  inputs      List and control external inputs on your TV
  power       Control the power state of the TV
  volume      Control the volume of the TV

Use "bravia [command] --help" for more information about a command.
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.