= Bug 相关
- 跳转后无法修改 subselected.y 的值, 由于我的 node 没有设置 fatherTree 信息以及坐标信息, 因此难以定位 ✅

= 计划
- 添加一个函数 `(t Tree) getTreePosi(tName string) point` 来获取 tree 的位置 , 需要判断调用者是不是 `root` tree ✅
  - 先解析 tName 为 levels, 找每一个 level 的 position 即可, 比如 hello/world/lala, 的位置:
    - 先找 root.Subtree 中匹配到 hello 的 position, 找到并存储 position 后, x=0, y++, 
    - 再找 hello.Subtree 中匹配到 world 的 position
    - 再找 world.Subtree 中匹配到 lala 的 position, 最后一层不需要 x=0, y++
    - 递归查找 levels 字符串数组, 递归结束条件是 `len(levels) == 0`
- 可以考虑, 为每一个 node 添加 default browser 的成员
