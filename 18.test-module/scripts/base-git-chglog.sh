#!/bin/bash

export BIN_DIR=/usr/local/bin
export APP_NAME=git-chglog

download-git-chglog()
{
    APP_BASE=/tmp
    APP_DIR=${APP_BASE}/${APP_NAME}/

    # get os, arch
    OS=$(uname | tr A-Z a-z)
    ARCH=$(uname -m | sed -E 's/^(aarch64|aarch64_be|armv6l|armv7l|armv8b|armv8l)$$/arm64/g' | sed -E 's/^x86_64$$/amd64/g')
    
    # get latest version
    DOWNLOAD_URL=`curl -s https://api.github.com/repos/git-chglog/git-chglog/releases/latest | jq .assets[].browser_download_url | sed 's/"//g' | grep ${OS} | grep ${ARCH}`
    if [ "X${DOWNLOAD_URL}" = "X" ] ; then
        echo "NOT FOUND"
        exit 1
    fi
    APP_DOWNFILE=`basename ${DOWNLOAD_URL}`

    rm -f ${APP_BASE}/${APP_DOWNFILE}
    rm -rf ${APP_DIR} && mkdir -p ${APP_DIR}

    echo "curl -L ${DOWNLOAD_URL} -o ${APP_BASE}/${APP_DOWNFILE}"
    curl -L ${DOWNLOAD_URL} -o ${APP_BASE}/${APP_DOWNFILE}
    tar xzvf ${APP_BASE}/${APP_DOWNFILE} -C ${APP_DIR}
    /bin/cp -f ${APP_DIR}/${APP_NAME} ${BIN_DIR}

    # 한번 더 체크
    if [ ! -f "${APP_DIR}/${APP_NAME}" ] ; then
        exit 1
    fi

    ${APP_DIR}/${APP_NAME} --version
}
