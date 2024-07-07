 edit 模式以及 testarea
首先, 按下 `e` 之后,  先经过 `Update` 函数, 设置 `edit` 模式, 之后可以用 `updateTextarea()` 来设置相关参数, 接下来再在 `getTreeView()` 函数中设置显示.

只能从 `display` 模式进入 `edit` 模式.

== 给 Textarea 绑定键位
几个细节, 如果是 `tree`, 则设置为无法换行.

一个 bug: 按下 `e` 进入 `edit` 模式后, 会输出 `e` 字符.

如何确认输入完成? 按下 `esc` 是直接退出, 按下 `ctrl+s` 是保存并退出 (`Blur`).

== `edit` 模式涉及功能
- 修改 tree name ✅
- 修改 node info ✅
- 修改 style

`case m.subSelected.content.(type) {}}` 的使用以判断当前节点是 tree 还是 node. ✅

对于 node, 需要一个函数判断 Alias, Link, Tags 这些哪一个最长, 用来作为 TextArea 的宽度, 先默认使用 Links. ✅ 

== 配置文件的更新
先获取 `TextArea` 中的内容 `Value()`, 将字符串解析为 `Tree` 和 `Node` 需要的成员值, 赋值后更新文件.

对于 node:
- 第一行是 Link, 用 `split` 拆分后取 `[1:]`
- 第二行是 Alias, 用 `split` 拆分后取 `[1:]`
- 第三行是 Desc, 用 `split` 拆分后取 `[1:]`
- 第四行是 Label, 用 `split` 拆分后取 `[1:]`

对于 tree:
- 需要更改层级, 比如原来是 `hello/baba/lala` 改成了 `hello/papa/lala`, 此时需要先在 `hello/papa` 下添加子树 `lala`, 再从 `hello/baba` 上删除 `lala` 子树, 

具体, 用 `AppendSubTree` 添加子树, 需要先找到父树的地址, 用 `GetRootTree()` 以及 `DeepFindSubTree`.

再用 `DelSubTree` 删除子树, 同样要找到父树.

`nextFatherTree` 若不存在则需要创建.

= help 信息相关
似乎现在会把 help 信息显示在光标位置.

= ESC 键位问题
在 `display` 模式下按 `esc` 会显示出 `search` 模式的输入框.

= Bug 相关
- 在 tree 没有 subtree 和 node 的时候 y 轴仍会增加 ✅
- 第二个 subtree 没法向下查看, 因为 x 清 0 了, 似乎是坐标系统有问题, 需要第三个参数记录上一行的横坐标, 或者说需要记录层级数
- 添加了新的 tab 之后, 没有更新布局

需要一个链表数据结构, 存储上一个 tree 的坐标
