package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	frases := []string{
		"El mundo global es global.",
		"No hay que ser bilingüe, basta con tener la intención de serlo.",
		"I guant to bi very cliar de economic policis shud not mak… shudnot meikus forged. ",
		" La clave de una buena relación es una buena relación.",
		"Los acercamientos ayudan a estar más cerca.",
		"Sólo las mujeres están obligadas a conocer el mundo real.",
		"No hay mayor problema que cuando la gente cree que hay un problema.",
		"Si no recuerda algo diga que en su momento lo precisó.",
		"Si quiere saber la verdad lea libros sobre las mentiras de otros libros.",
		" El desarrollo económico impulsa el desarrollo económico.",
		"Como ya lo ha comentado el presidente de Francia, los temas tratados fueron varios, abordando distintos temas",
	}
	api := initializeAPI()
	go tweet(api, frases)
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)

}

func initializeAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("APIKEY"))
	anaconda.SetConsumerSecret(os.Getenv("APISECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESSTOKEN"), os.Getenv("ACCESSTOKENSECRET"))
	return api
}

func tweet(api *anaconda.TwitterApi, frases []string) {
	for {
		time.Sleep(10 * time.Minute)
		randi := randint(len(frases))
		v := url.Values{}
		v.Set("count", "1")
		//searchResult, _ := api.GetHomeTimeline(v)
		results, _ := api.GetUserTimeline(v)
		if results[0].Text != frases[randi] {
			api.PostTweet(frases[randi], nil)
		}
	}

}

func randint(length int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(length)
}
