#!/bin/bash
#
# https://github.com/go-redis/redis/blob/v8/scripts/tag.sh

help() {
    cat <<- EOF
Usage: TAG=tag $0
Updates version in go.mod files and pushes a new brash to GitHub.
VARIABLES:
  TAG        git tag, for example, v1.0.0
EOF
    exit 0
}

validate_vtag() {
    TAG=${1}
    echo ">>validate_vtag>> tag[${TAG}]...1"
    TAG_REGEX="^v(0|[1-9][0-9]*)$"
    if ! [[ "${TAG}" =~ ${TAG_REGEX} ]]; then
        echo "<<validate_vtag<< tag[${TAG}]...r1"
        return 1
    fi
    echo "<<validate_vtag<< tag[${TAG}]...r0"
    return 0
}

validate_tag() {
    TAG=${1}
    echo ">>validate_tag>> tag[${TAG}]...1"
    TAG_REGEX="^v(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)(\\-[0-9A-Za-z-]+(\\.[0-9A-Za-z-]+)*)?(\\+[0-9A-Za-z-]+(\\.[0-9A-Za-z-]+)*)?$"
    if ! [[ "${TAG}" =~ ${TAG_REGEX} ]]; then
        echo "<<validate_tag<< tag[${TAG}]...r1"
        return 1
    fi
    echo "<<validate_tag<< tag[${TAG}]...r0"
    return 0
}

check_tag() {
    TAG=${1}

    if [ -z "$TAG" ]
    then
        printf "TAG is required\n\n"
        help
    fi

    validate_tag ${TAG}
    if [ $? -eq 1 ] ; then
        printf "TAG is not valid: ${TAG}\n\n"
        exit 1
    fi

    TAG_FOUND=`git tag --list ${TAG}`
    if [[ ${TAG_FOUND} = ${TAG} ]] ; then
        printf "tag ${TAG} already exists\n\n"
        exit 1
    fi

    if ! git diff --quiet
    then
        printf "working tree is not clean\n\n"
        git status
        exit 1
    fi
}
