# Agenda-test

Agenda测试文档

`Command`模块将使用命令进行手动测试

## 测试命令

### 用户

#### 用户注册

命令：`register`

- `user`：用户名
- `password`：密码
- `email`：用户邮箱
- `tel`：用户电话

功能：用户注册，判断用户名是否唯一，然后将用户信息存储到数据库

测试：注册3个用户test、test1、test2

```
GoAgenda register -u test -p test -e test@qq.com -t 1
Info: 2018/11/03 22:16:11 GoAgenda test  register succeed!
GoAgenda register -u test1 -p test1 -e test@qq.com -t 1
Info: 2018/11/03 22:16:42 GoAgenda test1  register succeed!
GoAgenda register -u test2 -p test2 -e test@qq.com -t 1
Info: 2018/11/03 22:16:55 GoAgenda test2  register succeed!
GoAgenda register -u test -p test -e test@qq.com -t 1
Error: 2018/11/03 22:17:04 GoAgenda test  register failed: username had been existed!
 # 注册失败，用户已存在
```

#### 用户登陆

命令：`login`

参数：

- `user`：用户名
- `password`：密码

功能：用户登陆，判断密码是否正确，如果正确则修改当前登陆的用户状态（当前用户登陆状态会在最后一项操作的3小时后自动登出）

测试：

```
GoAgenda login -u test3 -p test3
Error: 2018/11/03 22:23:01 GoAgenda test3  login failed: username does not exist!
#登录失败，用户不存在
GoAgenda login -u test -p test
Info: 2018/11/03 22:21:22 GoAgenda test  login succeed!
```

#### 用户登出

命令：`logout`

参数：无

功能：退出登陆，清理当前登陆状态

```
GoAgenda logout
Error: 2018/11/03 22:27:12 GoAgenda   logout failed: You did not login yet!
GoAgenda login -u test -p test
Info: 2018/11/03 22:28:12 GoAgenda test  login succeed!
GoAgenda logout
Info: 2018/11/03 22:28:34 GoAgenda test  logout succeed!
```

#### 用户查询

命令：`qryuser`

参数：无

功能：列出当前已注册的所有用户的用户名、邮箱及电话信息

```
GoAgenda qryuser
Error: 2018/11/03 22:29:05 GoAgenda   qryuser failed: You did not login yet!
# 未登录无法获取
GoAgenda login -u test -p test
Info: 2018/11/03 22:29:10 GoAgenda test  login succeed!
GoAgenda qryuser
Info: 2018/11/03 22:29:15 GoAgenda test  qryuser succeed!
There are  3  users!
Name-Email-Telephone
test   test@qq.com   1
test1   test@qq.com   1
test2   test@qq.com   1

```

#### 用户删除

命令：`deluser

参数：`password`：密码

功能：删除当前账户，清理登陆状态，移除相关的会议参与信息，并且删除无效会议

```
GoAgenda deluser -p test
Info: 2018/11/03 22:31:41 GoAgenda test  deluser succeed!
GoAgenda test1  qryuser succeed!
There are  2  users!
Name-Email-Telephone
test1   test@qq.com   1
test2   test@qq.com   1
```

### 会议

#### 创建会议

命令：`cm`

参数：

- `title`：会议主题
- `participator`：会议参与者（多个参与者用`+`分开）
- `start`：起始时间(`yyyy/MM/dd-hh:mm`，如：1998/03/07-11:23 要在当前时间之后)
- `end`：结束时间

功能：创建会议，检查参与者的合法性以及可行性。

```
GoAgenda cm -t meet -p test2 -s 2026-01-02/15:04:05 -e 2026-01-02/15:04:50
Info: 2018/11/03 22:38:36 GoAgenda test1  cm succeed! ---title=meet---participator=test2---starttime=2026-01-02/15:04:05---endtime=2026-01-02/15:04:50
```

#### 增加会议参与者

命令：`addpar`

参数：

- `participator`：新增的参与者
- `title`：会议标题

功能：增加会议参与者，检测合法性和可行性

```
GoAgenda addpar -t meet -p test3
Info: 2018/11/03 22:47:03 GoAgenda test1  addpar succeed! ---title=meet ---participator=test3
GoAgenda addpar -t meet -p test4
Error: 2018/11/03 22:47:24 GoAgenda test1  addpar failed: this participator is not registered # 不存在该用户
```

#### 移除会议参与者

命令：`delpar`

参数：

- `participator`：需要移除的参与者
- `title`：会议标题

功能：移除会议参与者，检测移除后的会议合法性

```
GoAgenda delpar -t meet -p test3
Info: 2018/11/03 22:48:25 GoAgenda test1  addpar succeed!---title=meet ---participator=test3
title:  meet
participator:  test3
```

#### 查询会议

命令： `qrymeet`

参数：

- `start`：开始的时间，默认为当前时间
- `end`：结束的时间，默认为10年后

功能：查询指定时间段与自己有关的（作为主持者或者参与者）的会议

```
GoAgenda qrymeet -s 2025-01-02/15:04:05 -e 2027-01-02/15:04:50
Info: 2018/11/03 23:02:12 GoAgenda test1  qrymeet succeed!
There are 1 meeting(s) you sponsor or participate:
Meeting 1
Title: meet
Sponsor: test1
Participators: [test2]
Start: 2026-01-02/15:04:05
End: 2026-01-02/15:04:50
```

#### 取消会议

命令：`canlmeet`

参数：

- `title`： Path参数，要取消的会议标题
- `password`: 密码

功能：取消指定标题的会议（自己发起的）

```
GoAgenda canlmeet -t meet -p test1
Info: 2018/11/03 22:53:46 GoAgenda test1  canlmeet succeed! ---title=meet
GoAgenda canlmeet -t meet -p test1
Error: 2018/11/03 22:54:05 GoAgenda test1  canlmeet failed: Meeting does not exist!
 #已删除，无会议
```

#### 退出会议

命令：`quitmeet`

参数：

- `title`： Path参数，要退出的会议标题

功能：退出指定标题的会议（自己参与的）

```
GoAgenda logout
Info: 2018/11/03 23:08:01 GoAgenda test1  logout succeed!
GoAgenda login -u test2 -p test2
Info: 2018/11/03 23:08:26 GoAgenda test2  login succeed!

```

#### 清空会议

命令：`emptymeet`

参数：`password`:密码

功能：清空自己发起的所有会议安排

```
GoAgenda qrymeet -s 2025-01-02/15:04:05 -e 2027-01-02/15:04:50
Info: 2018/11/03 23:04:29 GoAgenda test1  qrymeet succeed!
There are 1 meeting(s) you sponsor or participate:
Meeting 1
Title: meet
Sponsor: test1
Participators: [test2]
Start: 2026-01-02/15:04:05
End: 2026-01-02/15:04:50
GoAgenda emptymeet -p test1
Info: 2018/11/03 23:04:45 GoAgenda test1  emptymeet succeed!
GoAgenda qrymeet -s 2025-01-02/15:04:05 -e 2027-01-02/15:04:50
Info: 2018/11/03 23:04:49 GoAgenda test1  qrymeet succeed!
There are 0 meeting(s) you sponsor or participate:
```

### 其他

#### 帮助

命令：

`GoAgenda help`

列出所有命令说明

`GoAgenda [command] help`

列出具体命令说明

功能：查看帮助信息

