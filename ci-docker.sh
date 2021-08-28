# should not need to change anything
docker --version
docker login -u gitlab-ci-token -p ${CI_BUILD_TOKEN} ${GITLAB_REGISTRY}
docker build -t ${GITLAB_IMAGE_PIPELINE} -t ${GITLAB_IMAGE_LATEST} -f ${DOCKERFILE} .
docker push ${GITLAB_IMAGE_PIPELINE}
docker push ${GITLAB_IMAGE_LATEST}
