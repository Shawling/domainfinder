#!/bin/bash
# 进入文件目录
work_path=$(dirname $0)
cd ${work_path}
# build 子项目
echo Building synonyms...
cd synonyms
go build -o ../lib/synonyms
echo Building available...
cd ../available
go build -o ../lib/available
echo Building sprinkle...
cd ../sprinkle
cp ./transforms.txt ../lib/transforms.txt
go build -o ../lib/sprinkle
echo Building coolify...
cd ../coolify
go build -o ../lib/coolify
echo Building domainify...
cd ../domainify
cp ./tlds.txt ../lib/tlds.txt
go build -o ../lib/domainify
echo Building domainfinder...
# build main
cd ..
go build -o domainfinder
echo Done.