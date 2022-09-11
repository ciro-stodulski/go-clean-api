#!/bin/bash
function readJson {  
  UNAMESTR=`uname`
  if [[ "$UNAMESTR" == 'Linux' ]]; then
    SED_EXTENDED='-r'
  elif [[ "$UNAMESTR" == 'Darwin' ]]; then
    SED_EXTENDED='-E'
  fi; 

  VALUE=`grep -m 1 "\"${2}\"" ${1} | sed ${SED_EXTENDED} 's/^ *//;s/.*: *"//;s/",?//'`

  if [ ! "$VALUE" ]; then
    echo "Error: Cannot find \"${2}\" in ${1}" >&2;
    exit 1;
  else
    echo $VALUE ;
  fi; 
}

APP_VERSION=`readJson ../../package.json version` || exit 1;
NAMESPACE=`readJson ../../package.json namespace` || exit 1;

sed -i  "s/\$PROJECT_NAME/${APP_NAME}/"   deployment.yaml
sed -i  "s/\$APP_VERSION/${APP_VERSION}/"     deployment.yaml
sed -i  "s/\$NAMESPACE/${NAMESPACE}/"     deployment.yaml
sed -i  "s/\$DOCKER_IMAGE/${APP_IMAGE_REPO}\/${APP_IMAGE}:${APP_IMAGE_VERSION}/"     deployment.yaml
sed -i  "s/\$SERVICEPORT/${SERVICEPORT}/"     deployment.yaml
sed -i  "s/\$SERVICEPORT/${SERVICEPORT}/"  deployment.yaml

sed -i  "s/\$PROJECT_NAME/${APP_NAME}/"   hpa.yaml
sed -i  "s/\$NAMESPACE/${NAMESPACE}/"     hpa.yaml


sed -i  "s/\$PROJECT_NAME/${APP_NAME}/"   service.yaml
sed -i  "s/\$NAMESPACE/${NAMESPACE}/"     service.yaml
sed -i  "s/\$SERVICEPORT/${SERVICEPORT}/"  service.yaml