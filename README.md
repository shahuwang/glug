模仿Elixir的Plug库

并没有写完，懒了

主要提供两个功能，路由，过滤器，每个请求，先过滤一遍，然后决定是否还需要处理。

Go的http库本身已经很完整了，所以直接把request和response暴露给handler反而更方便。

