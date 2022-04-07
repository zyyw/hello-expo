GIT_COMMIT=`git rev-parse HEAD`

COUNT=1000
if [ -n "$1" ] && [ $1 -gt 0 ]; then
  COUNT=$1
fi

# docker build the controller-manager image for tams-extension
docker build \
--build-arg GIT_COMMIT=$GIT_COMMIT \
--build-arg GOPROXY=http://10.202.250.221:3000 \
--build-arg COUNT=$COUNT \
-t zigzag18/hello-expo:$GIT_COMMIT \
-f Dockerfile .

# push the hello-expo image
#docker push zigzag18/hello-expo:$GIT_COMMIT
