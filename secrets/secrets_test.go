package secrets

import (
	"fmt"
	"testing"
)

func TestPublishHandler(t *testing.T){
	fmt.Println(GetSecretData("T01ATBKH4UU"))

}
