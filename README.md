# hostctl
A CLI tool for blocking websites by managing /etc/hosts on **macOS**.

## Requirements
- macOS
- Go 1.21+

## Installation
```bash
go install github.com/kenvez/hostctl@latest
```

## Usage
Root access is required. Run all commands with `sudo`.
```bash
sudo hostctl block reddit.com    # block a website
sudo hostctl unblock reddit.com  # unblock a website
sudo hostctl list                # list all blocked websites
```

## How it works
hostctl adds and removes entries in `/etc/hosts`, redirecting blocked
domains to `0.0.0.0`. Changes take effect immediately after the 
system DNS cache is flushed.

## License
MIT