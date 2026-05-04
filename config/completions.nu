# cpp-forge nushell completions.
#
# Add this to your nushell config (e.g. ~/.config/nushell/config.nu):
#     source ($env.FORGE_ROOT | path join "config" "completions.nu")
#
# After sourcing, `work`, `run`, and `stress` get tab completion that knows
# about platforms, contests, and .cpp files in the current directory.

def "nu-complete forge-platforms" [] {
    let root = ($env.FORGE_ROOT? | default ".")
    let dir = ([$root, "platforms"] | path join)
    if not ($dir | path exists) { return [] }
    ls $dir | where type == dir | get name | each { |it| $it | path basename }
}

def "nu-complete forge-contests" [context: string] {
    let root = ($env.FORGE_ROOT? | default ".")
    let parts = ($context | split row " " | where { |it| ($it | str length) > 0 })
    if ($parts | length) < 2 { return [] }
    let platform = ($parts | get 1)
    let dir = ([$root, "platforms", $platform] | path join)
    if not ($dir | path exists) { return [] }
    ls $dir | where type == dir | get name | each { |it| $it | path basename }
}

def "nu-complete forge-cpp" [] {
    glob *.cpp | each { |it| $it | path basename }
}

# Wrapper for `work` with platform/contest completion.
export def --wrapped work [
    platform: string@"nu-complete forge-platforms"
    contest?: string@"nu-complete forge-contests"
    ...problems: string
] {
    ^work $platform $contest ...$problems
}

# Wrapper for `run` with .cpp file completion.
export def --wrapped run [
    file: string@"nu-complete forge-cpp"
    --debug (-d)
    --input (-i): string
] {
    if $debug and $input != null {
        ^run $file --debug --input $input
    } else if $debug {
        ^run $file --debug
    } else if $input != null {
        ^run $file --input $input
    } else {
        ^run $file
    }
}

# Wrapper for `stress` with .cpp file completion.
export def --wrapped stress [
    main_src: string@"nu-complete forge-cpp"
    naive_src: string@"nu-complete forge-cpp"
    gen_src: string@"nu-complete forge-cpp"
    --count (-c): int = 100
] {
    ^stress $main_src $naive_src $gen_src --count $count
}
