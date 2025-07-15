# Hello

This directory contains the "Hello, World!" application from the ["Step by Step
Tutorial"][tutorial] section of the [Service Weaver documentation][docs]. To run
the application, run `go run .`.  Then, curl the `/hello` endpoint (e.g., `curl
localhost:12345/hello?name=Alice`).

```mermaid
%%{init: {"flowchart": {"defaultRenderer": "elk"}} }%%
graph TD
    %% Nodes.
    github.com/thunur/weaver/Main(weaver.Main)
    github.com/thunur/weaver/examples/hello/Reverser(hello.Reverser)

    %% Edges.
    github.com/thunur/weaver/Main --> github.com/thunur/weaver/examples/hello/Reverser
```

[docs]: https://serviceweaver.dev/docs.html
[tutorial]: https://serviceweaver.dev/docs.html#step-by-step-tutorial
