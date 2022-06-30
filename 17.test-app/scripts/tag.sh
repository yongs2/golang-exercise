#!/bin/bash
#
# https://github.com/go-redis/redis/blob/v8/scripts/tag.sh

BASE_NAME=$(dirname $0)
if [ "X${BASE_NAME}" = "X" ] ; then
   BASE_PATH=${PWD}
else
   BASE_PATH=${PWD}/${BASE_NAME}
fi

# include
. ${BASE_PATH}/base-tag.sh
. ${BASE_PATH}/base-git-chglog.sh

set -e

check_tag ${TAG}

if [ ! -f ${BIN_DIR}/${APP_NAME} ] ; then
    download-git-chglog
fi

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo ">> Generate CHANGELOG.md with ${TAG} on branch ${CURRENT_BRANCH}"
# https://github.com/git-chglog/git-chglog/blob/master/Makefile
${BIN_DIR}/${APP_NAME} --next-tag ${TAG} -o CHANGELOG.md
git add CHANGELOG.md
git commit -m "chore: tag $TAG (tag.sh)"
git tag ${TAG}
git push origin ${CURRENT_BRANCH} ${TAG}

PACKAGE_DIRS=$(find . -mindepth 2 -type f -name 'go.mod' -exec dirname {} \; \
  | grep -E -v "example|internal" \
  | sed 's/^\.\///' \
  | sort)

for dir in $PACKAGE_DIRS
do
    printf "tagging ${dir}/${TAG}\n"
    git tag ${dir}/${TAG}
    git push origin ${CURRENT_BRANCH} ${dir}/${TAG}
done
