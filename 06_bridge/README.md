桥接模式

有两种类型的电脑 mac 和 windows，同时有两种类型的打印机 epson 和 hp。

电脑和打印机需要联合工作，客户只需要在电脑上操作而不需要关心打印机是怎么工作的。

- 第一种设计方案是为每一种电脑都提供一个打印机
- 第二种设计方案是根据需要对电脑和打印机进行组合

对于有两个变化维度（即两个变化的原因）的系统，采用方案二来进行设计系统中类的个数更少，且系统扩展更为方便。设计方案二即是桥接模式的应用。桥接模式将继承关系转换为关联关系，从而降低了类与类之间的耦合，减少了代码编写量。
