package oracle

import (
	"encoding/json"
	"net/http"
	"sync"
)

const (
	// Updated around 2021-07-01T10:57:17Z.
	DefaultAtomPrice float64 = 11.01
	DefaultIrisPrice float64 = 0.075752
	DefaultDvpnPrice float64 = 0.01979
	DefaultXprtPrice float64 = 8.3
	DefaultCroPrice  float64 = 0.113824
	DefaultAktPrice  float64 = 3.52
)

func GetAllPrices() [6]float64 {
	var waitGroup sync.WaitGroup

	waitGroup.Add(6)
	prices := [6]float64{
		GetAtomPrice(&waitGroup),
		GetIrisPrice(&waitGroup),
		GetDvpnPrice(&waitGroup),
		GetXprtPrice(&waitGroup),
		GetCroPrice(&waitGroup),
		GetAktPrice(&waitGroup),
	}
	waitGroup.Wait()

	return prices
}

func GetAtomPrice(waitGroup *sync.WaitGroup) float64 {
	type AtomPrice struct {
		Atom struct {
			Usd float64 `json:"usd"`
		} `json:"cosmos"`
	}

	defer waitGroup.Done()

	var price = AtomPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=cosmos&vs_currencies=usd")
	if err != nil {
		return 1
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	if price.Atom.Usd != 0 {
		return price.Atom.Usd
	} else {
		return DefaultAtomPrice
	}
}

func GetIrisPrice(waitGroup *sync.WaitGroup) float64 {
	type IrisPrice struct {
		Iris struct {
			Usd float64 `json:"usd"`
		} `json:"iris-network"`
	}

	defer waitGroup.Done()

	var price = IrisPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=iris-network&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	if price.Iris.Usd != 0 {
		return price.Iris.Usd
	} else {
		return DefaultIrisPrice
	}
}

func GetDvpnPrice(waitGroup *sync.WaitGroup) float64 {
	type DvpnPrice struct {
		Dvpn struct {
			Usd float64 `json:"usd"`
		} `json:"sentinel"`
	}

	defer waitGroup.Done()

	var price = DvpnPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=sentinel&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	if price.Dvpn.Usd != 0 {
		return price.Dvpn.Usd
	} else {
		return DefaultDvpnPrice
	}
}

func GetXprtPrice(waitGroup *sync.WaitGroup) float64 {
	type XprtPrice struct {
		Xprt struct {
			Usd float64 `json:"usd"`
		} `json:"persistence"`
	}

	defer waitGroup.Done()

	var price = XprtPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=persistence&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	if price.Xprt.Usd != 0 {
		return price.Xprt.Usd
	} else {
		return DefaultXprtPrice
	}
}

func GetCroPrice(waitGroup *sync.WaitGroup) float64 {
	type CroPrice struct {
		Cro struct {
			Usd float64 `json:"usd"`
		} `json:"crypto-com-chain"`
	}

	defer waitGroup.Done()

	var price = CroPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=crypto-com-chain&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	if price.Cro.Usd != 0 {
		return price.Cro.Usd
	} else {
		return DefaultCroPrice
	}
}

func GetAktPrice(waitGroup *sync.WaitGroup) float64 {
	type AktPrice struct {
		Akt struct {
			Usd float64 `json:"usd"`
		} `json:"akash-network"`
	}

	defer waitGroup.Done()

	var price = AktPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=akash-network&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	if price.Akt.Usd != 0 {
		return price.Akt.Usd
	} else {
		return DefaultAktPrice
	}
}
