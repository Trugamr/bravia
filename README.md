# Bravia TV CLI

A command-line interface tool for controlling Sony Bravia TVs. Control your TV's power, volume, inputs, and apps directly from your terminal.

## Features

- **CLI Tool**: Control your TV from the command line
  - Power control (on/off/status)
  - Volume adjustment
  - Input switching
  - App management
- **Web Remote**: Modern web-based remote control interface
  - Full-screen responsive grid layout
  - Real-time TV state updates via Server-Sent Events (SSE)
  - All IRCC remote control commands
  - App launcher with icons
  - Number pad, playback controls, and more
- **API Library**: Use the API package in your Go projects
- Secure PSK authentication

## Installation

Download the latest binary for your platform from the [releases page](https://github.com/trugamr/bravia/releases).

## Building Locally

This project uses [Task](https://taskfile.dev) for build automation and [GoReleaser](https://goreleaser.com) for releases.

### Quick Start

```bash
# Install dependencies
task install

# Build for current platform
task dev:build

# Build with GoReleaser (all platforms)
task build

# Build for current platform only (faster)
task release:local
```

### Available Tasks

```bash
task --list              # List all available tasks
task build              # Build using goreleaser (all platforms)
task release:local      # Build for current platform only
task dev:build          # Quick build for development
task run:cli            # Run bravia CLI
task run:remote         # Run remote server
task clean              # Clean build artifacts
task test               # Run tests
task fmt                # Format code
task lint               # Run linter
```

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

## Web Remote

The project includes a web-based remote control interface with a modern, responsive design:

```bash
# Start the remote server
remote

# Or with custom configuration
remote --base-url="http://your-tv-ip" --psk="your-pre-shared-key" --port=3000
```

The web remote provides:
- Full remote control functionality via IRCC commands
- Real-time TV state monitoring (power, volume, mute status)
- App launcher with application icons
- Input selection with visual feedback
- Number pad for channel entry
- Playback controls (play, pause, stop, rewind, forward, etc.)
- Power controls (power, wake, sleep)
- Quick access to HDMI inputs

Access the remote at `http://localhost:3000` (or your configured port).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.