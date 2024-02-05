# Desc
Using bubbletea and libgloss to write a web favorates handler for command line usage.

Multiple ways to sort: label, folder
Concept image:

# Targets

    Have a nice UI, customization of UI components'position, maybe drag with mouse and store the new position parameter
    Add/delete web
    Description for each web
    Multiple name alias for each web page
    Scroll bar
    Open a web page through target web browser
    Icon setting for each web page
    Fuzzy finder
    Classification support
    Vim mode for moving around and search space, maybe customization for keybindings
    YMAL configuration file
    Custom sort way: folder, label
    Flip between description and link
    Copy the link to clipboard
    Edit link and description in real time
    Basic help info on the top, could be hidden
    Check whether the keybindings is conflict
    Configuration file should be placed at "~/.config/web-tree/"
    Every tree have a storage file? data 目录, 每创建一个 tree, 就添加一个文件
    Check configuration
    Move a node to another position
    Browser list show
    Root tree 的设计
    如何不继承给某一个
    通过逗号分隔成列表, 斜杠分隔成层级, 检查值的合法性, 如 "," 等不出现在 alias 中(在 Run 之前运行检查程序)
    检查是否为 url
    全局查找一个 node

Possible TUI for it:
- text input
- help
- key
- autocomplete
- composable-views
- http
- list-fancy
- progress-animated
- spinners
- table
- tabs
- tui-daemon-combo

1. Command line args
2. Draw the outer box
3. help, key menu
4. text input and autocomplete for search space
5. Add function for http testing

Alias could not be duplicated

## Cmd
Sub-command and function.

Without the sub-command, open the node with browser.

决定是否 persistent 以及类型, 长命令, 短命令, 默认值, help 信息. 哪些需要捆绑使用

根据 RootTree 这个变量来添加文件. 在 Add, Delete 这些操作之后, 需要更新 RootTree 等变量的信息.

默认创建一个 root.yaml 文件(临时文件), 包含其他 tree 信息, 通过读取这个文件来更新 RootTree 变量. 循环对 Root 之下的每一个 Tree 做写入操作以 Update. -Fi-

判断一个 Tree 是否改变. 比较内存中的大小是否可行. -Fi-

手动保存, 或退出时一并写入.

普通树名不能为 root -Fi-

添加 sub tree 可接收 list, map 等.

在 NewTree 和 NewNode 时对参数进行检验 -Fi-

判断一个 alias 以及 link 是否已经存在

去除重复的 hint

去除空字符串

edit 的问题, tree name 被修改导致找不到原文件

为每一个函数添加错误处理.

### 主要信息
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
### root

> Open tui

```sh
webtree
```

> Open with browser in conf file by alias

```sh
webtree --tree=<name> --node --alias=<alias>
```

> Open with browser in conf file by link

```sh
webtree --tree=<name> --node --link=<url>
```

> Target browser

```sh
webtree --tree=<name> --node --alias=<alias> --browser=firefox
```

### add
> Add a new web folder:

```sh
webtree add --tree=<name> 
```
> Add a new web folder with sub folder:

```sh
webtree add --tree=<name>/<subname>/<subsubname> 
```
> Add multiple tree

```sh
webtree add --tree=<name>/<subname>/<subsubname>,<another>
```
> Add a new node to a tree (at lease url and alias):

```sh
webtree add --tree=<name>/<subname> --node --link=<url1>,<url2> --alias=<alias1>,<alias2> --desc=<desc1>,<desc2> --label=<label1>,<label2>
```

### delete
> Delete a tree

```sh
webtree del --tree=<name>/<subname>
```
> Delete multiple trees

```sh
webtree del --tree=<name>/<subname>,<another>
```
> Delete a node by url

```sh
webtree del --tree=<name> --node --link=<url>
```
> Delete a node by alias

```sh
webtree del --tree=<name> --node --alias=<alias>
```

### edit
> Edit a tree name

```sh
webtree edit --tree=<name> --tname=<newname>
```

> Edit a node chosen by link

```sh
webtree edit --tree=<name> --node --link=<url> --nlink=<newlink> --nalias=<newalias> --ndesc=<newdesc> --nlabel=<newlabel>
```

> Edit a node chosen by alias

```sh
webtree edit --tree=<name> --node --alias=<alias> --nlink=<newlink> --nalias=<newalias> --ndesc=<newdesc> --nlabel=<newlabel>
```

### move
> Move a tree as a subtree of another tree

```sh
webtree move --tree=<name> <another tree name>
```

> Move a node to another tree

```sh
webtree move --tree=<name> --node --alias=<alias> <another tree name>
```

### list
> List all trees

```sh
webtree list tree
```
> List all nodes

```sh
webtree list node
```
> List all tags/labels

```sh
webtree list label
```
> List all styles

```sh
webtree list style
```
### show
> Show detail of a tree

```sh
webtree show --tree=<name>
```
> Show detail of a node

```sh
webtree show --tree=<name> --node --link=<url>
```
or:
```sh
webtree show --tree=<name> --node --alias=<alias>
```
> Show detail of a label

```sh
webtree show --label=<name>
```
# Usage
## Move around

    `?`: Help manual
    `<Tab>`: Moving between several space
    `j`: Move down
    `k`: Move up
    `h`: Move left
    `l`: Move right
    `o`: Open in web browser
    `a`: Add new web
    `d`: Delete a selected web
    `t`: Toggle from description and web link

## Work space

    Search, on the top of the UI
    Web tree, 




