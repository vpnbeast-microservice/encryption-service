#!/bin/bash

function increment_minor_version() {
    local version_major=$(echo $1 | cut -d "." -f 1)
    local version_minor=$(echo $1 | cut -d "." -f 2)
    local version_patch=$(echo $1 | cut -d "." -f 3)
    version_patch=`expr ${version_patch} + 1`
    echo "${version_major}.${version_minor}.${version_patch}"
}

function set_version() {
    sed -i "s/${1}/${2}/g" version.properties
}

function set_chart_version() {
    sed -i "s/${1}/${2}/g" charts/${CHART_NAME}/Chart.yaml
    sed -i "s/${1}/${2}/g" charts/${CHART_NAME}/values.yaml
}

set -ex
CHART_NAME=encryption-service
CURRENT_VERSION=$(grep RELEASE_VERSION version.properties | cut -d "=" -f2)
RELEASE_VERSION=$(increment_minor_version ${CURRENT_VERSION})
set_version ${CURRENT_VERSION} ${RELEASE_VERSION}
set_chart_version ${CURRENT_VERSION} ${RELEASE_VERSION}
