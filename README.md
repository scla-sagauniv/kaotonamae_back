# 顔と名前が一致しない＿バックエンド

### 環境構築手順
#### 前提
docker環境であること  
wsl環境であること

1. クローン  
   ```
   git clone git@github.com:scla-sagauniv/kaotonamae_back.git
   ```
2. 環境変数を設定  
   以下、環境変数の項目に従い`.env`ファイルを作成する
3. ルートディレクトリでビルド
   ```
   docker compose build
   ```
4. 起動
   ```
   docker compose up
   ```

### 環境変数
rootディレクトリに.env

```
DB_HOST=db
DB_PORT=3306
DB_USER=testuser
DB_PASSWORD=testpassword
DB_NAME=test_database
DB_ROOT_PASSWORD=rootpassword
```

### docker 再起動
- コンテナの削除
  ```
   docker compose down --volumes --remove-orphans
  ```
- キャッシュを消してビルド
  ```
  docker compose build --no-cache
  ```