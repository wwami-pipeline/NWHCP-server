#should not need to change anything
kubectl version --client
kubectl config set-cluster ${KUBE_CLUSTER} --server="${KUBE_SERVER}"
kubectl config set-credentials ${KUBE_USER} --token="${KUBE_USER_TOKEN}"
kubectl config set clusters.${KUBE_CLUSTER}.certificate-authority-data ${KUBE_CERTIFICATE_AUTHORITY_DATA}
kubectl config set-context ${KUBE_CLUSTER} --cluster=${KUBE_CLUSTER} --user=${KUBE_USER} && kubectl config use-context ${KUBE_CLUSTER}
kubectl config set-context --current --namespace=$(echo $(grep "namespace" ${KUBE_YML}) | awk '{print $2}')
kubectl version
sed -i "s~${GITLAB_IMAGE_LATEST}~${GITLAB_IMAGE_PIPELINE}~g" ${KUBE_YML} # write name of image to deploy
cat ${KUBE_YML} # show yml contents in log
kubectl apply -f ${KUBE_YML}
