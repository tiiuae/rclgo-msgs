#!/bin/bash

if [ $# != 2 ]; then
    echo "usage: $0 <ros-package-path> <dest-dir>"
    exit 1
fi

ros_package_path="$(realpath "$1")"
dest_dir="$(realpath "$2")"

export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"

ros_package="$(basename "$ros_package_path")"

go install github.com/tiiuae/rclgo/cmd/rclgo-gen@latest || exit 1

cd "$ros_package_path"

git config --global user.name "$(git log --format='%an' HEAD^! || exit 1)" || exit 1
git config --global user.email "$(git log --format='%ae' HEAD^! || exit 1)" || exit 1

cd "$dest_dir"

retry_delay=3s
for i in {1..5}; do
    if [ $i -gt 1 ]; then
        echo "Trying again after $retry_delay"
        sleep "$retry_delay"
    fi
    git reset --hard "origin/$(git branch --show-current)"
    git clean -dxf
    git pull -X theirs || continue
    rm -rf "$ros_package"
    rclgo-gen generate --root-path "$ros_package_path" || continue
    git add -A || continue
    commit_output="$(git commit -m "Update bindings for $ros_package")"
    ret_code=$?
    echo "$commit_output"
    if [[ $ret_code != 0 ]]; then
        if grep -iq "nothing to commit, working tree clean" <<< "$commit_output"; then
            exit 0
        else
            continue
        fi
    fi
    git push || continue
    exit 0
done
exit 1
