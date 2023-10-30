## Structs, methods & interfaces 

#### What to learn

- 使用structs自定义类型

- 添加方法

- 声明接口

- table driven tests

#### Refactor

1. 使用struct存储数据，不同形状的数据

    ```go
      type Circle struct {
        Radius float64
      }

      type Rectangles struct {
        Width float64
        Height float64
      }
    ```
    
2. 添加方法，不同的形状添加对应的计算Area

    `func (receiverName ReceiverType) MethodName(args)`

3. 对test重构

    - 获取形状的struct，调用area，编写某种可以传递矩形和圆形的checkArea函数

    - 使用interfaces完成

4. `table driven tests`

5. 使用`table driven tests` 其中`{Rectangle{12, 6}, 72.0}`不明确，需要再次重构
