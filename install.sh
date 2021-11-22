#!/usr/bin/bash

echo "Installing JT CLI ..."

OS=linux
ARCH="$(arch)"

TAG=$(curl -s "https://api.github.com/repos/jarvis-technologies/cli/releases/latest" |
    grep '"tag_name":' |
    sed -E 's/.*"([^"]+)".*/\1/')

case "$OSTYPE" in
  darwin*)  OS=darwin ;; 
  linux*)   OS=linux ;;
  *)        OS=linux ;;
esac

case "$ARCH" in
   x86_64) ARCH=amd64 ;;
   arm) ARCH=arm64 ;;
esac

curl -LO https://github.com/jarvis-technologies/cli/releases/download/$TAG/jt-$OS-$ARCH
sudo mv ./jt-$OS-$ARCH /usr/local/bin/jt
chmod +x /usr/local/bin/jt

