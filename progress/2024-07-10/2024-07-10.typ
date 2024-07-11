= 16:18
- 把 del mode 改为 `delete` flag ✅
- 可以用 `add` 添加 `tab` ✅
- 跳转到新添加的 tab:
  - 用 `sort` 排序, 并找到新添加 tab 的 index, 赋值给 `m.selectedtab.index` ✅
  - 计算 `paginator` 要翻多少页, 新建一个函数, 通过当前 index 来返回其页数, 比如每页 `5` 个, 当前为 `3`, 其页数为 `3/5+1 = 1`, 需要翻动的页数为两次页数之差 ✅


= 预计添加的功能
- 从 CONFIG 文件中读取用于打开链接的浏览器, 可以选择:
  - 利用 `os/exec` 库来运行命令 ✅
  - 用一个 `textinput` 显示浏览器, 支持自己输入浏览器命令, 也可以从 suggestion 中补全, suggestion list 从 CONF 中读取 ✅
  - 用一个方框包裹 `textinput`, 默认显示 CONF 中读取的第一个 ✅
  - 选择要打开的链接, 用 confirm 获取输入 index ✅
  - 打开的键位为 o, 具体行为: 在 `display` mode 下按 `x` 进入 `confirm` mode, 提示获取输入, `enter` 之后打开网页 ✅
  - 添加 `browser` mode ✅
- 复制 link ✅


= Bug 相关
- 删掉 tab 会报错弹出 ✅

