package main

import (
	"crypto/tls"
	"fmt"
	"github.com/raihankhan/golab/queue"
	"github.com/raihankhan/golab/stack"
)

func main() {

	// ........... Queue ................
	fmt.Println(".......Queue........")
	q := queue.New()
	q.Push(3, "dass", 4)
	fmt.Println(q.Front())
	fmt.Println(q.Get())

	q.Pop()
	fmt.Println(q.Front())
	fmt.Println(q.Get())

	q.Clear()
	fmt.Println(q.Size())

	// ........... Stack ................
	fmt.Println(".......Stack........")
	st := stack.New()
	st.Push(3, "dass", 4)
	fmt.Println(st.Top())
	fmt.Println(st.Get())

	st.Pop()
	fmt.Println(st.Top())
	fmt.Println(st.Get())

	st.Clear()
	fmt.Println(st.Size())

}

type TLS struct {
	enable   bool
	CertSpec CertSpec
}

type CertSpec struct {
	name          string
	host          string
	dnsName       []string
	acmeKey       string
	SelfSignedKey string
}

func (t *TLS) GetTLS(certtype string) TLS {
	if certtype == "ACME" {
		cert := ACMEcert{name: "", host: ""}
		cert.GetName("", "")
	} else if certtype == "Self-Singsdf" {
		cert := SelfSigned{name: "", host: ""}
		cert.GetName("", "")
	}

}

type IssuerType struct {
	ACME string
}

type ACMEcert struct {
	name string
	host string
}

type SelfSigned struct {
	name string
	host string
}

type Cert struct {
	Cert tls.Certificate
}

func (c *ACMEcert) GetName(name string, host string) Cert {
	x := "............."
	return Cert{Cert: x}
}

func (c *SelfSigned) GetName(name string, host string) Cert {
	x := "............."
	return Cert{Cert: x}
}

type certificate interface {
	GetName(name string, host string)
}
