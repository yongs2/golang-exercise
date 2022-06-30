#!/bin/sh

DEFAULT_BRANCH=main

make_test_md_file() {
  CONTENT=${1}

  cat << EOF | tee ${CONTENT}.md
${CONTENT}
EOF

}

commit_base() {
  # ex) BRANCH=main,VERSION=v1.0.0
  BASE_BRANCH=${1}
  VERSION=${2}
  
  DASH_VER=$(echo ${VERSION} | sed 's/v//g' | sed 's/\./-/g')

  make_test_md_file "commit-${DASH_VER}"

  git add .
  COMMIT_TITLE="feat(core): Add commit-${DASH_VER}"
  COMMIT_BODY="some commit-${DASH_VER}"
  COMMIT_TAIL="issue ${DASH_VER}"
  git commit -m "${COMMIT_TITLE}" -m "${COMMIT_BODY}" -m "${COMMIT_TAIL}"

  git push origin ${BASE_BRANCH}
}

create_feature_branch() {
  # ex) BRANCH=main,VERSION=v1.0.0, TASK=task1
  BASE_BRANCH=${1}
  VERSION=${2}
  TASK=${3}

  DASH_VER=$(echo ${VERSION} | sed 's/v//g' | sed 's/\./-/g')
  FEATURE_BRANCH_NAME=feature/${VERSION}-${TASK}

  git checkout ${BASE_BRANCH}

  git branch ${FEATURE_BRANCH_NAME}
  git push origin ${FEATURE_BRANCH_NAME}
  git branch --set-upstream-to origin/${FEATURE_BRANCH_NAME}
  git checkout ${FEATURE_BRANCH_NAME}
  git branch -a

  make_test_md_file "feature-${DASH_VER}"

  git add feature-${DASH_VER}.md
  COMMIT_TITLE="feat(core): Add feature-${DASH_VER}"
  COMMIT_BODY="some feature-${DASH_VER}"
  COMMIT_TAIL="issue ${DASH_VER}"
  git commit -m "${COMMIT_TITLE}" -m "${COMMIT_BODY}" -m "${COMMIT_TAIL}"
  git push origin ${BASE_BRANCH}
}

merge_feature_branch() {
  # ex) BRANCH=main,VERSION=v1.0.0, TASK=task1
  BASE_BRANCH=${1}
  VERSION=${2}
  TASK=${3}

  FEATURE_BRANCH_NAME=feature/${VERSION}-${TASK}

  git checkout ${BASE_BRANCH}
  git branch -a
  git merge --no-edit --no-ff ${FEATURE_BRANCH_NAME}
  git push origin ${BASE_BRANCH}
}

delete_feature_branch() {
  # ex) BRANCH=main,VERSION=v1.0.0, TASK=task1
  BASE_BRANCH=${1}
  VERSION=${2}
  TASK=${3}

  FEATURE_BRANCH_NAME=feature/${VERSION}-${TASK}

  git branch -D ${FEATURE_BRANCH_NAME}
  git push origin :${FEATURE_BRANCH_NAME}
  git branch --unset-upstream
}

create_hotfix_release() {
  # ex) BRANCH=main, VERSION=v1.0.0, HOTFIX=hotfix-01
  BASE_BRANCH=${1}
  VERSION=${2}
  HOTFIX=${3}

  DASH_VER=$(echo ${VERSION} | sed 's/v//g' | sed 's/\./-/g')
  RELEASE_BRANCH_NAME=release/${VERSION}
  echo ">> RELEASE_BRANCH_NAME=${RELEASE_BRANCH_NAME}"

  git checkout ${RELEASE_BRANCH_NAME}
  git branch -a

  make_test_md_file "feature-release-${DASH_VER}-${HOTFIX}"

  git add feature-release-${DASH_VER}-${HOTFIX}.md
  COMMIT_TITLE="feat(core): Add feature-release-${DASH_VER}-${HOTFIX}"
  COMMIT_BODY="some feature-release-${DASH_VER}-${HOTFIX}"
  COMMIT_TAIL="issue ${DASH_VER}-${HOTFIX}"
  git commit -m "${COMMIT_TITLE}" -m "${COMMIT_BODY}" -m "${COMMIT_TAIL}"
  git push origin ${RELEASE_BRANCH_NAME}
}

change_release() {
  # ex) RELEASE_BRANCH_NAME=release/v1.0.0
  RELEASE_BRANCH_NAME=${1}

  git checkout ${RELEASE_BRANCH_NAME}
  git branch -a
}

merge_release_to() {
  # ex) RELEASE_BRANCH_NAME=release/v1.0.0, BASE_BRANCH=main
  RELEASE_BRANCH_NAME=${1}
  BASE_BRANCH=${2}

  echo ">> RELEASE_BRANCH_NAME=${RELEASE_BRANCH_NAME}"

  git checkout ${RELEASE_BRANCH_NAME}
  git branch -a
  
  git checkout main
  git branch -a
  git merge --no-edit --no-ff ${RELEASE_BRANCH_NAME}
  git push origin ${BASE_BRANCH}
}

case $1 in
commit_base)
  commit_base ${2} ${3}
  ;;
create_feature_branch)
  create_feature_branch ${2} ${3} ${4}
  ;;
merge_feature_branch)
  merge_feature_branch ${2} ${3} ${4}
  ;;
delete_feature_branch)
  delete_feature_branch ${2} ${3} ${4}
  ;;
create_hotfix_release)
  create_hotfix_release ${2} ${3} ${4}
  ;;
change_release)
  change_release ${2}
  ;;
merge_release_to)
  merge_release_to ${2} ${3}
  ;;
*)
  echo "Usage: ${0} [create_feature_branch|merge_feature_branch|delete_feature_branch|create_hotfix_release|merge_release_to] <VERSION> <TASK>"
  ;;
esac
