# Readme in test/module

- git commit 메시지 template 를 이용하여 branch 전략을 관리하는 예제
- go module 관리 방식으로 버전 변경시 v2 branch 를 별도로 최신 상태로 유지

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

# Example.md for test-module

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
scripts/release.sh tag main v1.0.1
```

### 3.1 v2 생성

```sh
scripts/release.sh version main v2
```

### 3.2 feature 2-0-0 작업

```sh
scripts/base-git-cmd.sh create_feature_branch v2 v2.0.0 task-03
scripts/base-git-cmd.sh merge_feature_branch v2 v2.0.0 task-03
scripts/base-git-cmd.sh delete_feature_branch v2 v2.0.0 task-03
scripts/release.sh tag v2 v2.0.0
```

### 4.1 v3 생성

```sh
scripts/release.sh version v2 v3
```

### 4.2 feature 3-0-0 작업

```sh
scripts/base-git-cmd.sh create_feature_branch v3 v3.0.0 task-04
scripts/base-git-cmd.sh merge_feature_branch v3 v3.0.0 task-04
scripts/base-git-cmd.sh delete_feature_branch v3 v3.0.0 task-04
scripts/release.sh tag v3 v3.0.0
```
