# go-func
<p>提供针对Golang的函数式编程接口</p>
<p>Provide functional programming interface in golang</p>

使用指南:
<p>0.拉取工具包</p>
<p><code>go get github.com/ZaneYork/go-func/stream</code></p>
<p>1.引入工具包:</p>
<p><code>import . "github.com/ZaneYork/go-func/stream"</code></p>
<p>2.使用<code>NewSteam</code>或者<code>NewParallelStream</code>函数开启数据流</p>
<p>3.然后追加流操作函数以修改数据</p>
<p>4.使用如Collect、Reduce等流终结操作函数拿到处理结果集</p>
注意：
<p>本工具包提供的并发数据流在处理压力较低或者单核机器上时，并发处理为负优化，请斟酌选择是否并发</p>
<hr>
Usage:
<p>0.Clone source code</p>
<p><code>go get github.com/ZaneYork/go-func/stream</code></p>
<p>1.import package:</p>
<p><code>import . "github.com/ZaneYork/go-func/stream"</code></p>
<p>2.Start steam with function <code>NewSteam</code> or <code>NewParallelStream</code></p>
<p>3.Append operation function to modify element in stream</p>
<p>4.Terminate stream with any termination operation like Collect,Reduce,etc. </p>
Notice:
<p>The ParallelStream provided by this toolkit may result in negative optimization when meeting low processing pressure or on a single-core machine. Please choose whether or not to use</p>
