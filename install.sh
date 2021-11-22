#!/usr/bin/bash

echo "Installing JT CLI ..."

OS=linux

case "$OSTYPE" in
  darwin*)  OS=darwin ;; 
  linux*)   OS=linux ;;
  *)        OS=linux ;;
esac

curl -LO https://github.com/jarvis-technologies/cli/releases/download/v0.1.0/jt-$OS-amd64
sudo mv ./jt-$OS-amd64 /usr/local/bin/jt
chmod +x /usr/local/bin/jt

