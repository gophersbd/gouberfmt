#!/usr/bin/env bash

#Get the highest tag number
VERSION=`git describe --abbrev=0 --tags`
VERSION=${VERSION:-'0.0.0'}

echo 'Last version ' ${VERSION}

#Get number parts
MAJOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
MINOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
PATCH="${VERSION%%.*}"; VERSION="${VERSION#*.}"

#Increase version
PATCH=$((PATCH+1))

#Get current hash and see if it already has a tag
GIT_COMMIT=`git rev-parse HEAD`
NEEDS_TAG=`git describe --contains ${GIT_COMMIT}`

#Create new tag
NEW_TAG="$MAJOR.$MINOR.$PATCH"
if [[ $1 != '' ]]; then
  NEW_TAG="$1"
fi

echo $NEW_TAG

#Only tag if no tag already (would be better if the git describe command above could have a silent option)
if [[ -z "$NEEDS_TAG" ]]; then
    echo "Updating to $NEW_TAG"
    echo "Tagged with $NEW_TAG (Ignoring fatal:cannot describe - this means commit is untagged) "
    git tag ${NEW_TAG}
else
    echo "Already a tag on this commit, please use" ${VERSION}
fi
