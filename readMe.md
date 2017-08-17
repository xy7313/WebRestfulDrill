之前实习的时候需要实现一个自用drill，看了几篇关于goweb和restful的博客，在本地先写了个简陋的测试版。

## 用法：
##### 在终端
1. 在WebRestfulDrill上层目录下执行go install WebRestfulDrill
2. 执行 WebRestfulDrill

##### 在浏览器
1. localhost:8080
welcome!
2. localhost:8080/todos
a form to type queries as input

##### 另外
需要在本地启动drill服务，这样在todos页面输入SQL查询语句，点击查询之后，查询结果会在查询btn下面显示

##### 后来
这个东西对工作的项目来说并没有什么卵用，工作项目里，boss已经把一切都准备好了，只需要处理收集的数据，设置drill环境，在前端画个表就可以了┑(￣Д ￣)┍
