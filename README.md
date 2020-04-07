# go-xo-game

## API Spec
### Create New Game
POST `http://localhost:3000/new_game`

Request Body :
```json
{
    "player_one": {
        "symbol": "X",
        "name": "KA"
    },
    "player_two": {
        "symbol": "O",
        "name": "PK"
    }
}
```

Response Body :
```json
{
    "board": [
        [
            "",
            "",
            ""
        ],
        [
            "",
            "",
            ""
        ],
        [
            "",
            "",
            ""
        ]
    ],
    "player_one": {
        "symbol": "X",
        "name": "KA",
        "score": 0
    },
    "player_two": {
        "symbol": "O",
        "name": "PK",
        "score": 0
    },
    "current_turn": "KA"
}
```
### Player Play
POST `http://localhost:3000/play`

Request Body :
```json
{
    "player": {
        "symbol": "X",
        "name": "KA"
    },
    "location_x":0,
    "location_y":0
}
```

Response Body :
```json
// game scenario
{
    "message": "NO WIN"
}
```
```json
// x mark player win scenario
{
    "message": "X WIN"
}
```
---
## Coding Convention
## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
datecalculate
```

## File Name
- camelCase ขึ้นต้นด้วยตัวใหญ่ เช่น
```
OrderService.java
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
datecalculate
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_ActionToTest_Input_152_Should_Be_150
```

## Variable Name
- ชื่อตัวแปรเป็นคำเดียวให้ตั้งชื่อเป็นพิมพ์เล็กทั้งหมด เช่น
```
day, month, year
```

- ชื่อตัวแปรมีความยาวตั้งแต่ 2 คำขึ้นไป ให้คำหลังขึ้นตันด้วยตัวอักษรตัวใหญ่เสมอ ในรูปแบบ **camelCase** เช่น
```
startDay, endMonth
```

- ชื่อตัวแปรเก็บค่าให้เติม "List" ต่อท้ายตัวแปรเสมอ เช่น
```
orderList

```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวอักษรพิมพ์ใหญ่ทั้งหมด เช่น
```
HOUR, MINUTE
```

## ข้อตกลง Commit Message ร่วมกัน
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน
