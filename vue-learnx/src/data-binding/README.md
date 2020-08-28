# Vue2 数据绑定


## 数据绑定的通用流程

数据变化后更新视图
视图变化后更新数据

## Vue2 的数据绑定流程

![](./assets/data-binding-flow.png)

* Watcher 负责做具体通知更新对应的 View 元素
* Data的 会被 Vue Obserser 化
  * 通过 Object.defineProperty 方法，重写对象的属性的存取修饰器的方式, 为对象属性提供 getter 和 setter
* Watcher 和 Data 的联动
  * 涉及的概念：
    * 订阅发布对象（Dep 类）：每个对象和对象属性都会搭配一个 Dep 对象，用来存储 Watcher 引用
    * 全局对象：Vue2数据绑定中最黑的部分, 主要做共享数据
  * Watcher 与 Data 如何发生联系的
    * Watcher 所需的参数
      * Watcher 需要监听的 Data 对象（代码片段所示的 testObj），还有监听属性的访问表达式（'g.4.h')，最后还有一个回调函数
    * Watcher 初始化
      * 解析 访问表达式 得到 getter
      * 然后对 getter 进行封装 - this.get
        * 这里会利用 `全局对象` 把自己的引用暂存
        * 调用 getter 拿到最新的值
          * 这里会调用 Data 中对应值的 getter 取值，会将 `全局对象` 里的 Watcher 引用添加到值对应的 Dep 中
        * 清除下 `全局对象` 中的值引用
      * 调用 `this.get` 存储当前值到 `this.value` 中
    * Data 某值的更新
      * testObj.a = ?? 的表达式会触发对应的 `setter`
      * 更新属性值，然后调用 Dep 中 所有 watcher 的 updata 方法
      * watcher 调用callback


```js
new Watcher(testObj, 'g.4.h', function(newVal) {
  assert.strictEqual(newVal, 5);
});
```

