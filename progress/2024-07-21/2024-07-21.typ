= 08:21
- 先加入当前所在位置的标号, 在右下角, 示例 `2/10`, 2 为当前 `m.subselected.y` 的值, `10` 为总深度, 每次增加 subTree 时要更新总深度
  - 主要思路, 直接追加到 viewport 的后面 `indexView()` , Box 把两者框起来
  - 开始位置为: m.viewport.Width-len(index)
  - 如何计算出总深度: 当前 y 加上当前 tree 的 subtree 层数, 最后一层有 node, 则还要加 1
  - 设计一个函数: `(t Tree) GetTreeDepth() int`
  - 如果当前选中的是 tree, 则用 `GetTreeDepth` 更新 index; 如果选中的是 node, 则 index 更新为 y 值
  ✅
