// Code generaTed by fileb0x at "2017-11-26 17:57:18.158384185 +0600 +06 m=+3.914702659" from config file "b0x.yml" DO NOT EDIT.

package swaggerFiles

import (
	"log"
	"os"
)

// FileOauth2RedirectHTML is "/oauth2-redirect.html"
var FileOauth2RedirectHTML = []byte("\x3c\x21\x64\x6f\x63\x74\x79\x70\x65\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x20\x6c\x61\x6e\x67\x3d\x22\x65\x6e\x2d\x55\x53\x22\x3e\x0a\x3c\x62\x6f\x64\x79\x20\x6f\x6e\x6c\x6f\x61\x64\x3d\x22\x72\x75\x6e\x28\x29\x22\x3e\x0a\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x3c\x2f\x68\x74\x6d\x6c\x3e\x0a\x3c\x73\x63\x72\x69\x70\x74\x3e\x0a\x20\x20\x20\x20\x27\x75\x73\x65\x20\x73\x74\x72\x69\x63\x74\x27\x3b\x0a\x20\x20\x20\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x72\x75\x6e\x20\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x6f\x61\x75\x74\x68\x32\x20\x3d\x20\x77\x69\x6e\x64\x6f\x77\x2e\x6f\x70\x65\x6e\x65\x72\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x49\x52\x65\x64\x69\x72\x65\x63\x74\x4f\x61\x75\x74\x68\x32\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x73\x65\x6e\x74\x53\x74\x61\x74\x65\x20\x3d\x20\x6f\x61\x75\x74\x68\x32\x2e\x73\x74\x61\x74\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x20\x3d\x20\x6f\x61\x75\x74\x68\x32\x2e\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x69\x73\x56\x61\x6c\x69\x64\x2c\x20\x71\x70\x2c\x20\x61\x72\x72\x3b\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x2f\x63\x6f\x64\x65\x7c\x74\x6f\x6b\x65\x6e\x7c\x65\x72\x72\x6f\x72\x2f\x2e\x74\x65\x73\x74\x28\x77\x69\x6e\x64\x6f\x77\x2e\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x2e\x68\x61\x73\x68\x29\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x71\x70\x20\x3d\x20\x77\x69\x6e\x64\x6f\x77\x2e\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x2e\x68\x61\x73\x68\x2e\x73\x75\x62\x73\x74\x72\x69\x6e\x67\x28\x31\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x71\x70\x20\x3d\x20\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x2e\x73\x65\x61\x72\x63\x68\x2e\x73\x75\x62\x73\x74\x72\x69\x6e\x67\x28\x31\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x61\x72\x72\x20\x3d\x20\x71\x70\x2e\x73\x70\x6c\x69\x74\x28\x22\x26\x22\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x61\x72\x72\x2e\x66\x6f\x72\x45\x61\x63\x68\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x76\x2c\x69\x2c\x5f\x61\x72\x72\x29\x20\x7b\x20\x5f\x61\x72\x72\x5b\x69\x5d\x20\x3d\x20\x27\x22\x27\x20\x2b\x20\x76\x2e\x72\x65\x70\x6c\x61\x63\x65\x28\x27\x3d\x27\x2c\x20\x27\x22\x3a\x22\x27\x29\x20\x2b\x20\x27\x22\x27\x3b\x7d\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x71\x70\x20\x3d\x20\x71\x70\x20\x3f\x20\x4a\x53\x4f\x4e\x2e\x70\x61\x72\x73\x65\x28\x27\x7b\x27\x20\x2b\x20\x61\x72\x72\x2e\x6a\x6f\x69\x6e\x28\x29\x20\x2b\x20\x27\x7d\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x6b\x65\x79\x2c\x20\x76\x61\x6c\x75\x65\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x6b\x65\x79\x20\x3d\x3d\x3d\x20\x22\x22\x20\x3f\x20\x76\x61\x6c\x75\x65\x20\x3a\x20\x64\x65\x63\x6f\x64\x65\x55\x52\x49\x43\x6f\x6d\x70\x6f\x6e\x65\x6e\x74\x28\x76\x61\x6c\x75\x65\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x29\x20\x3a\x20\x7b\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x73\x56\x61\x6c\x69\x64\x20\x3d\x20\x71\x70\x2e\x73\x74\x61\x74\x65\x20\x3d\x3d\x3d\x20\x73\x65\x6e\x74\x53\x74\x61\x74\x65\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x28\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2e\x73\x63\x68\x65\x6d\x61\x2e\x67\x65\x74\x28\x22\x66\x6c\x6f\x77\x22\x29\x20\x3d\x3d\x3d\x20\x22\x61\x63\x63\x65\x73\x73\x43\x6f\x64\x65\x22\x7c\x7c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2e\x73\x63\x68\x65\x6d\x61\x2e\x67\x65\x74\x28\x22\x66\x6c\x6f\x77\x22\x29\x20\x3d\x3d\x3d\x20\x22\x61\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x43\x6f\x64\x65\x22\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x29\x20\x26\x26\x20\x21\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2e\x63\x6f\x64\x65\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x21\x69\x73\x56\x61\x6c\x69\x64\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x2e\x65\x72\x72\x43\x62\x28\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x61\x75\x74\x68\x49\x64\x3a\x20\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2e\x6e\x61\x6d\x65\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x73\x6f\x75\x72\x63\x65\x3a\x20\x22\x61\x75\x74\x68\x22\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x65\x76\x65\x6c\x3a\x20\x22\x77\x61\x72\x6e\x69\x6e\x67\x22\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6d\x65\x73\x73\x61\x67\x65\x3a\x20\x22\x41\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x20\x6d\x61\x79\x20\x62\x65\x20\x75\x6e\x73\x61\x66\x65\x2c\x20\x70\x61\x73\x73\x65\x64\x20\x73\x74\x61\x74\x65\x20\x77\x61\x73\x20\x63\x68\x61\x6e\x67\x65\x64\x20\x69\x6e\x20\x73\x65\x72\x76\x65\x72\x20\x50\x61\x73\x73\x65\x64\x20\x73\x74\x61\x74\x65\x20\x77\x61\x73\x6e\x27\x74\x20\x72\x65\x74\x75\x72\x6e\x65\x64\x20\x66\x72\x6f\x6d\x20\x61\x75\x74\x68\x20\x73\x65\x72\x76\x65\x72\x22\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x71\x70\x2e\x63\x6f\x64\x65\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x65\x6c\x65\x74\x65\x20\x6f\x61\x75\x74\x68\x32\x2e\x73\x74\x61\x74\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2e\x63\x6f\x64\x65\x20\x3d\x20\x71\x70\x2e\x63\x6f\x64\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x2e\x63\x61\x6c\x6c\x62\x61\x63\x6b\x28\x7b\x61\x75\x74\x68\x3a\x20\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2c\x20\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x3a\x20\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x2e\x65\x72\x72\x43\x62\x28\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x61\x75\x74\x68\x49\x64\x3a\x20\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2e\x6e\x61\x6d\x65\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x73\x6f\x75\x72\x63\x65\x3a\x20\x22\x61\x75\x74\x68\x22\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x65\x76\x65\x6c\x3a\x20\x22\x65\x72\x72\x6f\x72\x22\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6d\x65\x73\x73\x61\x67\x65\x3a\x20\x22\x41\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x20\x66\x61\x69\x6c\x65\x64\x3a\x20\x6e\x6f\x20\x61\x63\x63\x65\x73\x73\x43\x6f\x64\x65\x20\x72\x65\x63\x65\x69\x76\x65\x64\x20\x66\x72\x6f\x6d\x20\x74\x68\x65\x20\x73\x65\x72\x76\x65\x72\x22\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x2e\x63\x61\x6c\x6c\x62\x61\x63\x6b\x28\x7b\x61\x75\x74\x68\x3a\x20\x6f\x61\x75\x74\x68\x32\x2e\x61\x75\x74\x68\x2c\x20\x74\x6f\x6b\x65\x6e\x3a\x20\x71\x70\x2c\x20\x69\x73\x56\x61\x6c\x69\x64\x3a\x20\x69\x73\x56\x61\x6c\x69\x64\x2c\x20\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x3a\x20\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x63\x6c\x6f\x73\x65\x28\x29\x3b\x0a\x20\x20\x20\x20\x7d\x0a\x3c\x2f\x73\x63\x72\x69\x70\x74\x3e\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "/oauth2-redirect.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileOauth2RedirectHTML)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
