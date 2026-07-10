## Development setup

```sh
git clone https://github.com/inflame-ue/pastiche
cd pastiche
go build .
```

The project has no external runtime dependencies beyond the formatter tools you choose to use.

### Running the daemon during development

```sh
# Create a default config if you haven't already
go run . configure

# Run the daemon in the foreground
go run . daemon
```

### Running tests

```sh
go test ./...
```

### Project layout

```
cmd/      — cobra commands (root, configure, install, status, service)
internal/ — library code (config, pipeline, formatter, trigger, tui)
```
