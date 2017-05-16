package main

import (
	"log"

	"github.com/pjh130/go/common/uuid"
)

func main() {
	log.Println("uuid V1:", uuid.NewV1().String())
	log.Println("uuid V2:", uuid.NewV2(uuid.DomainPerson).String())
	log.Println("uuid V2:", uuid.NewV2(uuid.DomainGroup).String())
	log.Println("uuid V2:", uuid.NewV2(uuid.DomainOrg).String())
	//	log.Println("uuid V3:", uuid.NewV3().String())
	log.Println("uuid V4:", uuid.NewV4().String())
	//	log.Println("uuid V5:", uuid.NewV5().String())
}
