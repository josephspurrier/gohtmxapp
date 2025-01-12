# gohtmxapp

This is a sample application that uses a [Go](https://go.dev/) backend to serve content using [HTMX](https://htmx.org/) and [Templ](https://templ.guide/). For styling, it uses [Tailwind](https://tailwindcss.com/) and [Flowbite](https://flowbite.com/). It supports live reloading (using [Air](https://github.com/air-verse/air)) and browser reloading (simply using server-sent events). Assets are hashed using [hashfs](https://github.com/benbjohnson/hashfs). There are helpers in Go for setting active pages based on page URL.

The reason why I like this approach is it keeps the majority of your front-end logic in Go and HTML so it does not require a front-end framework like React or Vue, it uses the HTMX library to provide many of the features of dynamic web applications like loading content without having to do a full page render.

**Note:** The UI itself is very basic, it's not a real web application, it's just designed to show how the different components work together in code.

## Quick Start

To test the application locally, you can run these commands. You will need to install Go and NPM before running these commands.

```bash
# Clone the repo.
git git@github.com:josephspurrier/gohtmxapp.git

# Change to the project directory.
cd gohtmxapp

# Create the config.
cp sample.env .env

# Install the dependencies.
make install

# Start the web server with reload capability.
make watch

# Open your browser to view the application: http://localhost:8080/dashboard
# Most of the pages are placeholders, but a few pages do have content: Dashboard, Test Form, and Settings.
```

## Overview

This projects uses a combination of tools to make management of dependencies, reloading, and a few other things easier.

- **Makefile** contains helpful commands to interact with the project during development.
- **Air** is used to rebuild the Go binary with files changes.
- **package.json** contains dependencies for this project. There is no advanced system outside of a few commands like creating a tailwind CSS output file.
- **templ** is used for templating in Go. You may want to set up your IDE using these [suggestions](https://templ.guide/developer-tools/ide-support). If you're using VSCode and don't want file nesting in the explorer window when using the `templ` extension, you can add this to your VSCode settings.json: `"explorer.fileNesting.enabled": false`.

To install any dependencies using npm, you should install them as dev dependencies like this: `npm install -D htmx.org@2`. The `make install` command does call the npm scripts to copy the dependencies to the correct `web/assets/js` folder.

This project was created initially using Go-Blueprint: https://docs.go-blueprint.dev/creating-project/project-init/

You should update the assets/favicon folder with your own Favicon: http://realfavicongenerator.net.

### Config 

You should add to your `.env` or `.envrc` file.

For .env:

```bash
# App Config.
PORT=8080
HASH_ASSETS=true
APPLICATION_URL="https://app-name-here.fly.dev"
HOT_RELOAD=true
```

For .envrc ([direnv](https://direnv.net/) required):

```bash
# App Config.
export PORT=8080
export HASH_ASSETS=true
export APPLICATION_URL="https://app-name-here.fly.dev"
export HOT_RELOAD=true
```

### MakeFile

You can run these commands from your terminal:

```bash
# Install and setup the project dependencies.
make install

# Start the web server and watch the files for changes.
make watch

```

Other commands:

```bash
# Clean up binary from the last build.
make clean

# Build the application.
make build

# Run the application.
make run

# Run the test suite
make test
```