# music-guide
Music theory and guitar practice tool

## Installation
Here's a list of everything you'll need to compile the project.
The application can be built from scratch or with Docker.

### Building and running from scratch
#### Prerequisites
go 1.23 or higher: \
https://go.dev/doc/install
templ:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```
air (optional):
```bash
go install github.com/air-verse/air@latest
```
npm:
https://docs.npmjs.com/downloading-and-installing-node-js-and-npm
tailwind:
```bash
npm install -D tailwindcss
npm install -D @tailwindcss/forms
```
#### Build
Generate templates, build tailwind, and compile the project:
```bash
make build
```
Build and start the website with air:
```bash
make build-watch
```

#### Run
Start the website:
```bash
make run
```
#### Clean
Remove the compiled binaries with:
```bash
make clean
```

### Building and running with Docker
#### Build
Build a docker image:
```bash
make docker-build
```
#### Run
Start the website:
```
make docker-run
```
Run a watchful container with air:
```bash
make docker-watch
```
