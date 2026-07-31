# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added `groups` command to list groups: retrieves call groups (ring groups) configured for your team.  
- Added `groups get <group_id>` command: retrieves detailed information about a specific call group

## [0.3.0-beta] - 2026-07-31

### Added

- Added `users` command.
- Added `users list` command that retrieves the list of users from the `/users` endpoint.
- Added `--format` flag to `users list` command to choose between different output formats (table, json).
- Added `users get <user_id>` command to retrieve a user's information based on a user ID.
- Added `users plannings get <user_id>` command to retrieve a user's planning based on a user ID.
- Added `users presences <user_id>` command to retrieve a user's presence status based on a user ID.
- Added `users snooze log <user_id>` command to retrieve a specific user's snooze log.

### Modified

- `teams` command now supports `--format` flag to chosse between different output formats (table, json).
- `teams plan` command now supports `--format` flag to chosse between different output formats (table, json).

### Removed

- Run function for `auth` command.

## [0.2.0-beta] - 2026-07-28

### Added

- Added `teams` command that retrieves global information (team ID, team name, total numbers count, total users count, etc.).
- Added `teams plans` command that retrieves a team's plans data including number of licenses.

## [0.1.0-beta] - 2026-07-27

### Added

- Added `auth` commands including `login`, `logout` and `status`.
