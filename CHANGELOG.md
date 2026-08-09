# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- [Pull Request #3](https://github.com/matthieukhl/rgvr/pull/3): Added `numbers list` command to retrieve the team's phone numbers.
- [Pull Request #13](https://github.com/matthieukhl/rgvr/pull/13): Added `numbers get <number>` command to retrieve details of a specific number.
- [Pull Request #14](https://github.com/matthieukhl/rgvr/pull/14): Added `ivrs list` command to retrieve all IVR configurations for the user's team.
- [Pull Request *15](https://github.com/matthieukhl/rgvr/pull/15): Added `ivrs get <ivr_id>` command that allows user to retrieve detailed information about a specific ivr.
- [Pull Request *16](https://github.com/matthieukhl/rgvr/pull/16): Added `ivrs scenarios list <ivr_id>` command that allows user to list the scenarios for a specific IVR.
- Added `ivrs scenarios get <ivr_id> <scenario_id>` command that allows user to get a specific scneario of a specific IVR.

### Fixed

- [Pull Request #2](https://github.com/matthieukhl/rgvr/pull/2): Fixed an issue in which the `groups` command returned a null `users` key.
- Fixed an issue that caused the `--format` flag to be available on each command as it was defined as a persistent flag on the root command.

## [0.4.0-beta] - 2026-08-02

### Added

- Added `groups` command to list groups: retrieves call groups (ring groups) configured for your team.  
- Added `groups get <group_id>` command: retrieves detailed information about a specific call group
- Added `groups add-users` command: adds one or more user(s) to a specific group.
- Added `groups remove-users` command: removes on or more user(s) froma a specific group.
- Added `groups set-ring-duration` command: sets the ring duration for a specific user in a specific group.
- Added `groups toggle-access` command: toggles the `is_jumper` flag on a specific group.

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
