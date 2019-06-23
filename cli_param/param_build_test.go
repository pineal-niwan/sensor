package cli_param

import (
	"github.com/pineal-niwan/sensor/cypher"
	"testing"
)

func TestBuildDb(t *testing.T) {
	key, err := hashCypherKey(`x-9@0!q1t-35-x7zz-ip@!q1357z1I^$35=7P8TYA-2+34`)
	t.Logf("key:%+v", key)
	if err != nil {
		t.Fail()
	}
	x, err := cypher.Encrypt([]byte(key), `a5-zkq!7l0-o@pp1-JNZF#`)
	if err != nil {
		t.Fail()
	}
	x, err = cypher.Decrypt([]byte(key), x)
	if err != nil {
		t.Fail()
	}
	t.Log(x)
	if x != `a5-zkq!7l0-o@pp1-JNZF#` {
		t.Fail()
	}
}
