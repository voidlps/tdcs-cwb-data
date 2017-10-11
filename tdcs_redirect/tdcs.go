package main

import (
  "fmt"
  "time"
  "net/http"
  "strings"
)

var (
  Taipei = time.FixedZone("Pacific", +8*3600)
  FiveMin = 5 * time.Minute
)

func tdcs_url(datatype string,t time.Time) string {
  d:=t.Format("20060102")
  h:=t.Format("15")
  f:=t.Format("20060102_1504")
  url:="http://tisvcloud.freeway.gov.tw/history/TDCS/"+datatype+"/"+d+"/"+h+"/TDCS_"+datatype+"_"+f+"00.csv"
  return url
}

func handler(w http.ResponseWriter, r *http.Request) {
  datatype:=strings.Replace(strings.ToUpper(r.RequestURI),"/","",1)
  fmt.Printf("datatype:%q\n",datatype)
  t:=time.Now().In(Taipei).Truncate(FiveMin)
  for n := 0; n <= 5; n++ {
     url:=tdcs_url(datatype,t.Add(time.Duration(-n)*FiveMin))
     fmt.Printf("try url:%q\n", url)
     resp, err := http.Head(url)
     if err != nil {
       fmt.Printf("error: %q\n", err)
     } else if resp.StatusCode == 200 {
       http.Redirect(w,r,url,302)
       fmt.Printf("return url:%q\n", url)
       return
     }
  }
  http.NotFound(w,r)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":1968", nil)
}
