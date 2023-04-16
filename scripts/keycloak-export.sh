#!/bin/bash
rm ./keycloak/config/*
docker compose exec grpc-keycloak /opt/keycloak/bin/kc.sh export \
  --dir /work \
  --realm test
