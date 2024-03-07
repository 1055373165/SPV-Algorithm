# SPV-Algorithm

简单支付验证算法实现

## 算法正确性演示
- 构建出来的默克尔树
![](resources/2024-03-07-16-33-19.png)
- 算法遍历结果
![](resources/2024-03-07-16-32-50.png)
- merkle branch（必须与图 1 中黄色部分完全对应）
![](resources/2024-03-07-16-33-59.png)
- 沿着默克尔路径重新构造出默克尔根
![](resources/2024-03-07-16-34-58.png)
- 算法压测结果
![image](https://github.com/1055373165/SPV-Algorithm/assets/33158355/e44c4e47-eefa-4db4-b185-00a15f26ed7f)
