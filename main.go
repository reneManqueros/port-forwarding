package main

var settings Settings

func init() {
	settings.Load()
}

func main() {
	for key, _ := range settings.Redirections {
		thisRedir := settings.Redirections[key]
		go func(r *Redirection) {
			r.Listen()
		}(&thisRedir)
	}
	waitChan := make(chan int)
	<-waitChan
}
