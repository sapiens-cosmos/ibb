package oracle

import (
	"encoding/json"
	"net/http"
)

func GetAllPrices() [6]float64 {
	return [6]float64{
		GetAktPrice(),
		GetAtomPrice(),
		GetCroPrice(),
		GetDvpnPrice(),
		GetIrisPrice(),
		GetXprtPrice(),
	}
}

func GetAtomPrice() float64 {
	type AtomPrice struct {
		Atom struct {
			Usd float64 `json:"usd"`
		} `json:"cosmos"`
	}

	var price = AtomPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=cosmos&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	return price.Atom.Usd
}

func GetIrisPrice() float64 {
	type IrisPrice struct {
		Iris struct {
			Usd float64 `json:"usd"`
		} `json:"iris-network"`
	}

	var price = IrisPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=iris-network&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	return price.Iris.Usd
}

func GetDvpnPrice() float64 {
	type DvpnPrice struct {
		Dvpn struct {
			Usd float64 `json:"usd"`
		} `json:"sentinel"`
	}

	var price = DvpnPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=sentinel&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	return price.Dvpn.Usd
}

func GetXprtPrice() float64 {
	type XprtPrice struct {
		Xprt struct {
			Usd float64 `json:"usd"`
		} `json:"persistence"`
	}

	var price = XprtPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=persistence&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	return price.Xprt.Usd
}

func GetCroPrice() float64 {
	type CroPrice struct {
		Cro struct {
			Usd float64 `json:"usd"`
		} `json:"crypto-com-chain"`
	}

	var price = CroPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=crypto-com-chain&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	return price.Cro.Usd
}

func GetAktPrice() float64 {
	type AktPrice struct {
		Akt struct {
			Usd float64 `json:"usd"`
		} `json:"akash-network"`
	}

	var price = AktPrice{}
	r, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=akash-network&vs_currencies=usd")
	if err != nil {
		return 0
	}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&price)

	return price.Akt.Usd
}
