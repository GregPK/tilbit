#!/usr/bin/env bash

# Script based on https://gist.githubusercontent.com/makeworld-the-better-one/e1bb127979ae4195f43aaa3ad46b1097/raw/90e2c8bd9ed7c4332b902a2765dd0ff502343977/go-build-all.sh

type setopt >/dev/null 2>&1

contains() {
    # Source: https://stackoverflow.com/a/8063398/7361270
    [[ $1 =~ (^|[[:space:]])$2($|[[:space:]]) ]]
}

SOURCE_FILE=$(echo "$@" | sed 's/\.go//')
CURRENT_DIRECTORY="${PWD##*/}"
VERSION=${VERSION:-99.88.77}
OUTPUT="dist/${SOURCE_FILE:-$CURRENT_DIRECTORY}" # if no src file given, use current dir name
FAILURES=""

# You can set your own flags on the command line
FLAGS=${FLAGS:-"-ldflags=\"-s -w\""}

# A list of OSes to not build for, space-separated
# It can be set from the command line when the script is called.
NOT_ALLOWED_OS=${NOT_ALLOWED_OS:-"aix android dragonfly ios js  illumos plan9 solaris freebsd netbsd openbsd"}
NOT_ALLOWED_ARCH=${NOT_ALLOWED_ARCH:-"s390x mips mipsle mips64 mips64le ppc ppc64 ppc64le riscv64"}

# Get all targets
while IFS= read -r target; do
    GOOS=${target%/*}
    GOARCH=${target#*/}
    BIN_FILENAME="${OUTPUT}-${VERSION}-${GOOS}-${GOARCH}"

    if contains "$NOT_ALLOWED_OS" "$GOOS" ; then
        continue
    fi
    if contains "$NOT_ALLOWED_ARCH" "$GOARCH" ; then
        continue
    fi

    # Modified: cut down the arm architecture to just 1
    # Check for arm and set arm version
    if [[ $GOARCH == "arm" ]]; then
        # Set what arm versions each platform supports
        if [[ $GOOS == "darwin" ]]; then
            arms="7"
        elif [[ $GOOS == "windows" ]]; then
            arms="7"
        elif [[ $GOOS == *"bsd"  ]]; then
            arms="7"
        else
            # Linux goes here
            arms="7"
        fi

        # Now do the arm build
        for GOARM in $arms; do
            BIN_FILENAME="${OUTPUT}-${VERSION}-${GOOS}-${GOARCH}${GOARM}"
            if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
            CMD="GOARM=${GOARM} GOOS=${GOOS} GOARCH=${GOARCH} go build $FLAGS -o ${BIN_FILENAME} $@"
            echo "${CMD}"
            eval "${CMD}" || FAILURES="${FAILURES} ${GOOS}/${GOARCH}${GOARM}"
        done
    else
        # Build non-arm here
        if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
        CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build $FLAGS -o ${BIN_FILENAME} $@"
        echo "${CMD}"
        eval "${CMD}" || FAILURES="${FAILURES} ${GOOS}/${GOARCH}"
    fi
done <<< "$(go tool dist list)"

if [[ "${FAILURES}" != "" ]]; then
    echo ""
    echo "${SCRIPT_NAME} failed on: ${FAILURES}"
    exit 1
fi
