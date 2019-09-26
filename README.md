# LevelDB Server设计

## 什么是LevelDB Server

LevelDB Server 是基于Golang开发的LevelDB的服务端，对外提供http服务，可接收单次写入，批量写入，删除，范围读取等操作

## 结构说明

LevelDB服务监听8360端口，主协程负责接收外部请求，并同时启动10个子协程，其中五个子协程负责写操作，五个子协程负责读操作，并设定一个长度为10的写channel，所有的写入操作进入channel之后由写协程进行读取操作



## 接口说明

#### readOne

Method: GET

Params:  

- key:string

Result:

```json
{
    "message":"success",
    "result":true,
    "data":bytes
}
```

#### delete

Method: DELETE

Params:

- key:string

Result:

```json
{
    "message":"success",
    "result":true,
    "data":null
}
```

#### deleteAndGet

Method: DELETE

Params:

- key:string

Result:

```json
{
    "message":"success",
    "result":true,
    "data":bytes
}
```

### put

Method: PUT

Body:

- key:string
- value:object

Result:

```json
{
    "message":"success",
    "result":true,
    "data":null
}
```

#### batchWrite

Method: PUT

Body:

- data:list
  - key:string
  - value:object

Result:

```json
{
    "message":"success",
    "result":true,
    "data":null
}
```

#### range

Method: GET

Params:

- start:string
- end:string

Result:

```json
{
    "message":"success",
    "result":true,
    "data":[]bytes
}
```

#### rangeByPrefix

Method:GET

Params:

- prefix:string

Result:

```json
{
    "message":"success",
    "result":true,
    "data":[]bytes
}
```

