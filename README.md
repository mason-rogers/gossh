# gossh

A simple terminal-based SSH host manager with support for nested groups, jumphosts, and interactive host selection.

[![asciicast](https://asciinema.org/a/8IyIwNWVApcpWEO8eIij9verZ.svg)](https://asciinema.org/a/8IyIwNWVApcpWEO8eIij9verZ)

## Features

- 📁 Hierarchical host groups
- 🔄 Jump host support
- 🎯 Interactive host selector
- ⌨️ Direct connection using host paths
- 🔧 YAML configuration
- 📥 Import from Termius (Coming soon!)

## Installation

Download the pre-compiled binary from the [releases page](https://github.com/mason-rogers/gossh/releases).

### macOS & Linux:
```bash
curl -L https://github.com/mason-rogers/gossh/releases/latest/download/gossh_/gossh_$(uname -s)_$(uname -m).tar.gz".tar.gz | tar xz
sudo mv gossh /usr/local/bin
```

### Windows:
Download the zip file for Windows from the [releases page](https://github.com/mason-rogers/gossh/releases) and extract `gossh.exe` to a location in your PATH.

### Build from Source
```bash
go install github.com/mason-rogers/gossh@latest
```

## Usage

### Interactive Mode
```bash
gossh
```
Launches an interactive selector showing all available hosts organized by groups.

### Direct Connection
```bash
gossh production/europe/web01
```
Directly connects to a host using its full path.

### Version Information
```bash
gossh version
```

## Configuration

Create `~/.gossh.yaml`:

```yaml
jumphosts:
  - name: bastion
    hostname: jump.viction.dev
    user: jumper
    port: 22
    keyfile: ~/.ssh/id_ed25519

groups:
  - name: production
    groups:
      - name: eu
        hosts:
          - name: web01
            hostname: eu-web01.viction.dev
            user: root
            port: 22
            keyfile: ~/.ssh/id_ed25519
            jumphost: bastion
      - name: us
        hosts:
          - name: web01
            hostname: us-web01.viction.dev
            user: root
            port: 22
            keyfile: ~/.ssh/id_ed25519
            jumphost: bastion
```

### Configuration Structure

- `jumphosts`: List of jump hosts (bastions) that can be referenced by hosts
    - `name`: Unique identifier for the jump host (required)
    - `hostname`: Jump host address (required)
    - `user`: SSH username (default: root)
    - `port`: SSH port (default: 22)
    - `keyfile`: Path to SSH private key (optional)

- `groups`: Nested groups containing hosts
    - `name`: Group name (required)
    - `groups`: Nested subgroups (optional)
    - `hosts`: List of hosts in this group (optional)
        - `name`: Host identifier (required)
        - `hostname`: Host address (required)
        - `user`: SSH username (default: root)
        - `port`: SSH port (default: 22)
        - `keyfile`: Path to SSH private key (optional)
        - `jumphost`: Reference to a jumphost by name (optional)

## Development

### Requirements
- Go 1.23 or later

### Local Development
```bash
# Build locally
go build

# Test release build
goreleaser release --snapshot --clean
```

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)