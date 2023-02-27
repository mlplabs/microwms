#!/bin/bash

BASE_DIR=/opt/microwms
DEPLOY_DIR=/home/mike/deploy_mwms


if [ -d $BASE_DIR ]; then
  echo "Stopped service..."
  systemctl stop microwms
  sleep 5
  echo "Copyng microwms files..."
  mv $BASE_DIR/mwms-daemon $BASE_DIR/mwms-daemon.bak
  mv $BASE_DIR/.env $BASE_DIR/.env.bak
  cp $DEPLOY_DIR/build/mwms-daemon $BASE_DIR/mwms-daemon
#  cp $DEPLOY_DIR/build/.env $BASE_DIR/.env
  echo "Starting service..."
  systemctl start microwms
else
  echo "New deploy started"
  mkdir -p $BASE_DIR
  echo "Copyng microwms files..."
  cp $DEPLOY_DIR/build/mwms-daemon $BASE_DIR/mwms-daemon
#  cp $DEPLOY_DIR/build/.env $BASE_DIR/.env
  echo "Registering and starting service"
  cp $DEPLOY_DIR/scripts/microwms.service /etc/systemd/system/microwms.service
  systemctl enable microwms.service
  systemctl start microwms
fi
