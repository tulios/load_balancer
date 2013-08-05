package load_balancer

import (
  "fmt"
  "net/url"
  "net/http"
  "net/http/httputil"
  "math/rand"
  "time"
)

type LoadBalancer struct {
  hosts []Host
}

type Host struct {
  name string
  active int
  proxy *httputil.ReverseProxy
}

func New(hosts ...string) (*LoadBalancer, error) {
  balancer := new(LoadBalancer)
  balancer.hosts = make([]Host, len(hosts))

  for i, host := range hosts {
    url, err := url.Parse(host)
    if err != nil {
      fmt.Println("Failed to parse " + host)
      return nil, err
    }

    proxy := httputil.NewSingleHostReverseProxy(url)
    balancer.hosts[i] = Host{name: host, proxy: proxy}
  }

  return balancer, nil
}

func (l *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  host := l.selectHost()
  fmt.Printf("income connection to %s (active:%d)\n", host.name, host.active)
  host.active++
  host.proxy.ServeHTTP(w, r)
  host.active--
}

func (l *LoadBalancer) selectHost() Host {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  return l.hosts[r.Intn(len(l.hosts))]
}
