# Preparation

## Prerequisites

Expected background knowledge and skills for the workshop:

- Have 1-year hands-on experience in other programming language;
- Know how to work with the Command Line Interface in Linux or OSX;
- basic knowledge about Git.

Recommended: you should have an Github or Gitlab account.

## Your workstation

- Linux (can be a virtual machine with VirtualBox or VMWare) or OSX;
- Installed:

  - Golang;
  - Git;
  - Docker.

- An IDE or code editor to work with Golang, e.g.:
   
  - vscode,
  - Jetbrains Goland.

- for the exercises with SQL and no{\small SQL} databases:

  - Postgres CLI,
  - MongoDB CLI.

You will find on [github.com/golang/go/wiki/IDEsAndTextEditorPlugins](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) information about how to configure your favorite editor to work with Golang.

## Guides

### Ubuntu

We recommend Ubuntu, one of the LTSes - [wiki.ubuntu.com/Releases](https://wiki.ubuntu.com/Releases).

1. Install Golang, following the instructions from [github.com/golang/go/wiki/Ubuntu](https://github.com/golang/go/wiki/Ubuntu)

   ```bash
   sudo add-apt-repository ppa:longsleep/golang-backports
   sudo apt-get update
   sudo apt-get install golang-go
   ```

   ```bash
   go version
   ```

2. Install vscode ([code.visualstudio.com/](https://code.visualstudio.com)) or Goland ([jetbrains.com/go/](https://www.jetbrains.com/go/)):

   ```bash
   # vscode with snap
   # https://code.visualstudio.com/docs/setup/linux
   sudo snap install --classic code
   ```

   For Jetbrain Goland follow the standalone installation: [www.jetbrains.com/help/go/installation-guide.html](https://www.jetbrains.com/help/go/installation-guide.html#c2dfc2).

3. Install Git:

   ```bash
   sudo apt-get update
   sudo apt-get install -y git
   ```

4. Install Docker after [docs.docker.com/engine/install/ubuntu](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository), copy and paste to your terminal:

   ```bash
   # install necessary packages
   sudo apt-get update
   sudo apt-get install -y ca-certificates curl gnupg

   # Add Docker’s official GPG key:
   sudo install -m 0755 -d /etc/apt/keyrings
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg \
       | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
   ```

   ```bash
   # Setting up the repository:
   echo \
      "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
      sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
   ```
  
   Install `docker-ce`:

   ```bash
   sudo apt-get update
   sudo apt-get install docker-ce docker-ce-cli
   ```

   Check whether it works:

   ```bash
   sudo docker run hello-world
   ```

## MacOS

1. Install *Homebrew*, a package manager for MacOS, follow the instructions from the official website - [brew.sh](https://brew.sh/).

2. With *Homebrew*, you can easily install Golang:

   ```bash
   brew install golang
   ```

   ```bash
   go version
   ```

3. Choose your IDE - *vscode* (recommended for this workshop) or *Goland* (goto IDE for most of Golang devs):

   ```bash
   # for vscode
   brew install --cask visual-studio-code
   ```

   ```bash
   # for goland community edition
   brew install --cask goland
   ```

3. To install Docker, follow instructions from the oficial website - [docs.docker.com/desktop/install/mac-install/](https://docs.docker.com/desktop/install/mac-install/).

## References

- The Missing Semester of Your CS Education: [missing.csail.mit.edu](https://missing.csail.mit.edu/);

- Quick start with \emph{Goland}: [jetbrains.com/help/idea/quick-start-guide-goland.html](https://www.jetbrains.com/help/idea/quick-start-guide-goland.html);

- Quick start with VSCode: [learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code). 
