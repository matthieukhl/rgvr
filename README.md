# rgvr

`rgvr` is a command-line tool written in Go for interacting with [Ringover](https://www.ringover.com)'s public API — manage and query your teams, users, calls, and more, directly from your terminal.

> **Disclaimer**: rgvr is an independent, community-built tool. It is not affiliated with, endorsed by, or supported by Ringover.  
> **Status**: Work in progress. Commands are being added incrementally as Ringover's API endpoints are covered. Expect breaking changes until a `v1.0.0` release.

## Prerequisites

- Go 1.21+ (for building from source)
- A valid Ringover API key (create one from your Ringover Dashboard under **Developer > API key**)

### API key permissions

Ringover API keys have granular permissions across 7 categories (Calls, Contacts, Users, Numbers, IVRs, Conversations, Empower), each configurable as None, Read, Write, or Read+Write, plus a separate **Monitoring** toggle that controls whether a key can access your own data only (Monitoring OFF) or your entire team's data (Monitoring ON).

`rgvr` commands map to different permission categories and monitoring requirements depending on the endpoint they call — for example, `rgvr users list` requires `Users Read`, while `rgvr groups toggle-access` requires Monitoring to be enabled. **To use all commands, your API key should have Read+Write access to every category, with Monitoring enabled.**

If a command fails with a `401 Unauthorized` error, check that your key has the required permission and/or Monitoring enabled for that specific operation. See the [Ringover API documentation](https://developer.ringover.com) for the exact requirements of each endpoint.

## Installation

### From source

```bash
git clone https://github.com/matthieukhl/rgvr.git
cd rgvr
make build
```

This produces a `rgvr` binary at the root of the directory. Move it somewhere on your `$PATH`, for example:

```bash
mv rgvr /usr/local/bin/rgvr
```

### Via `go install`

```bash
go install github.com/matthieukhl/rgvr@latest
```

## Getting started

Authenticate with your Ringover API key:

```bash
rgvr auth login
```

This will prompt for your API key and your Ringover region (EU or US), and store them locally in `~/.config/rgvr/config.yaml`.

Verify you're properly authenticated:

```bash
rgvr auth status
```

## Authentication

`rgvr` resolves your API key in the following order of precedence:

1. `--api-key` flag
2. `RGVR_API_KEY` environment variable
3. `~/.config/rgvr/config.yaml` (set via `rgvr auth login`)

This means you can use `rgvr` in scripts or CI/CD pipelines without ever running `auth login`, by exporting the environment variable instead:

```bash
export RGVR_API_KEY=your-api-key-here
rgvr teams
```

## Development

Build the binary:

```bash
make build
```

## Changelog

See [Changelog](CHANGELOG.md).

## Roadmap

- 🟢 `/teams` endpoint
- 🟡 `/users` endpoint (`GET` endpoints have been implemented)
- 🟢 `/groups` endpoint
- 🟢 `/numbers` endpoint
- 🟢 `/ivrs` endpoint
- 🟢 `/tags` endpoint
- 🔴 `/conferences` endpoint
- 🔴 `/calls` endpoint
- 🔴 `/campaigns` endpoint
- 🔴 `/channels` endpoint
- 🔴 `/contacts` endpoint
- 🔴 `/conversations` endpoint
- 🔴 `/messages` endpoint
- 🔴 `/sms` endpoint
- 🔴 `/profiles` endpoint
- 🔴 `/blacklist` endpoint
- 🔴 `/webhook` endpoint
- 🔴 `/callbacks` endpoint
- 🔴 `/surveys` endpoint
- 🔴 `/snoozes` endpoint
- 🔴 `/transcriptions` endpoint
- 🔴 `/empower` endpoint
- 🔴 `/whatsapp` endpoint
- 🔴 `/task` endpoint
- 🔴 `/mcp` endpoint

## Links

- [Ringover API documentation](https://developer.ringover.com)

## License

rgvr is licensed under **AGPL-3.0-only**. See [LICENSE](LICENSE) for the full license text.
