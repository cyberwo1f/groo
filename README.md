# groo: fast access to a website on git remote

**groo** is a shell script to help you easily access to a website on git remote in current.

## Typical usage

In a git managed directory: (default target is “origin”)
```bash
groo
```

If you want to specify remote:
```bash
groo -r [remote]
```

## Supported Domains
- github.com
- gitlab.com

## Installation

### Installation on OS X using Homebrew

First, you need to run the tap command:
```bash
brew tap cyberwo1f/groo
```

Next, run the install command:
```bash
brew install groo
```

**Currently, only Mac OS is supported.**

## License
MIT
