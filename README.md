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
├── .bin/                   # ワークスペース管理コマンド (nushell)
│   ├── work                # 問題セットの作成 + zellij タブを開く
│   ├── run                 # ビルド & 実行
│   ├── stress              # ストレステスト
│   ├── clean               # ビルド成果物の削除
│   └── done                # zellij タブを閉じる
├── .helix/                 # Helix プロジェクトローカル設定 (自動適用)
│   ├── config.toml         # エディタ設定 (行番号・自動保存・キーバインド等)
│   └── languages.toml      # clangd / clang-format 設定
├── config/                 # ツール設定をまとめるディレクトリ
│   ├── forge.kdl           # zellij レイアウト (work コマンドで使用)
│   ├── zellij.kdl          # zellij 本体設定 (devbox.json で自動指定)
│   ├── completions.nu      # nushell タブ補完 (source して使う)
│   └── nu-config.nu        # nushell プロンプト・設定 (source して使う)
├── platforms/              # 問題ごとのフォルダを配置
│   └── <platform>/<contest>/<problem>/
│       ├── main.cpp        # 解答コード
│       ├── problem.md      # 問題文メモ
│       └── input*.txt      # サンプル入力 (任意)
├── template.cpp            # 新しい問題のひな形
├── compile_flags.txt       # clangd 用 (LSPで補完が効く)
├── .clang-format           # 整形ルール
└── devbox.json             # ツールチェイン + 環境変数の宣言
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

`platforms/<platform>/<contest>/<problem>/` が問題ごとに作られ、それぞれに以下が入ります:

- `main.cpp`  — `template.cpp` のコピー
- `problem.md` — 空の問題文メモ (適宜書き足し)

zellij で Editor / Terminal / Input の3ペイン構成のタブが開きます。

### 2. ビルドして実行

問題ディレクトリ (例: `A/`) に `cd` してから実行するのが基本の流れです。

```sh
cd A

# リリースビルド (-O2)
run main.cpp

# デバッグビルド (-g -DDEBUG, dbg() マクロが有効)
run main.cpp --debug
run main.cpp -d

# 入力ファイル (input1.txt など) を stdin に流す
run main.cpp --input input1.txt
run main.cpp -i input1.txt
```

コンテストディレクトリにいる場合は `run A/main.cpp` のように相対パスでも OK です。

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

## Helix 設定

`.helix/` ディレクトリにプロジェクトローカル設定があります。Helix を開くと自動で適用されます (ユーザー設定とマージ)。

| 設定 | 内容 |
|---|---|
| **相対行番号** | `5j` / `10k` のジャンプがしやすい |
| **ルーラー 100** | `.clang-format` の列幅と一致 |
| **保存時自動整形** | `clang-format --style=file` で保存するたびに整形 |
| **inlay hints** | 変数の型・引数名を薄く表示 |
| **自動保存** | フォーカスを外すと保存 |
| **Ctrl+S** | ノーマル・インサートモード両方で保存 |
| **Alt+. / Alt+,** | バッファ切り替え |

## Zellij 設定

`zellij.kdl` がプロジェクトローカルの Zellij 設定です。`.envrc` で `ZELLIJ_CONFIG_FILE` が自動的にこのファイルを指すため、devbox 環境に入るだけで有効になります。

| キー | 動作 |
|---|---|
| `Alt+h/l` | ペイン左右移動 (タブをまたいでも移動) |
| `Alt+k/j` | ペイン上下移動 |
| `Alt+H/L/K/J` | ペインリサイズ |
| `Alt+[` / `Alt+]` | タブ切り替え |
| `Alt+t` | 新規タブ |
| `Alt+z` | 現在のペインを全画面表示/解除 |
| `Alt+s` | スクロールモード (`d`/`u` で半ページ、`/` で検索) |

## Nushell 設定

`nu-config.nu` を config から読み込むと、プロンプトと操作性が強化されます。

```nu
# ~/.config/nushell/config.nu
source ($env.FORGE_ROOT | path join "config" "nu-config.nu")
source ($env.FORGE_ROOT | path join "config" "completions.nu")
```

| 機能 | 内容 |
|---|---|
| **プロンプト** | `[platform/contest]` のコンテキスト + git ブランチを表示 |
| **vi モード** | コマンドライン編集を vi キーバインドに (helix と統一) |
| **fuzzy 補完** | タイポしても候補が出る |
| **SQLite 履歴** | 全セッションで履歴を共有、10万件保持 |
| `fmt` | `clang-format --style=file -i` のエイリアス |
| `ls-problems` | 今のディレクトリの `.cpp` ファイルを整形表示 |
| `dbgrun` | `gdb --args` のエイリアス |

## ライセンス

個人用ワークスペースのためライセンス未設定。
