package main

// Incorrect
//func First(query string, replicas ...Search) Result {
//	c := make(chan Result)
//	searchReplica := func(i int) { c <- replicas[i](query) }
//	for i := range replicas {
//		go searchReplica(i)
//	}
//	return <-c
//}

//right way
/*
func First(query string, replicas ...Search) Result {
	c := make(chan Result,len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result,1)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		default:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}
*/

func main() {
	// Blocked Goroutines and Resource Leaks
	// TODO something
}
