= Bug 相关
- 在改变终端大小后, 现在是优先显示底部字符串, 如何优先显示顶部呢?
- 删除第一个 tab 会报错 ✅
- 没有 tab 也会报错 ✅
- help 的显示也有问题 ✅
- node 横向显示重叠
- add 的 input 无法一次清空输入 ✅
- add 中, 同时添加新 tab 和 node 时不会添加 node ✅
- 测试添加 Box 对 string height 的影响 ✅


= 梳理下 add 功能
- 如果 `treelevel == 1 && tree 不存在`, 则添加 tree file, 更新 root tree
