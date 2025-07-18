# Service Weaver [![Go Reference](https://pkg.go.dev/badge/github.com/thunur/weaver.svg)](https://pkg.go.dev/github.com/thunur/weaver)

## 🚨 Important Announcement

> **Service Weaver** began as an exploratory initiative to understand the challenges of developing, deploying, and maintaining distributed applications. We were excited by the strong interest from the developer community, which led us to open-source the project.
>
> We greatly appreciate the continued advocacy and support of the Service Weaver community. However, we realized that it was hard for users to adopt Service Weaver directly since it required rewriting large parts of existing applications. Therefore, Service Weaver did not see much direct use, and **effective December 5, 2024**, we will transition Service Weaver into maintenance mode. After this date, for the next 6 months, we will only push critical commits to the GitHub repository, respond to critical issues, merge critical pull requests, and patch new releases. We recommend that users fork the repository and report any issues preventing them from maintaining a stable version of the code.
>
> **On June 6, 2025**, we plan to permanently freeze and archive the GitHub repository, after which no new commits or releases will be made.

Service Weaver is a programming framework for writing, deploying, and managing
distributed applications. You can run, test, and debug a Service Weaver
application locally on your machine, and then deploy it to the
cloud with a single command.

```bash
$ go run .                      # Run locally.
$ weaver gke deploy weaver.toml # Run in the cloud.
```

Visit [https://serviceweaver.dev][website] to learn more about Service Weaver.

## Installation and Getting Started

Visit [https://serviceweaver.dev/docs.html][docs] for installation
instructions and information on getting started.

## Contributing

Please read our [contribution guide](./CONTRIBUTING.md) for details on how
to contribute.

[website]: https://serviceweaver.dev
[docs]: https://serviceweaver.dev/docs.html
