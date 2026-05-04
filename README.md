# cpp-forge

競技プログラミング用のワークスペース。devbox + zellij + nushell で C++ の問題ごとのファイル管理・ビルド・実行をひとまとめにします。

## 必要なもの

- [devbox](https://www.jetify.com/devbox)
- [direnv](https://direnv.net/) (任意、自動で環境を有効化できる)
- [zellij](https://zellij.dev/) (タブ機能を使うため)
- [nushell](https://www.nushell.sh/) (`.bin/` 内のスクリプト用)

## セットアップ

```sh
# 環境に入る (direnv を使っていれば cd した時点で自動)
devbox shell
```

`devbox.json` で `gcc`, `clang-tools`, `gdb` が入ります。`PATH` には `.bin/` が、`CPLUS_INCLUDE_PATH` にはリポジトリルートが追加されるので、`#include <lib/...>` のような書き方が可能です。

## ディレクトリ構成

```
.
├── .bin/             # ワークスペース管理コマンド (nushell)
│   ├── work          # 問題セットの作成 + zellij タブを開く
│   ├── run           # ビルド & 実行
│   ├── stress        # ストレステスト
│   ├── clean         # ビルド成果物の削除
│   └── done          # zellij タブを閉じる
├── platforms/        # 問題ごとのソースを配置 (例: platforms/aoj/itp1_1/A.cpp)
├── template.cpp      # 新しい問題のひな形
├── compile_flags.txt # clangd 用 (LSPで補完が効く)
├── .clang-format     # 整形ルール
├── devbox.json       # ツールチェインの宣言
└── forge.kdl         # zellij レイアウト
```

## 基本的な使い方

### 1. 問題セットを作る

```sh
# AtCoder の abc300 で A〜F を作る (デフォルト)
work atcoder abc300

# 問題を明示
work atcoder abc300 A B C

# 既存のコンテストならタブだけ開く
work atcoder abc300
```

`platforms/<platform>/<contest>/` 以下に `template.cpp` がコピーされ、zellij で Editor / Terminal / Input の3ペイン構成のタブが開きます。

### 2. ビルドして実行

```sh
# リリースビルド (-O2)
run A.cpp

# デバッグビルド (-g -DDEBUG, dbg() マクロが有効)
run A.cpp --debug
run A.cpp -d

# 入力ファイルから実行
run A.cpp --input sample.txt
run A.cpp -i sample.txt
```

### 3. ストレステスト (WAデバッグ用)

```sh
# main.cpp の出力と naive.cpp (愚直解) の出力を比較
# gen.cpp はテストケース生成器
stress main.cpp naive.cpp gen.cpp

# 試行回数を指定 (デフォルト 100)
stress main.cpp naive.cpp gen.cpp --count 500
```

差分が出た時点で停止し、入力ケースを表示します。

### 4. 後片付け

```sh
clean   # 各 cpp と並んでいるバイナリを一括削除
done    # 現在の zellij タブを閉じる
```

## テンプレート (`template.cpp`)

定型のインクルード、エイリアス (`ll`, `pii`, …)、便利マクロ (`rep`, `all`, `chmin/chmax`)、デバッグ用 `dbg()` を備えています。`run -d` でビルドしたときだけ `dbg(x)` が標準エラーに出力されます。

```cpp
vector<int> v = {1, 2, 3};
dbg(v);   // -d ビルド時のみ "v = [1, 2, 3]" が cerr に出る
```

`vector` / `pair` / `map` 等もそのまま `dbg()` で出せます。

## エディタ統合

- **clangd**: `compile_flags.txt` を読むので追加設定不要で補完・診断が効きます
- **整形**: `.clang-format` (Google ベース、インデント 2、列幅 100)

## シェル補完 (任意)

`completions.nu` を nushell の config から読み込むと、`work` / `run` / `stress` で Tab 補完が効くようになります。

```nu
# ~/.config/nushell/config.nu
source ($env.FORGE_ROOT | path join "completions.nu")
```

- `work <Tab>` → `platforms/` 直下のディレクトリ (atcoder, aoj, ...)
- `work atcoder <Tab>` → そのプラットフォーム配下のコンテスト一覧
- `run <Tab>` / `stress <Tab>` → カレントディレクトリの `*.cpp`

## ライセンス

個人用ワークスペースのためライセンス未設定。
