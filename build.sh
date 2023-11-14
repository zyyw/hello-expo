GIT_COMMIT=`git rev-parse --short HEAD`

# docker build the hello-expo image
docker build \
--build-arg GIT_COMMIT=$GIT_COMMIT \
--build-arg GOPROXY=https://goproxy.cn,direct \
-t zigzag18/hello-expo:$GIT_COMMIT \
-f Dockerfile .

# push the hello-expo image
#docker push zigzag18/hello-expo:$GIT_COMMIT
