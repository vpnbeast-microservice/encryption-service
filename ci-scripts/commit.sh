#!/bin/bash

function git_global_settings() {
    git config --global user.name ${USERNAME}
    git config --global user.email ${EMAIL}
}

function git_commit_and_push() {
    git --no-pager diff
    git add --all
    git commit -am "[ci-skip] version ${RELEASE_VERSION}.RELEASE"
    git tag -a "v${RELEASE_VERSION}" -m "v${RELEASE_VERSION} tagged"
    git status
    git push --force --follow-tags ${PUSH_URL} HEAD:${BRANCH}
}

set -ex
PROJECT_NAME=encryption-service
BRANCH=master
USERNAME=vpnbeast-ci
EMAIL=info@thevpnbeast.com
GIT_ACCESS_TOKEN=$1
RELEASE_VERSION=$2
PUSH_URL=https://${USERNAME}:${GIT_ACCESS_TOKEN}@github.com/vpnbeast/${PROJECT_NAME}.git

git_global_settings
git_commit_and_push
