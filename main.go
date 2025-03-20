package main

func main() {
	cmd := (&Dic{}).Create()
	_ = cmd.Execute()
}
