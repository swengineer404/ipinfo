package ipinfo

import "testing"

func TestGet(t *testing.T) {
	info, err := Get("8.8.8.8")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", info)
}
