package main

import "mycollection"

//例子程序如下：
/*
type Data struct{
}
func xxx()　*Data{
	var c Data
	return &c
}
func main(){
	xxx();
}
*/
//go run -gcflags "-m -l" main.go //-m表示进行内存分析　-l表示避免程序内联（避免程序优化）
//这样我们可以看到类似如下的输出：
//........
//./main.go: &c escapes to heap
//./main.go: moved to heap :c
//.........
func main() {
	//zhanglin.PrintfTest()
	//zhanglin.JsonTest()
	//zhanglin.LogicTest()
	//zhanglin.AlignTest()
	//mystruct.Mystruct()
	//mystruct.Mystruct2()
	//=========包================
	//mypkg.Mypkg()
	//=========变量================
	//myvar.Myvar()

	//=========基本语法================
	//mygolang.Mygolang()

	//=========集合================
	mycollection.Mycollection()
	//mycollection.MySort()
	//mycollection.TestMap()

	//=========函数================
	//myfunc.Myfunc()

	//=========延迟================
	//myfunc.Mydefer()
	//=========类================
	//mystruct.Mystruct()
	//mystruct.Mystruct2()

	//=========方法================
	//mymethod.Mymethod()
	//mymethod.Finallizefunc()

	//=========接口================
	//myintf.Myintf()
	//myintf.Myintf2()
	//myintf.Myintf3()
	//myintf.Myintf4()
	//myintf.TestAAABBBCCC()
	//myref.Myref()

	//=========错误处理================
	//myexception.Panic()
	//myexception.Myexception()
	//myexception.DeferTest()

	//=========基本api================
	//bytes.BytesBase()
	//bytes.BytesReader()
	//bytes.BytesBuffer()
	//myapi.Mystr()
	//myapi.MyString2()
	//myapi.MyReg()
	//myapi.Mytime()
	//myapi.MyJson()
	//myapi.MyLineFilter()
	//myapi.MyRandom()
	//myapi.MySha1()
	//myapi.MyBase64()
	//myapi.MyUrlParser()

	//=========并发================
	//mygoroutine.Mygoroutine1()
	//mygoroutine.Mygoroutine2()
	//mygoroutine.TestGoroutine()
	//mygoroutine.Myselect()
	//myruntime.Myruntime()
	//myruntime.NumCpuTest()
	//mysync.WaitGroupTest()
	//mysync.MyMutx()
	//mysync.MyRWLock()
	//mysync.MyAtomic()
	//mysync.MyOnce()
	//mysync.MyOnce2()
	//mysync.MyPooltest()
	//mytime.MyTicker()
	//mytime.MyTimeout()
	//mytime.MyTimer()

	//=========pprof================
	//mypprof.Testpprof()
	//mypprof.Testmypprof2()
	//=========文件io================
	//myfile.Myfile()
	//myfile.MyFile2()

	//=========命令行参数================
	//mycmdline.Mycmdline()

	//=========网络io================
	//go mytcp.MyTcpServer()
	// mytcp.MyTcpClient()
	//go onetoone.MyUdpServer()
	//onetoone.MyUdpCli()
	//go multi.MyMultiServer()
	//multi.MyMultiCli()
	//broadcast.MyBCServer()
	//broadcast.MyBCCli()

	//=========网络io-http相关================
	//mynet.MyHttp1()
	//mynet.MyHttp2()
	//mynet.MyHttp21()
	//mynet.MyHttp22()
	//mynet.MyHttp3()
	//mynet.MyHttp31()
	//mynet.Myhttpcli()

	//myupload.MyUpServer()
	//myupload.MyUpCli()

	//myssl.MySsl()
	//myssl.MySslCli()

	//myhttp2.MyHttp2Server()
	//myhttp2.MyHttp2Cli()

	//myws.MyWsServer()

	//=========系统================
	//mysys.MyExec()
	//mysys.MyExec2()
	//mysys.MySignal()
	//mysys.MySyscall()
	//mysys.MySyscall2()

	//=========go调用c================
	//goandc.Mygoandc()

	//=========虚拟化================
	//mycontainer.MyUts()
	//mycontainer.MyUts2()
	//mycontainer.MyPid()
	//mycontainer.MyCgrouptest()

}
