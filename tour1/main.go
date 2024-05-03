
//собес вопросы

package main

func main() {
  var urls = []string{
    "http://ozon.ru",
    "https://ozon.ru",
    "http://google.com",
    "http://somesite.com",
    "http://non-existent.domain.tld",
    "https://ya.ru",
    "http://ya.ru",
    "http://ёёёё",
  }

    n := 2
    chGo := make(chan struct{}, n)
 
    ch := make(chan string)
    wg = &sync.WaitGroup
  for _, url := range urls {
    wg.add(1)
    chGo<-struct{}{}
    go getUrl(url, ch, chGo, wg) 
  }

  go func(){
wg.Wait()
close(ch)
}()
  
  for code := range ch {
     fmt.Println(code)
  }
    
}


func getUrl(url string, ch chan string, wg *sync.WaitGroup) {

    defer wg.Done()
   code := http.Get(url); 
   if code == "200" {
      ch<-"ok"
    } else {
      ch<-code
    }
    <-chGo
}
=====================

type Cache interface {
    Set(k, v string)
    Get(k string) (v string, ok bool)
}


type MemCache struct{
    mx sync.RWMutex
    mem map[string]string
}

func NewMemCache() &MemCache {
     c := &MemCache{
        mx sync.Mutex
        mem make(map[string]string)
        }
    return c
}


func (c *MemCache) Set(k, v string){
    c.mx.Lock()
defer c.mx.Unlock()
    c.mem[k] = v
}

func (c *MemCache) Get(k string) (v string, ok bool) {
    c.mx.RWLock()
    defer c.mx.RWUnlock()
    var ok bool
    var v string
    if ok, v := mem[k]; ok {
        ok = true
      return 
    }

    return 
}

============================

SELECT * FROM employee WHERE sex = 'm' AND salary > 300000 AND age = 20 ORDER BY created_at

age 
salary
m
created_at
