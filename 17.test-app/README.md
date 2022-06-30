# Readme in test/app

- git commit 메시지 template 를 이용하여 branch 전략을 관리하는 예제
- release 별로 별도의 branch 를 관리
- 항상 main branch 를 최신 상태로 유지

## Initialize

- 프로젝트 생성
- git commit template 설정

```sh
git config --local commit.template .gitmessage.txt
```

## Post-Step

- git-chglog 관련 설정 파일과 배포 관련 스크립트를 추가

```sh
cp -Rp ${BASE_DIR}/.chglog/ ${PRJECT_APP}/.chglog
cp -Rp ${BASE_DIR}/scripts/ ${PRJECT_APP}/scripts
```

- 커밋

```sh
git add .
git commit -m "feat(core): Add base script files" -m "some base script files" -m "issue 1";
git push -u origin main
```

## Example.md for test-app

### 1.2 초기 작업

```sh
scripts/base-git-cmd.sh commit_base main v0.0.1
```

### 1.3 tag v0.0.1 생성

```sh
scripts/release.sh tag main v0.0.1
```

### 2.1 feature 1-0-0 작업

```sh
scripts/base-git-cmd.sh create_feature_branch main v1.0.0 task-01
scripts/base-git-cmd.sh merge_feature_branch main v1.0.0 task-01
scripts/base-git-cmd.sh delete_feature_branch main v1.0.0 task-01
scripts/release.sh tag main v1.0.0
```

### 2.2 feature 1-0-1 작업

```sh
scripts/base-git-cmd.sh create_feature_branch main v1.0.1 task-02
scripts/base-git-cmd.sh merge_feature_branch main v1.0.1 task-02
scripts/base-git-cmd.sh delete_feature_branch main v1.0.1 task-02
```

### 3.1 release/v1.0.1 생성

```sh
scripts/release.sh branch main v1.0.1
```

### 3.2 release/v1.0.1 에서 hotfix 추가

```sh
scripts/base-git-cmd.sh create_hotfix_release main v1.0.1 hotfix-01
```

### 4.1 feature 1-0-2 작업

```sh
scripts/base-git-cmd.sh create_feature_branch main v1.0.2 task-03
scripts/base-git-cmd.sh merge_feature_branch main v1.0.2 task-03
scripts/base-git-cmd.sh delete_feature_branch main v1.0.2 task-03
```

### 5.1 release/v1.0.1 검증 완료

```sh
scripts/base-git-cmd.sh change_release release/v1.0.1
scripts/release.sh tag release/v1.0.1 v1.0.1
scripts/base-git-cmd.sh merge_release_to release/v1.0.1 main
```

### 6.1 release/v1.0.2 생성

```sh
scripts/release.sh branch main v1.0.2
```

### 6.2 release/v1.0.2 에서 hotfix 추가

```sh
scripts/base-git-cmd.sh create_hotfix_release main v1.0.2 hotfix-02
```

### 7.1 release/v1.0.2 검증 완료

```sh
scripts/base-git-cmd.sh change_release release/v1.0.2
scripts/release.sh tag release/v1.0.2 v1.0.2
scripts/base-git-cmd.sh merge_release_to release/v1.0.2 main
```
