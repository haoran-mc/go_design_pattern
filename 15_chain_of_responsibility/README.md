责任链模式

It lets you create a chain of request handlers. For every incoming request, it is passed through the chain and each of the handler:

1. Processes the request or skips the processing.
2. Decides whether to pass the request to the next handler in the chain or not

一个医院的例子，医院会有很多部门：

- 接待部门 Reception
- 会诊部门 Doctor
- 药房部门 Medicine Room
- 收费部门 Cashier

显然当有病人来医院，他需要按顺序看病，首先到接待处，然后会诊，然后到医疗室，最后付费离开。











