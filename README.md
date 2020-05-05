# convertTsvToDDL

クリップボードのタブ区切り文字列をテーブル定義 DDL に変換する。
エクセルからコピーして、テーブル定義を生成する用途を想定。
カラムは、次の内容になっていること。
```tsv
列名	データ型	サイズ	PK	-	-	-	-	NOT NULL	Default
```

例
```tsv
id	varchar	50	○					◯	
num_id	bigint unsigned	20						◯	
inv	tinyint	4						◯	0
inv_at	datetime								NULL
created	datetime							◯	1000-01-01 00:00:00
```


出力結果

```sql
CREATE TABLE IF NOT EXISTS TTTTT (
  `id` VARCHAR(50) NOT NULL ,
  `num_id` BIGINT(20) UNSIGNED NOT NULL ,
  `inv` TINYINT(4) NOT NULL DEFAULT 0,
  `inv_at` DATETIME  DEFAULT NULL,
  `created` DATETIME NOT NULL DEFAULT '1000-01-01 00:00:00',
  PRIMARY KEY(`id`)
);
```
