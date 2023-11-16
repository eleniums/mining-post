#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

cd $BASE_DIR

NAME=miningpost
VERSION=$(git describe --tags --always --long --dirty)
RELEASE=$BASE_DIR/release/$VERSION

echo "Creating release of $NAME with version: $VERSION"

# build all binaries
./scripts/build.sh

# create release folder
rm -rf $RELEASE
mkdir -p $RELEASE

# package windows binary
echo "Packaging Windows binary..."
mv "$NAME"_windows_amd64.exe $RELEASE/$NAME.exe
cd $RELEASE
zip "$NAME"_windows.zip $NAME.exe
rm -rf $NAME.exe
cd $BASE_DIR

# package linux binary
echo "Packaging Linux binary..."
mv "$NAME"_linux_amd64 $RELEASE/$NAME
cd $RELEASE
tar -czvf "$NAME"_linux.tar.gz $NAME
rm -rf $NAME
cd $BASE_DIR

# package mac binary
echo "Packaging macOS binary..."
mv "$NAME"_darwin_amd64 $RELEASE/$NAME
cd $RELEASE
tar -czvf "$NAME"_mac.tar.gz $NAME
rm -rf $NAME
cd $BASE_DIR

echo "Release of $NAME created: $RELEASE"