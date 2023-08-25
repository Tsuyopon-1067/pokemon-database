mariadbをインストール
```
sudo apt install mariadb-server
```
mariadb初期設定
```
sudo mysql_secure_installation
```

mysqlログイン
```
sudo mysql -u root -p
```

データベース「Pokemon」を作成
```
CREATE DATABASE Pokemon;
```

--全てのSQL文を発行できるユーザ「uuu」の作成
```
GRANT ALL PRIVILEGES ON Pokemon.* TO uuu@localhost IDENTIFIED BY 'hogehoge';
```
ユーザ「uuu」でログイン
```
mysql -D Pokemon -u uuu -p
```
テーブル「Hgss」を作成
```
CREATE TABLE Hgss(
    id INT NOT NULL,
    name VARCHAR(30) NOT NULL ,
    HP INT NOT NULL,
    Attack INT NOT NULL,
    Defense INT NOT NULL,
    Special Attack INT NOT NULL,
    Special Defense INT NOT NULL,
    Speed INT NOT NULL
);
```
collyのインストール
```
go get github.com/gocolly/colly/v2
```