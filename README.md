# dtrack-auditor-go

dtrack-auditor-go is a Go program to facilitate usage of [DependencyTrack](https://dependencytrack.org/) in CI systems or locally, failing the build based on different parameters.

dtrack-auditor-go relies on the metrics API to analyze results.

## Features

1. Auto mode for project creation given project name and version. Creates new project with version if already not found.
2. Fail based on severity type (critical, high, medium, low, unassigned) and numbers, example: if number of critical is higher or equal to 10.
3. Return 0 or 1 exit status after analysis.

## Quick Install

dtrack-auditor-go can be installed in 2 separate ways:

1. Download binary from Github Releases
2. Docker image for use in containerized CI systems (`akath19/dtrack-auditor-go`)

## Configuration

dtrack-auditor-go is configured via environment variables:

* `API_URL`: Base URL of dependency-track to connect to.
* `API_TOKEN`: Token used for authentication.
* `API_TIMEOUT`: Timeout in seconds to use for API calls, defaults to 30
* `PROJECT_NAME`: Name of the project to find or create if `autocreate` is enabled.
* `PROJECT_VERSION`: Version of the project to find or create if `autocreate` is enabled.
* `PROJECT_AUTOCREATE`: Toggle to automatically create the project in dependency-track if it does not exist, token must have `PORTFOLIO_MANAGEMENT` or `PROJECT_CREATION_UPLOAD` permissions to create the project successfully.
* `BOM_PATH`: Relative path to the BOM file to analyze.
* `CRITICAL_THRESHOLD`: Numeric threshold of CRITICAL vulnerabilities (inclusive), setting to `0` ignores the threshold entirely, defaults to `0`.
* `HIGH_THRESHOLD`: Numeric threshold of HIGH vulnerabilities (inclusive), setting to `0` ignores the threshold entirely, defaults to `0`.
* `MEDIUM_THRESHOLD`: Numeric threshold of MEDIUM vulnerabilities (inclusive), setting to `0` ignores the threshold entirely, defaults to `0`.
* `LOW_THRESHOLD`: Numeric threshold of LOW vulnerabilities (inclusive), setting to `0` ignores the threshold entirely, defaults to `0`.
* `UNASSIGNED_THRESHOLD`: Numeric threshold of UNASSIGNED vulnerabilities (inclusive), setting to `0` ignores the threshold entirely, defaults to `0`.

## Upcoming Features

* License violations threshold
* Security violations threshold
* Overall violations threshold
