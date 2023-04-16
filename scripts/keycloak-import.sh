#!/bin/bash
docker compose exec grpc-keycloak /opt/keycloak/bin/kc.sh import \
  --dir /work \
  --override false
