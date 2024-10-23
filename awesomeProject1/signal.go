package main

func main() {
	/*
		sigs := make(chan os.Signal, 1)
		done := make(chan bool, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //将sigint和sigterm的信号传给sig
		//process start
		//get a sig interrupt,这个interrupt就是syscall.SIGINT返回的信号
		go func() {
			sig := <-sigs //将信号发送给sig
			fmt.Println("get a sig", sig)
			done <- true //标为结束
		}()
		fmt.Println("process start")
		<-done
	*/

}
