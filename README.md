# hello-word
This is my notes when I sutdy Golang.


1. 一个int 变量A: 它的本质是：
    1) 计算机的内存中有一小块内存，这个小块内存的名字是A
    2) 这个名字为A的小块内存，只能存放整形值：int,比如 2，放2.1就不行。
    3）这小块内存存放的值可以随时被程序修改。比如修改为 3.
    4）这小块内存可以被另外一个小块内存赋值，比如A=B,但前提是B也是整形。   
    4）这个小块内存的地址是：0x12345，可以用&A打印
    5）一个指针变量B如果被赋值为&A,则名称为B的这小块内存里存的值为16进制：0x12345


2. 一个函数func A（int x) int: 它的本质是：
    1) 计算机的内存中有一小块内存，这个小块内存的名字是A
    2) 这个名字为A的小块内存，存放了一些代码。 比如说是计算x的平方
    3）这小块内存存放的代码可以随时被程序修改。比如修改为计算x的根
    4）这个存放函数的小块内存与存放变量的小块内存的区别是：
       一是：存放函数的内存稍大一些。
       二是：这块内存的类型是由这小块内存的输入和输出参数决定，称为签名
       同样类型的变量可以相互赋值，同样签名的函数也可以相互赋值。
    5）这小内存存放的代码可以被另外一小块存放代码的内存赋值，比如A=B, 但前提是A的输入输出参数
       与B完全一样。也就是说B这小块内存的输入是一个int, 输出也是一个int.
    6）这个小块内存的地址是：0x12345，也可以打印
    7）一个指针变量B如果被赋值为A,则名称为B的这小块内存里存的值为16进制：0x12345
 
3. 结论：函数其实就是变量，是一种特殊的变量。      
