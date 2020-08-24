#!/bin/bash
git pull

git add -A

# Commit changes.
msg="update `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"


# 推送到 github pages
# nusr.github.io 只能使用 master分支
# -f 强制推送
git push -f git@github.com:AverageJoeWang/go-blog-service.git master

