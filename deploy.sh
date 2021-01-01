#!/bin/bash

# -e: Exit immediately if a command exits with a non-zero status, -v: Print shell input lines as they are read.
set -evx

# Constants
PROJECT_ID="PROJECT_ID"
HOSTNAME="HOSTNAME"
SOURCE_IMAGE="be_dashboard_bert_faqclass"
REMOTE_IMAGE=${HOSTNAME}/${PROJECT_ID}/${SOURCE_IMAGE}

COMPONENT="be-dashboard-bert-faqclass"
SERVICE_ACCOUNT="${COMPONENT}@${PROJECT_ID}.iam.gserviceaccount.com"

REGION=europe-west1

# Deploy
gcloud run deploy \
    ${COMPONENT} \
    --project ${PROJECT_ID} \
    --platform managed \
    --region  ${REGION}\
    --image ${REMOTE_IMAGE} \
    --service-account ${SERVICE_ACCOUNT} \
    --memory 200M \
    --port 8081