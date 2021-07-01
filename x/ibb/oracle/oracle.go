package oracle

import (
	"encoding/json"
	"net/http"
	"sync"
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

	return price.Atom.Usd
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

	return price.Iris.Usd
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

	return price.Dvpn.Usd
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

	return price.Xprt.Usd
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

	return price.Cro.Usd
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

	return price.Akt.Usd
}
