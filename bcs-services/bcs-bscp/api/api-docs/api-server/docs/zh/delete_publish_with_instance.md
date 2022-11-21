### 描述
该接口提供版本：v1.0.0+
 
删除实例发布。

### 输入参数
| 参数名称     | 参数类型     | 必选   | 描述             |
| ------------ | ------------ | ------ | ---------------- |
| biz_id         | uint32       | 是     | 业务ID     |
| app_id         | uint32       | 是     | 应用ID     |
| id         | string       | 是     | 实例发布记录ID     |


### 调用示例
```json
```

### 响应示例
```json
{
    "code": 0,
    "message": "ok"
}
```

### 响应参数说明
| 参数名称     | 参数类型   | 描述                           |
| ------------ | ---------- | ------------------------------ |
|      code        |      int32      |            错误码                   |
|      message        |      string      |             请求信息                  |