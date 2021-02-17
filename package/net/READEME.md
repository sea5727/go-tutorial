# Package net

[Overview](## Overview)
 
## Overfview

TCP/IP, UDP, DNS, Unix Domain Socket 을 포함한 네트워크 I/O 인터페이스를 제공한다.

이 패키지는 low-level 네트워킹 접근을 제공하지만, 대부분 어플리케이션에선 기본 인터페이스만을 사용하면 된다. (Dial, Listen, Accept함수와 Conn, Listener 인터페이스가 있다. )

crypto/tls 패키지는 Dial과 Listen 펑션과 비슷하고, 같은 인터페이스를 사용한다.

```go
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
    // handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
```