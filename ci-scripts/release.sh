#!/bin/bash

function git_checkout() {
    git checkout ${BRANCH}
    git pull origin ${BRANCH}
}

function git_global_settings() {
    git config user.userName ${USERNAME}
    git config user.email ${EMAIL}
}

function git_commit_and_push() {
    git --no-pager diff
    git add --all
    git commit -am "[ci-skip] version ${RELEASE_VERSION}.RELEASE"
    git tag -a "${RELEASE_VERSION}.RELEASE" -m "v${RELEASE_VERSION} tagged"
    git status
    git push --follow-tags http://${USERNAME}:${GIT_ACCESS_TOKEN}@gitlab.com/${GROUP_NAME}/${PROJECT_NAME}.git HEAD:${BRANCH}
}

function increment_minor_version() {
    local version_major=$(echo $1 | cut -d "." -f 1)
    local version_patch=$(echo $1 | cut -d "." -f 2)
    local version_minor=$(echo $1 | cut -d "." -f 3)
    version_minor=`expr ${version_minor} + 1`
    echo "${version_major}.${version_patch}.${version_minor}"
}

function set_version() {
    sed -i "s/${1}/${2}/g" version.properties
}

function docker_login() {
    docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}
}

function docker_build_image() {
    echo "building docker image with maven..."
    docker build -t ${DOCKER_USERNAME}/${PROJECT_NAME}:latest .
}

function docker_tag_image() {
    echo "tagging builded image with the version info..."
    docker tag ${DOCKER_USERNAME}/${PROJECT_NAME}:latest ${DOCKER_USERNAME}/${PROJECT_NAME}:${RELEASE_VERSION}
}

function docker_push_image() {
    echo "pushing with version info..."
    docker push ${DOCKER_USERNAME}/${PROJECT_NAME}:${RELEASE_VERSION}
    echo "pushing with latest tag..."
    docker push ${DOCKER_USERNAME}/${PROJECT_NAME}:latest
}

function set_chart_version() {
    echo "previous docker image is ${DOCKER_USERNAME}/${PROJECT}:${CURRENT_VERSION}, changing it as ${DOCKER_USERNAME}/${PROJECT}:${RELEASE_VERSION}!"
    sed -i "s/${DOCKER_USERNAME}\/${PROJECT_NAME}:${CURRENT_VERSION}/${DOCKER_USERNAME}\/${PROJECT_NAME}:${RELEASE_VERSION}/g" charts/${PROJECT_NAME}/values.yaml
    sed -i "s/${CURRENT_VERSION}/${RELEASE_VERSION}/g" charts/${PROJECT_NAME}/Chart.yaml
}

set -ex
USERNAME=bilalcaliskan
EMAIL=bilalcaliskan@protonmail.com
PROJECT_NAME=encryption-service
GROUP_NAME="vpnbeast/backend"
CURRENT_VERSION=`grep RELEASE_VERSION version.properties | cut -d "=" -f2`
RELEASE_VERSION=$(increment_minor_version ${CURRENT_VERSION})
CLONE_URL=git@gitlab.com:vpnbeast/${GROUP_NAME}/${PROJECT_NAME}.git
PUSH_URL=http://${USERNAME}:${GIT_ACCESS_TOKEN}@gitlab.com/vpnbeast/${GROUP_NAME}/${PROJECT_NAME}.git
BRANCH=$1

docker_login
docker_build_image
docker_tag_image
docker_push_image

set_version $CURRENT_VERSION $RELEASE_VERSION
set_chart_version
git_global_settings
git_commit_and_push
