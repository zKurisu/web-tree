= 16:18
- 把 del mode 改为 `delete` flag ✅
- 可以用 `add` 添加 `tab`
- 跳转到新添加的 tab:
  - 用 `sort` 排序, 并找到新添加 tab 的 index, 赋值给 `m.selectedtab.index` ✅
  - 计算 `paginator` 要翻多少页, 新建一个函数, 通过当前 index 来返回其页数, 比如每页 `5` 个, 当前为 `3`, 其页数为 `3/5+1 = 1`, 需要翻动的页数为两次页数之差 ✅


= 预计添加的功能
- 从 CONFIG 文件中读取用于打开链接的浏览器, 可以选择
- 复制 link


= Bug 相关
- 删掉 tab 会报错弹出 ✅

