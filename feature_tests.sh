#!/bin/bash

go build osg.go

#===================================================
# Settings
#===================================================
export ghub_endpoint='https://codeload.github.com/AnthonyHewins'

# Remote vars
export branch='master'
export test_repo='one-time-pad-socket'

# Remote vars compiled for local usage
export filename=$test_repo-$branch

export master_file=fixtures/local-dir.test.out
#===================================================


echo Running master
./osg -d fixtures/$filename > $master_file

echo Running local tar
export tardir='fixtures/local-tar.test.out'
./osg -t fixtures/$filename.tar.gz > $tardir
echo Running diff between master and tardir
diff $master_file $tardir

echo Running local zip
export zipdir='fixtures/local-zip.test.out'
./osg -z fixtures/$filename.zip > $zipdir
echo Comparing local zip to master
diff $master_file $zipdir
