## 一个非常简陋的Onebot模拟客户端

用于测试onebot插件是否能正常使用

**Action只实现了两个, 所以会有很多功能无法使用**

### 用法

1. 修改`config.yml`
2. 命令:
```shell
private_message [user_id] [message]
group_message [group_id] [user_id] [message]
friend_increase [user_id]
friend_decrease [user_id]
group_increase [group_id] [user_id] [operator_id]
group_decrease [group_id] [user_id] [operator_id]
private_message_delete [message_id] [user_id]
group_message_delete [message_id] [group_id] [user_id] [operator_id]
```
为了方便使用, 所有的命令均可以用首字母代替, 如`pm`代替`private_message`

### 例子

```shell
-> pm 123 echo hello
-> 2023/01/06 00:16:57 send_message: private[123] :echo hello
```