# backend_design.md

### ***pyke elysia***

### 关于后端的结构设计

基本上包括三个主要的数据结构，user，team，item。

```json
// user
{
    "userUID": "uint",
    "userPassword": "uint",
    "teamBelong": [
        "teamUID|score|percentComplate"
    ],
    "messions":[
        "string"
    ],
    "teamOwn":[
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