package main
import (
"fmt"
"flag"
"sync"
"time"
"bufio"
"os"
"net/http"
"net/url"
"io/ioutil"
"strings"
)

const (
	BannerColor  = "\033[1;34m%s\033[0m\033[1;36m%s\033[0m"
	TextColor = "\033[1;0m%s\033[1;32m%s\n\033[0m"
        InfoColor    = "\033[1;0m%s\033[1;35m%s\033[0m"
        NoticeColor  = "\033[1;0m%s\033[1;34m%s\n\033[0m"
        ErrorColor   = "\033[1;31m%s%s\033[0m"
)
func main () {

	var c int
	var p string

	flag.IntVar(&c, "C", 30, "Set the Concurrency")
	flag.StringVar(&p, "p", "", "The XSS Payload")
	flag.Parse()
	fmt.Printf(BannerColor,`

             █████                                        
            ░░███                                         
 █████ ████ ███████    ██████  █████ █████  █████   █████ 
░░███ ░███ ░░░███░    ███░░███░░███ ░░███  ███░░   ███░░  
 ░███ ░███   ░███    ░███ ░███ ░░░█████░  ░░█████ ░░█████ 
 ░███ ░███   ░███ ███░███ ░███  ███░░░███  ░░░░███ ░░░░███
 ░░████████  ░░█████ ░░██████  █████ █████ ██████  ██████ 
  ░░░░░░░░    ░░░░░   ░░░░░░  ░░░░░ ░░░░░ ░░░░░░  ░░░░░░  
                                                          
                                                          
	                     
                    
	`, "-- Coded by @tojojo -- \n")

	if p == ""{
		fmt.Println("Some Argument are not set")
		return
	}else {
		var wg sync.WaitGroup
		for i:=0; i<c; i++ {
			wg.Add(1)
			go func () {
				xss(p)
				wg.Done()
			}()
			wg.Wait()
		}
	}

}

func xss (payload string){
	

	time.Sleep(500 * time.Millisecond)
	scanner:=bufio.NewScanner(os.Stdin)
	
	for scanner.Scan(){
		link:=scanner.Text()
		u, err := url.Parse(link)
		if err != nil {
			return
		}
		qs := url.Values{}
		for param := range u.Query() {
			fmt.Printf(TextColor,"[*] Parameter:  ", param)
            qs.Set(param, payload)
		}
		u.RawQuery = qs.Encode()
		fmt.Printf(InfoColor,"[-] Testing:  ",u.String())
		req,err := http.Get(u.String())
       		if err != nil {
			return
        	}
        
        
        body, err := ioutil.ReadAll(req.Body)
				if err != nil {
	      			fmt.Println(err)
	   			}
	   			sb := string(body)
	   			check_result := strings.Contains(sb , payload)
	   			// fmt.Println(check_result)
	   			if check_result != false {
	   				fmt.Println(InfoColor,":XSS FOUND")
	   			}else{
	   				fmt.Println(ErrorColor,":Not Vulnerable:")
	   			}
	}
}
