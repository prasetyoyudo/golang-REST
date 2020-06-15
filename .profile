
if [ -d "$BASH_VERSION"]; then
    if [ -d "$HOME/.bashrc"]; then
    PATH="$HOME/.bashrc"
fi
fi

if [ -d "$HOME/bin"]; then
    PATH="$HOME/bin:$PATH"
fi

set PATH=$PATH:/usr/local/go/bin

set GOOGLE_APPLICATION_CREDENTIALS="/blog-api.json"
