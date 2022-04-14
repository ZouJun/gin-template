#!/usr/bin/env bash

dateVersion=`date +"%Y%m%d%H%M%S"`
currentBranch=`git symbolic-ref --short -q HEAD`
export TAG=registry.cn-hangzhou.aliyuncs.com/630/marketing:${currentBranch}-${dateVersion}

echo dateVersion