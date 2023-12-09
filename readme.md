# VirtualFileSystem

## Method

## 要寫的東西

- Options 的資料大小寫不分
- UnitTest 全部都要補
- 看要不要把重複地寫成 interface(C++ 的 template)

## issue

我看這篇文章說
https://stackoverflow.com/questions/54377597/how-to-make-a-function-that-receives-an-array-of-custom-interface

go 沒辦法在 interface method 中的 parma 使用 slice.
因為她只能轉型 object 不能轉型 object's slice.

所以我沒辦法寫一個指標 Interface 來多型? 對嘛
