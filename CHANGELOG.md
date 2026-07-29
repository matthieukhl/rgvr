# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added `users` command.
- Added `users list` command that retrieves the list of users from the `/users` endpoint.
- Added `--format` flag to `users list` command to choose between different output formats (table, json).

### Modified

- `teams` command now supports `--format` flag to chosse between different output formats (table, json).

### Removed

- Run function for `auth` command.

## [0.2.0-beta] - 2026-07-28

### Added

- Added `teams` command that retrieves global information (team ID, team name, total numbers count, total users count, etc.).
- Added `teams plans` command that retrieves a team's plans data including number of licenses.

## [0.1.0-beta] - 2026-07-27

### Added

- Added `auth` commands including `login`, `logout` and `status`.
