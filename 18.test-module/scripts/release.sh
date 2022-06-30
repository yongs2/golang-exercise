#!/bin/bash
#
# https://github.com/go-redis/redis/blob/v8/scripts/release.sh

BASE_NAME=$(dirname $0)
if [ "X${BASE_NAME}" = "X" ]; then
  BASE_PATH=${PWD}
elif [ "${BASE_NAME}" = "." ]; then
  BASE_PATH=${PWD}
else
  BASE_PATH=${PWD}/${BASE_NAME}
fi
echo "BASE_NAME=${BASE_NAME}"
echo "BASE_PATH=${BASE_PATH}"

# include
. ${BASE_PATH}/base-tag.sh
. ${BASE_PATH}/base-git-chglog.sh

set -e

# DEFAULT_BRANCH=$(git rev-parse --abbrev-ref origin/HEAD | awk -F'/' '{print $2}')

update-go-mod() {
  TAG=${1}

  PACKAGE_DIRS=$(find . -mindepth 2 -type f -name 'go.mod' -exec dirname {} \; | sed 's/^\.\///' | sort)

  for dir in $PACKAGE_DIRS; do
    printf "${dir}: go mod download && go mod tidy\n"
    (cd ./${dir} &&
      echo ">> PWD=${PWD}" &&
      go mod download &&
      go mod tidy &&
      echo "<< DONE")
  done

  read GROUP PROJECT <<<$(git remote get-url origin | awk -F'/' '{print $4,$5}')
  echo "PWD=[${PWD}], GROUP=[${GROUP}], PROJECT=[${PROJECT}]"
  for dir in $PACKAGE_DIRS; do
    sed --in-place \
      "s/${GROUP}\/${PROJECT}.git\([^ ]*\) v.*/${GROUP}\/${PROJECT}.git\1 ${TAG}/" "${dir}/go.mod"
    echo "change tag[${TAG}] in ${dir}/go.mod"
    (cd ./${dir} &&
      echo ">> PWD=${PWD}" &&
      printf "${dir}: go mod download && go mod tidy\n" &&
      go mod download &&
      go mod tidy &&
      echo "<< DONE")
  done
}

create_changelog() {
  TAG=${1}
  
  update-go-mod ${TAG}

  if [ ! -f ${BIN_DIR}/${APP_NAME} ]; then
    download-git-chglog
  fi

  echo ">> Generate CHANGELOG.md with ${TAG}"
  # https://github.com/git-chglog/git-chglog/blob/master/Makefile
  ${BIN_DIR}/${APP_NAME} --next-tag ${TAG} -o CHANGELOG.md
}

release-base() {
  BASE_BRANCH=${1}
  TAG=${2}

  check_tag ${TAG}

  echo ">> Checkout ${BASE_BRANCH}"
  git checkout ${BASE_BRANCH}

  create_changelog ${TAG}
}

release-base-branch() {
  BASE_BRANCH=${1}
  TAG=${2}

  validate_tag ${TAG}

  echo ">> Checkout ${BASE_BRANCH}"
  git checkout ${BASE_BRANCH}

  create_changelog ${TAG}
}

release-branch() {
  BASE_BRANCH=${1}
  TAG=${2}

  BRANCH_NAME=release/${TAG}

  git checkout -b ${BRANCH_NAME} ${BASE_BRANCH}
  git add -u
  git commit -m "chore: release/$TAG (release.sh)"
  git push origin ${BRANCH_NAME}
}

release-base-version() {
  BASE_BRANCH=${1}
  TAG=${2}

  validate_vtag ${TAG}
  
  echo ">> Checkout ${BASE_BRANCH}"
  git checkout ${BASE_BRANCH}

  create_changelog ${TAG}
}

release-version() {
  BASE_BRANCH=${1}
  TAG=${2}

  git checkout -b ${TAG} ${BASE_BRANCH}
  git add -u
  git commit -m "chore: version $TAG (release.sh)"
  git push origin ${BASE_BRANCH} ${TAG}
}

release-tag() {
  BASE_BRANCH=${1}
  TAG=${2}

  git add CHANGELOG.md
  git commit -m "chore: tag $TAG (release.sh)"
  git tag ${TAG}
  git push origin ${BASE_BRANCH} ${TAG}
}

case $1 in
branch)
  release-base-branch $2 $3
  release-branch $2 $3
  ;;
tag)
  release-base $2 $3
  release-tag $2 $3
  ;;
version)
  release-base-version $2 $3
  release-version $2 $3
  ;;
*)
  echo "Usage: ${0} [branch|tag] <base_branch> <TAG>"
  ;;
esac
