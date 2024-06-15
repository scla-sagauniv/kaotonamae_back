# あなたを知る魔法（Dich zu Kennen）_backend

## Dich zu Kennen 
- 読み： ディヒ ツゥ ケネン
- 言語: ドイツ語
- 意味: あなたを知る
- 名づけ: 他者の名前だけでなく、趣味や好きな色など、さまざま側面でその人のことを認識する、"あなたを知る"ためのアプリにしたいという願いを込めた。読みは、"葬送のフリーレン"に登場する魔法の命名規則に従いドイツ語にしている。
  
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

### 注意書き
   `docker compose up`で立ち上げると、  
   dbより先にechoが立ち上がって接続がうまくいかないときがある  

   そういう時は  `docker compose up -d db`で  
   データベースをバックグラウンドで動かしてから`docker compose up`するとよい
