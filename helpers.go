package univush

import (
	"crypto/tls"

	"github.com/sideshow/apns2/certificate"
)

// CertBytes retuns a new certificate using a bytes
func CertBytes(bytes []byte, password string) (tls.Certificate, error) {
	cert, err := certificate.FromP12Bytes(bytes, password)
	if err != nil {
		cert, err = certificate.FromPemBytes(bytes, password)
	}
	return cert, err
}
