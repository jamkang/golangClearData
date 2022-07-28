# svn的基本操作

## 一、SVN启动模式

## 二、SVN链接和操作
- 库本版中需要增加一个readme的说明文件：**cat readme**
- 查看svn中文件状态：**svn status**
- **svn add**跟github中的一样
- 提交add了之后的代码，就算push了：**svn commit -m "SVN readme."**

## 三、版本控制
- 回退到未修改状态：**svn revert filename**,**svn revert -R trunkname**
- 查看历史信息：全部**svn log**，查看两个特定：**svn log -r 6:8**，单文件版本：**svn log 路劲**，目录的信息要加 **-v**，限定数量：**-l N**

## 四、分支管理
