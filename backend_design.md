# backend_design.md

### ***pyke elysia***

### 关于后端的结构设计

基本上包括三个主要的数据结构，user，team，item。

```json
// user
{
    "userUID": "uint",
    "userPassword": "uint",
    "teamsBelong": [
        "teamUID|score|percentComplete"
    ],
    "messions":[
        "string"
    ],
    "teamsOwn":[
        "string"
    ]
}
// team
{
    "teamUID": "uint",
    "teamLeader": "uint",
    "teamPassword": "uint",
    "membersInclude": [
        "string"
    ],
    "itemInclude": [
        "string"
    ]
}
// item
{
    "itemUID": "uint",
    "score": "uint",
    "shouldBeCompletedBy": "uint",
    "beCompletedBy": "uint",
    "isComplete": "bool"
}
```

其中 TeamsBelong 返回至 front end 的结构为

```json
{
    "TeamsBelong": [
        {
            "teamUID": "uint",
            "score": "uint",
            "percentComplete": "uint"
        }
    ]
}
```