package ja4

import (
	"fmt"

	utls "github.com/refraction-networking/utls"
)

// tlsVersion is the TLS version.
type tlsVersion uint16

// numberOfCipherSuites is the number of cipher suites.
type numberOfCipherSuites int

// numberOfExtensions is the number of extensions.
type numberOfExtensions int

// cipherSuites is a slice of cipher suites.
type cipherSuites []uint16

// extensions is a slice of extensions.
type extensions []uint16

// signatureAlgorithms is a slice of signature algorithms.
type signatureAlgorithms []uint16

func (x tlsVersion) String() string {
	switch uint16(x) {
	case utls.VersionTLS10:
		return "10"
	case utls.VersionTLS11:
		return "11"
	case utls.VersionTLS12:
		return "12"
	case utls.VersionTLS13:
		return "13"
	}
	return "00"
}
func (x numberOfCipherSuites) String() string { return fmt.Sprintf("%02d", min(x, 99)) }
func (x numberOfExtensions) String() string   { return fmt.Sprintf("%02d", min(x, 99)) }
func (x cipherSuites) String() string         { return joinUint16(x, cipherSuitesSeparator) }
func (x extensions) String() string           { return joinUint16(x, extensionsSeparator) }
func (x signatureAlgorithms) String() string  { return joinUint16(x, signatureAlgorithmSeparator) }
