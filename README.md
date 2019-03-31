Go で AWS S3 に画像アップロードするサンプルです。

# 必要環境

- Go 1.12.x
- Docker 18.09.3
- docker-compose 1.24.0-rc1
- aws-cli
- direnv

# セットアップ

```
git clone https://github.com/dyoshikawa/
go mod vendor
cp .envrc.example .envrc
direnv allow
```

## Localstack

```
docker-compose up -d
make localstack-setup
```

### aws-cli

```
aws configure --profile localstack
```

設定値はダミーの値で OK。

# 環境の切り替え

## .envrc

### Localstack S3 を使用

```
export STAGE=dev
```

### 本番の S3 を使用

```
export STAGE=prod
export S3_BUCKET=YOUR_S3_BUCKET_NAME
```

# アップロード開始

```
make run
```
