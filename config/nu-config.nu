# cpp-forge nushell settings.
# Source this in your nushell config:
#
#   source ($env.FORGE_ROOT | path join "config" "nu-config.nu")
#   source ($env.FORGE_ROOT | path join "config" "completions.nu")

# ---------- 履歴 ----------
$env.config.history = {
    max_size: 100_000
    sync_on_enter: true
    file_format: "sqlite"      # SQLite 形式で全セッション共有
    isolation: false
}

# ---------- 補完 ----------
$env.config.completions = {
    case_sensitive: false
    quick: true
    partial: true
    algorithm: "fuzzy"         # fuzzy で多少タイポしても補完できる
    use_ls_colors: true
}

# ---------- エディタ (コマンドライン編集) ----------
$env.config.edit_mode = "vi"  # vi キーバインドで helix と統一感を出す

# ---------- プロンプト ----------
# 競プロ作業中: [platform/contest] ユーザー @ カレントディレクトリ の形式で表示

def forge-context [] {
    let root = ($env.FORGE_ROOT? | default "")
    let cwd  = ($env.PWD)
    if $root != "" and ($cwd | str starts-with $root) {
        let rel = ($cwd | str replace $root "" | str trim -c "/")
        let parts = ($rel | split row "/")
        # platforms/<platform>/<contest> の中にいるときはコンテスト名を表示
        if ($parts | length) >= 3 and ($parts | first) == "platforms" {
            let platform = ($parts | get 1)
            let contest  = ($parts | get 2)
            $"(ansi cyan)[$platform/(ansi bold)($contest)(ansi reset)(ansi cyan)](ansi reset) "
        } else {
            ""
        }
    } else {
        ""
    }
}

def create_left_prompt [] {
    let ctx   = (forge-context)
    let dir   = ($env.PWD | str replace $env.HOME "~")
    let git   = (do { git branch --show-current } | complete)
    let branch = if $git.exit_code == 0 {
        $" (ansi yellow)($git.stdout | str trim)(ansi reset)"
    } else { "" }

    $"($ctx)(ansi green)($dir)(ansi reset)($branch) "
}

def create_right_prompt [] {
    let t = (date now | format date "%H:%M:%S")
    $"(ansi dark_gray)($t)(ansi reset)"
}

$env.PROMPT_COMMAND       = { create_left_prompt }
$env.PROMPT_COMMAND_RIGHT = { create_right_prompt }
$env.PROMPT_INDICATOR     = $"(ansi green)❯(ansi reset) "
$env.PROMPT_INDICATOR_VI_INSERT = $"(ansi green)❯(ansi reset) "
$env.PROMPT_INDICATOR_VI_NORMAL = $"(ansi yellow)❮(ansi reset) "

# ---------- エイリアス ----------

# 現在のコンテストディレクトリを素早く確認
alias ls-problems = ls *.cpp | select name size modified

# clang-format を手動適用
alias fmt = clang-format --style=file -i

# gdb で手軽にデバッグ
alias dbgrun = gdb --args

# ---------- 表示設定 ----------
$env.config.table = {
    mode: rounded
    index_mode: always
    show_empty: true
    trim: { methodology: wrapping, wrapping_try_keep_words: true, truncating_suffix: "..." }
}
$env.config.datetime_format = { normal: "%Y-%m-%d %H:%M:%S" }
$env.config.show_banner = false  # 起動メッセージを非表示

# ---------- 色 ----------
$env.config.color_config = {
    separator:               { fg: "dark_gray" }
    leading_trailing_space_bg: { attr: "n" }
    header:                  { fg: "green" attr: "b" }
    empty:                   { fg: "dark_gray" }
    int:                     { fg: "white" }
    float:                   { fg: "white" }
    string:                  { fg: "cyan" }
    bool:                    { fg: "light_blue" }
    filesize:                { fg: "cyan" }
    duration:                { fg: "white" }
    date:                    { fg: "purple" }
    range:                   { fg: "white" }
    record:                  { fg: "white" }
    list:                    { fg: "white" }
    block:                   { fg: "white" }
    nothing:                 { fg: "dark_gray" }
    binary:                  { fg: "dark_gray" }
    cell-path:               { fg: "white" }
    row_index:               { fg: "dark_gray" }
    record_field_separator:  { fg: "dark_gray" }
    list_stream:             { fg: "white" }
    hints:                   { fg: "dark_gray" attr: "i" }
}
