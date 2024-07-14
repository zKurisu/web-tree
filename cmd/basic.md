# Develop info
root 的 flag 有:
- `--version`, `-v`, "0.01"
- `--help`, `-h`, 暂时不用管
- `--tree`, `-t`, persistent, string, "root", "target web tree"
- `--node`, `-n`, persistent, bool, false, "target node"
- `--alias`, `-a`, persistent, string, "", "Alias of target node"
- `--link`, `-l`, persistent, string, "", "Urls of target node"
- `--browser`, `-b`, local, string, ""(从 conf 获取), "Open url with target browser"

root 的 sub-command 有:
- add
- del
- edit
- move
- list
- show

add 除继承以外的 flag:
- `--desc`, `-d`, local, string, "", "Descriptions of target node"
- `--label`, `none`, local, string, "", "Labels of target node"

del 无额外 flag.

edit 除继承以外的 flag:
- `--tname`, `-N`, local, string, "", "New tree name"
- `--nlink`, `-L`, local, string, "", "New links for target node"
- `--nalias`, `-A`, local, string, "", "New alias for target node"
- `--ndesc`, `-D`, local, string, "", "New descriptions for target node"
- `--nlabel`, `none`, local, string, "", "New labels for target node"

move 无额外 flag.

list 有 subcommand.

show 无额外 flag.

