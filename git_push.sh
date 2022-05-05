#!/usr/bin/env bash

BRANCH=`git rev-parse --abbrev-ref HEAD`

echo "Pushing to branch ${BRANCH}"
echo "-----------------------------------------"
echo -n "Enter commit message: "
read -r COMMIT

git add .
git commit -m "$COMMIT"
git push origin ${BRANCH}