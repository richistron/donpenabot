package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("APIKEY"))
	anaconda.SetConsumerSecret(os.Getenv("APISECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESSTOKEN"), os.Getenv("ACCESSTOKENSECRET"))
	v := url.Values{}
	v.Set("count", "300")
	searchResult, err := api.GetSearch("epn", v)

	if err != nil {
		panic(err)
	}

	f, err := os.Create("seeds.txt")
	defer f.Close()
	buffer := bytes.Buffer{}
	for _, tweet := range searchResult.Statuses {
		tweetbytes := []byte(fmt.Sprintf(`"%s"\n`, tweet.Text))
		buffer.Write(tweetbytes)
	}
	f.Write(buffer.Bytes())
	f.Sync()
}
