<div align="center">
  <h1>tu</h1>

  <p>tu is a tiny 🤏🏾 wrapper for your package 📦 manager.</p>
  <p> It'll make managing 🖇️ your packages fun 😊 again!</p>

  <hr/>

  <a href="#why">Why?</a>
  <span> • </span>
  <a href="#install">Install</a>
  <span> • </span>
  <a href="#usage">Usage</a>

  <hr/>

  <p>Tu is <a href="https://en.wikipedia.org/wiki/Swahili_language">swahili</a>.</p>
  <p>It means <em>only</em>, <em>just</em> and <em>simply</em> which all describe this package manager wrapper perfectly.</p>

  <hr/>


[![Go](https://img.shields.io/badge/Made%20with%20Go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=ffffff)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/mistweaverco/tu?style=for-the-badge)](https://github.com/mistweaverco/tu/releases/latest)

  <p></p>

[![GitHub Workflow Build Status](https://img.shields.io/github/actions/workflow/status/mistweaverco/tu/build.yml?style=for-the-badge)](https://github.com/mistweaverco/tu/actions/workflows/build.yml)
[![GitHub Workflow Lint Status](https://img.shields.io/github/actions/workflow/status/mistweaverco/tu/lint.yml?style=for-the-badge&label=lint)](https://github.com/mistweaverco/tu/actions/workflows/lint.yml)
[![codecov](https://img.shields.io/codecov/c/github/mistweaverco/tu?token=yoLIXBNzAZ&style=for-the-badge)](https://codecov.io/github/mistweaverco/tu)


</div>

# Why

Why not simply use aptitute, brew or choco?

I wanted to have a simple package manager wrapper
that I can use to install, remove and update packages from the command line.

If I switch to a different OS/Distro, I can simply continue using `tu` with the
new package manager.

# Install

Grab the latest, greatest and hottest release from the
[releases page](https://github.com/mistweaverco/tu/releases),

or use the following commands depending on your OS:

- [Linux](#install-on-linux)
- [macOS](#install-on-macos)
- [Windows](#install-on-windows)

## Install on Linux

On linux you can use `curl` command to execute the install script directly.

```sh
/bin/bash $(curl -fsSL https://raw.githubusercontent.com/mistweaverco/tu/HEAD/install.sh)
```

## Install on macOS

On macOS you can use `curl` command to execute the install script directly.

```sh
/bin/zsh $(curl -fsSL https://raw.githubusercontent.com/mistweaverco/tu/HEAD/install.sh)
```

## Install on Windows

> [!NOTE]
> To be honest I don't know if package management is a "real" thing on Windows,
> but if it is, you can use `tu` as a wrapper for your choco?!

On Windows you can use the powershell to execute the install script directly.

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/mistweaverco/tu/HEAD/install.ps1 -OutFile tu-install.ps1 | .\tu-install.ps1
```

## Usage

On the command line you can use `tu` to install, remove and update packages.

### Install a package

You can install a package by using the `tu a` command.

> [!TIP]
> Aliases for the `a` subcommand are `add`, `i`, `install` and `get`.

```sh
# Install neovim and zsh
tu a neovim zsh
```
> [!NOTE]
> If you want to install the packages via Homebrew, you can pass the `-b` or `--brew` flag to `tu a`.
> This only works on macOS and Linux.

> [!TIP]
> Maybe you want to update the mirrors before installing packages.
> Easy peasy, just pass the `-s` or `--sync` flag to `tu a`.

```sh
# Update the mirrors and install neovim and zsh
tu a -s neovim zsh
```
> [!NOTE]
> If you want to sync via Homebrew,
> you can pass the `-b` or `--brew` flag to `tu a`.
> This only works on macOS and Linux.

### Remove a package

You can remove a package by using the `tu d` command.

> [!TIP]
> Aliases for the `d` subcommand are `remove`, `uninstall`, `rm`, `delete` and `del`.

```sh
# Remove a neovim and zsh
tu d neovim zsh
```

> [!NOTE]
> If you want to delete a package from Homebrew,
> you can pass the `-b` or `--brew` flag to `tu a`.
> This only works on macOS and Linux.

### Update all packages

You can update all packages by using the `tu u` command.

> [!TIP]
> Aliases for the `u` subcommand are `update`, `upgrade` and `up`.

```sh
# Update all packages
tu u
```

> [!NOTE]
> If you want to update all packages installed via Homebrew,
> you can pass the `-b` or `--brew` flag to `tu a`.
> This only works on macOS and Linux.

> [!TIP]
> Maybe you want to update the mirrors before updating all packages.
> Easy peasy, just pass the `-s` or `--sync` flag to `tu u`.

### Search for a package

You can search for a package by using the `tu q` command.

> [!TIP]
> Aliases for the `q` subcommand are `query` and `search`.

```sh
# Search for a package called flameshot
tu q flameshot
```

> [!NOTE]
> If you want to search for a package via Homebrew,
> you can pass the `-b` or `--brew` flag to `tu q`.
> This only works on macOS and Linux.

> [!TIP]
> Maybe you want to update the mirrors before searching for a package.
> Easy peasy, just pass the `-s` or `--sync` flag to `tu q`.

