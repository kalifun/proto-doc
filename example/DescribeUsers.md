# DescribeUsers

### 接口描述

- 接口路径: /api/v1/describe_users
- 接口描述: try get user info
- 访问限制: 
- 请求方式: post
- 接口版本: v1

### 输入参数

| 参数名称 | 受否必选 | 参数类型 | 参数注解 |
| -------- | -------- | -------- | -------- |
| UserIds | false | Array of Int64 | userids filter user ids |

### 输出参数

| 参数名称 | 受否必选 | 参数类型 | 参数注解 |
| -------- | -------- | -------- | -------- |
| Total | true | Int64 | total get all users count |
| User | true | Array of User | user list |

### 枚举值

#### UserStatus

| 枚举值 | 枚举注解 |
| ------ | -------- |
| UnLock | unlock  Unavailable |
| Lock | lock   Available |

### 复杂类型

#### User

> @desc: User info

| 参数名称 | 受否必选 | 参数类型 | 参数注解 |
| -------- | -------- | -------- | -------- |
| UserId | true | Int64 | userid description |
| UserName | true | String | user name |
| UserStatus | true | UserStatus | user status |
