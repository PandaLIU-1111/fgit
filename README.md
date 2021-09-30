# fgit

> 该项目是为了快速的使用一些常用的 git 组合命令

## 已内置命令

- version // 获取版本信息
- pushCommit "comment" // 代码上传，省略 `git add .`, `git commmit -m ""`,`git push` 三个命令
- cleanCheckout {{branch}} // 清理干净当前工作区，并且切换分支
- saveCheckout {{branch}} // 保存当前工作区修改，并且切换分支