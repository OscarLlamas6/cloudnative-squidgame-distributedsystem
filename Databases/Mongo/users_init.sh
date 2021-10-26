#!/bin/bash
set -e

# dbUser is the userName used from applicatoin code to interact with databases and dbPwd is the password for this user.
# MONGO_INITDB_ROOT_USERNAME & MONGO_INITDB_ROOT_PASSWORD is the config for db admin.
# admin user is expected to be already created when this script executes. We use it here to authenticate as admin to create
# dbUser and databases.

echo ">>>>>>> trying to create database and users"
if [ -n "${MONGO_INITDB_ROOT_USERNAME:-}" ] && [ -n "${MONGO_INITDB_ROOT_PASSWORD:-}" ] && [ -n "${dbUser:-}" ] && [ -n "${dbPwd:-}" ]; then
mongo -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD<<EOF
db=db.getSiblingDB('tomi');
db=db.getSiblingDB('imdb');
use tomi;
db.createUser({
  user:  '$dbUser',
  pwd: '$dbPwd',
  roles: [{
    role: 'readWrite',
    db: 'tomi'
  }]
});
use imdb;
db.createUser({
  user:  '$dbUser',
  pwd: '$dbPwd',
  roles: [{
    role: 'readWrite',
    db: 'imdb'
  }]
});
EOF
else
    echo "MONGO_INITDB_ROOT_USERNAME,MONGO_INITDB_ROOT_PASSWORD,dbUser and dbPwd must be provided. Some of these are missing, hence exiting database and user creatioin"
    exit 403
fi