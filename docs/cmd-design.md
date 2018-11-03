## Agenda的命令及参数

### GoAgenda help
列出所有命令说明

### GoAgenda [command] help
列出具体命令说明

### GoAgenda register -u username -p password -e email -t telephone
用户注册

### GoAgenda login -u username -p password
用户登陆

### GoAgenda logout 
用户登出

### GoAgenda qryuser
用户查询

### GoAgenda deluser -p password
删除本用户

### GoAgenda cm -t title -p participator -s starttime -e endtime
创建会议

### GoAgenda delpar -t title -p participator
删除会议参与者

### GoAgenda addpar -t title -p participator
增加会议参与者

### GoAgenda qrymeet -s starttime -e endtime
查询会议

### GoAgenda canlmeet -t title -p password
取消会议

### GoAgenda quitmeet -t title -p password
退出会议

### GoAgenda emptymeet -p password
清空会议