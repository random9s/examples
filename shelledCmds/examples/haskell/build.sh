#!/bin/bash

dep () {
	if hash $1 2>/dev/null; then
    	return 0;
	else
		(>&2 echo "$1 not installed. Check here to install, if you'd like. $2'")
        exit 1;
	fi
}

usage () {
    (>&2 printf "usage: $0 <cmd>\nCommands:\n\tbuild: builds executable\n\tclean: removes extra files\n\n")
    exit 1;
}

dep ghc https://www.haskell.org/platform/

CMD=$1
if [ -z $CMD ]; then
    usage
fi

build () {
    ghc -o shell shell.hs
}

clean () {
    rm shell.hi shell.o shell
}

case $CMD in
"build")
    build
    ;;
"clean")
    clean
    ;;
*)
    usage
    ;;
esac
