# <> Documentation

## Overview

Package `ja4` implements JA4 algorithm based on [utls](https://github.com/refraction-networking/utls).

---

## Index

- [type JA4Fingerprint](#types)
  - [func (j *JA4Fingerprint) String() string](#type-ja4fingerprint)
    - [func (j *JA4Fingerprint) Unmarshal(chs *utls.ClientHelloSpec, protocol byte) error](#func-ja4fingerprint-string)
    - [func (j *JA4Fingerprint) UnmarshalBytes(clientHelloRecord []byte, protocol byte error)](#func-ja4fingerprint-unmarshal)




 

---

## Constants

*This section is empty.*

---

## Variables

*This section is empty.*

---

## Functions

*This section is empty.*

---

## Types

### type [JA4Fingerprint](https://github.com/wu238121-a11y/go-ja4/blob/main/ja4.go#L21)

<pre>
<code class="go">
type JA4Fingerprint struct {
    Protocol    <a href="https://pkg.go.dev/builtin#byte" target="_blank">byte</a>    // (Go builtin)
    TLSVersion  tlsVersion
    SNI         <a href="https://pkg.go.dev/builtin#byte" target="_blank">byte</a>    // (Go builtin)
    NumberOfCipherSuites numberOfCipherSuites
    NumberOfExtensions   numberOfExtensions
    FirstALPN   <a href="https://pkg.go.dev/builtin#string" target="_blank">string</a>  // (Go builtin)

    CipherSuites         cipherSuites

    Extensions           extensions
    SignatureAlgorithms  signatureAlgorithms
}
</code>
</pre>

### func (*JA4Fingerprint) [String](https://github.com/wu238121-a11y/go-ja4/blob/main/ja4.go#L81)

<pre><code class="language-go">func (j *<a href="#type-ja4fingerprint">JA4Fingerprint</a>) String() <a href="https://pkg.go.dev/builtin#string">string</a></code></pre>


### func (*JA4Fingerprint) [Unmarshal](https://github.com/wu238121-a11y/go-ja4/blob/main/ja4.go#L58)

<pre><code class="language-go">func (j *<a href="#type-ja4fingerprint">JA4Fingerprint</a>) Unmarshal(chs *<a href="https://pkg.go.dev/github.com/refraction-networking/utls">utls</a>.<a href="https://pkg.go.dev/github.com/refraction-networking/utls#ClientHelloSpec">ClientHelloSpec</a>, protocol <a href="https://pkg.go.dev/builtin#byte">byte</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code></pre>

### func (*JA4Fingerprint) [UnmarshalBytes](https://github.com/wu238121-a11y/go-ja4/blob/main/ja4.go#L47)

<pre><code class="language-go">func (j *<a href="#type-ja4fingerprint">JA4Fingerprint</a>) UnmarshalBytes(clientHelloRecord []<a href="https://pkg.go.dev/builtin#byte">byte</a>, protocol <a href="https://pkg.go.dev/builtin#byte">byte</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code></pre>