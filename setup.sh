echo "Begining installation of reno"

# checks if go is installed and redirects stdout or stderr to dev/null
if command -v go >/dev/null 2>&1; then
    echo "Go compiler installed"
    go version
    go build -o renoi main.go
    chmod u+x renoi # sets permission for only user to execute
    echo "Build completed, adding reno1 to system path"
    echo "Creating $HOME/.local/bin"

    # Creating a local bin directory in HOME to avoid messing with Unix's bin 
    mkdir -p ~/.local/bin/
    # Copying the newly built binary to the ~/.local/bin
    cp -v ./renoi ~/.local/bin

    echo "Moved renoi to $HOME/.local/bin"

    # We are basically checking if the kind of shell the device uses. 
    # The scripts only checks for bash and zsh shells by default
    if [ "$SHELL" == "/bin/zsh" ];then
      echo "zsh shell detected...writing to ~/.zshrc"
      # Also checks if the .local/bin has already been added to their Path in their config file
      if ! grep -Fxq 'export PATH="$HOME/.local/bin:$PATH"' "$HOME/.zshrc";then
        echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
      fi
    elif [ "$SHELL" == "/bin/bash" ];then
      echo "bash shell detected...writing to ~/.bashrc"
      if ! grep -Fxq 'export PATH="$HOME/.local/bin:$PATH"' "$HOME/.bashrc";then
        echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
      fi
    fi


    echo "renoi successfully added to path"
    echo "Please restart your shell by running source ~/.bashrc or source ~/.zshrc"

else
    echo "Cannot continue setup process, please install the go compiler or add binary found in $PWD/bin to path"
fi

