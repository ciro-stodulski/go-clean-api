

sed -i  "s/\$PROJECT_NAME/${APP_NAME}/"   deployment.yaml
sed -i  "s/\$APP_VERSION/${APP_VERSION}/"     deployment.yaml
sed -i  "s/\$NAMESPACE/${NAMESPACE}/"     deployment.yaml
sed -i  "s/\$DOCKER_IMAGE/${APP_IMAGE_REPO}\/${APP_IMAGE}:${APP_IMAGE_VERSION}/"     deployment.yaml
sed -i  "s/\$SERVICEPORT/${SERVICEPORT}/"     deployment.yaml
sed -i  "s/\$SERVICEPORT/${SERVICEPORT}/"  deployment.yaml

sed -i  "s/\$PROJECT_NAME/${APP_NAME}/"   hpa.yaml
sed -i  "s/\$NAMESPACE/${NAMESPACE}/"     hpa.yaml
sed -i  "s/\$APP_VERSION/${APP_VERSION}/"     hpa.yaml


sed -i  "s/\$PROJECT_NAME/${APP_NAME}/"   service.yaml
sed -i  "s/\$NAMESPACE/${NAMESPACE}/"     service.yaml
sed -i  "s/\$SERVICEPORT/${SERVICEPORT}/"  service.yaml