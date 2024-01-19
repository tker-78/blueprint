# websocketを使ったチャットアプリケーション


## デプロイ

blueprint-410621

gcloud config set project blueprint-410621

gcloud services enable compute.googleapis.com \
  run.googleapis.com artifactregistry.googleapis.com \
  cloudbuild.googleapis.com servicenetworking.googleapis.com


$ gcloud iam service-accounts create blueprint-service-account \
  --display-name="Blueprint Service Account" 


環境変数はデプロイ時に指定してあげるか、
コンソールから登録する。
```
gcloud run deploy blueprint \
  --region=us-central1 \
  --source=. \
  --service-account="blueprint-service-account@blueprint-410621.iam.gserviceaccount.com" \
  --allow-unauthenticated \
  --set-env-vars security_key="" \
  --set-env-vars client_id="" \
  --set-env-vars client_secret="" \
  --set-env-vars url="https://blueprint-e6cexn42ya-uc.a.run.app/auth/callback/google" 
```


https://blueprint-e6cexn42ya-uc.a.run.app

websocketの接続には、`wss://`スキームを指定する。(`ws://`では不安全のため接続できない。)