= 14:01
== 基本安排
- 添加 Delete 功能
- 考虑 Delete 对 preSelectedTree 的影响
- 样式的添加
- Edit 对样式的修改
- 考虑 `display` 模式下批量选中的添加, `visual` 模式

= Delete 对 preSelectedTree 的影响
== 针对将选中的元素删除
若选中的是 node, 删除后:
- 如果右侧还有元素, 则 `subselected.x` 和 `subselected.y` 都不变
- 如果右侧没有元素, 但左边有元素, 则 `subselected.x` 减一, `subselected.y` 都不变
- 如果右侧没有元素, 左边也没有元素, 则 `subselected = preSelectedTree[len(preSelectedTree)-1]` , `preSelectedTree = preSelectedTree[:len(preSelectedTree)-1]`

若选中的是 tree, 删除后:
- 如果右侧还有元素, 则 `subselected.x` 和 `subselected.y` 都不变
- 如果右侧没有元素, 但左边有元素, 则 `subselected.x` 减一, `subselected.y` 都不变
- 如果右侧没有元素, 左边也没有元素, 则 `subselected = preSelectedTree[len(preSelectedTree)-1]` , `preSelectedTree = preSelectedTree[:len(preSelectedTree)-1]`
和 node 行为似乎一致.

= 具体的 Delete 功能
在 `display` 模式下, 按下 `d`, 删除光标下选择的元素, 即 `x == m.subselected.x && y == m.subselected.y` 的元素.

需要出现一个弹窗, 并询问是否删除. 有两个按钮, Yes 和 No. 弹窗出现在窗口正中.




