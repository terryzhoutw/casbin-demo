# casbin-demo

## 策略配置（authz_policy.csv）

## 测试样例
- 鉴权成功
```
curl --location --request POST 'localhost:8080/unit/getunitbyid' \
--header 'x-tif-uid: uid-1'

## 响应
{
    "message": "getunitbyid"
}
```

- 鉴权失败`
```
curl --location --request POST 'localhost:8080/unit/getunitbyid' \
--header 'x-tif-uid: uid-111'

## 响应
{"message":"authz fail"}
```