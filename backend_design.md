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
        {
            "teamUID": "uint",
            "score": "uint",
            "percentComplete": "float"
        }
    ],
    "mession":[
        "uint"
    ],
    "teamOwn":[
        "uint"
    ]
}
// team
{
    "teamUID": "uint",
    "teamPassword": "uint",
    "memberInclude": [
        "uint"
    ],
    "itemInclude": [
        "uint"
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