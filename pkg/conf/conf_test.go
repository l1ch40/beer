package conf

import "testing"

func TestAddSoftware(t *testing.T) {
	var args = [2]string{"3mux", "go get -u github.com/aaronjanse/3mux"}
	Add(args[:])
	yes := cfg.Section("").HasKey("3mux")
	if yes == false {
		t.Fatal("The 3mux doesn't exits.")
	}
	val := cfg.Section("").Key("3mux").Value()
	t.Log(val)
}

func TestInfo(t *testing.T) {
	value, err := Info("3mux")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
}

func TestList(t *testing.T) {
	beers, err := List()
	if err != nil {
		t.Error(err)
	} else {
		for _, beer := range beers {
			t.Logf("%+v", beer)
		}
	}
}

func TestUpdate(t *testing.T) {
	Update()
}
func TestUpgrade(t *testing.T) {
	Upgrade("3mux")
}

func TestRemove(t *testing.T) {
	var removeList = []string{"3mux"}
	Remove(removeList[:])
}