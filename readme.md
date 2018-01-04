﻿# 异步HttpClient原理实验


### 1、动机   

![此处输入图片的描述][1]
  设想一个旅行社，其信息系统由多个基本服务组成。这些服务可能使用不同的技术（JMS，EJB，WS，...）来构建。为了简单起见，我们假设服务可以通过HTTP方法调用（例如使用JAX-RS客户端）使用REST接口来使用。我们也假定我们需要的基本服务是：  

    客户服务 - 提供有关旅行社客户的信息。  

    目的地服务 - 提供已认证客户的已访问和推荐目的地列表。  

    天气服务 - 提供给定目的地的天气预报。  

    报价服务 - 为客户提供推荐目的地的价格计算。  

其任务是创建一个公开可用的功能，对于经过认证的用户，将显示最后访问的10个地点的列表，并显示10个新的推荐目的地列表，包括用户的天气预报和价格计算。请注意，某些请求（检索数据）取决于以前请求的结果。例如。获取推荐的目的地取决于首先获取有关认证用户的信息。获取天气预报取决于目的地信息等。一些请求之间的这种关系是问题的一个重要部分，并且您可以真正利用反应性编程模型。  

如何获取数据的一种方式是从客户端（例如移动设备）向所涉及的所有服务进行多个HTTP方法调用，并将检索到的数据组合在客户端上。但是，由于基本服务只能在内部网络中使用，我们宁愿创建公共编排层，而不是将所有内部服务暴露给外部世界。编配层只向公众公开所需的基本服务操作。为了限制流量并实现更低的延迟，我们希望通过一个响应将所有必要的信息返回给客户端。  

编排层如图6.1所示。该层接受来自外部的请求，并负责调用对内部服务的多个请求。当来自内部服务的响应在编排层中可用时，它们被合并成单个响应，并被发送回客户端。  
### 2、天真做法  
　![此处输入图片的描述][2]
　　
### 3、优化做法  
![此处输入图片的描述][3]

### 4、做法
天真做法结果：  
![此处输入图片的描述][4]  
优化做法结果：  
![此处输入图片的描述][5]  
对比两种做法，显然第二种做法延迟会低很多。

  [1]: https://raw.githubusercontent.com/thougr/Agency/master/screenshot/6.1.png
  [2]: https://raw.githubusercontent.com/thougr/Agency/master/screenshot/6.2.png
  [3]: https://raw.githubusercontent.com/thougr/Agency/master/screenshot/6.3.png
  [4]: https://raw.githubusercontent.com/thougr/Agency/master/screenshot/naivetime.png
  [5]: https://raw.githubusercontent.com/thougr/Agency/master/screenshot/optimized.png
