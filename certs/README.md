These certificates are to be used only in development mode.

## Generating certificate
First install certutil
```bash
sudo apt install libnss3-tools
```
Then you can install using Homebrew for Linux
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
```bash
brew install mkcert
```
or build from source(requires Go 1.13+):
```bash
git clone https://github.com/FiloSottile/mkcert && cd mkcert
go build -ldflags "-X main.Version=$(git describe --tags)"
```

Then, to install `mkcert` do:
```bash
mkcert -install
```
Then, with `mkcert` installed, go to your project folder and run:
```bash
mkcert localhost
```

