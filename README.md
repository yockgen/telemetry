# IaaS - Platform Components

## Building and testing

This project uses [Earthly](https://earthly.dev) for CI/CD targets.

Building this project:

`earthly +build-api`

Running all the tests:

`earthly +test`

## Directory Structure

```bash
iaas-platform
├── schemas                 # OpenAPI yaml and protobuf files
├── build                   # Build artifacts pushed by build containers
├── cmd                     # Go-only main entry points executables agents and managers
├── configs                 # Configuration artifacts for services and agents
├── docker                  # Holds container definitions and docker-compose files for artifacts
├── .github/workflows       # All git actions for CI/CD workflows
├── helm                    # Helm charts and default values files
├── internal                # Go-only source code implementation for executables, libraries and go-tests
├── src                     # Open: for Py can we keep similar go structure or dedicated /src
│   └── package_name
│       └── __init__.py
├── pkg                     # Go-only Library code for external applications to use
├── scripts                 # General purpose for scripts
├── tests                   # all api-level and e2e, system testing
├── ext                     # Imported 3rd party tools and utilities
├── tools                   # Support tools, that may be importing from /pkg and /internal
├── vendor                  # Application dependencies and go modules
├── README.md
├── requirements.txt        # Python pre-requisits
├── setup.cfg
├── setup.py
├── .editorconfig
├── .gitattributes
├── .gitignore
├── .golangci.yml
├── go.mod
├── go.sum
├── .snyk
├── LICENSE
├── Magefile
└── tox.ini
```

## Edge IaaS Components
1. [Telemetry](/helm/edge-iaas-platform/platform-director/telemetry)
