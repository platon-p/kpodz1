package main

func main() {
	dic := &Dic{}
	cmd := dic.Create()
	_ = cmd.Execute()
}
