package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"code.google.com/p/goauth2/oauth/jwt"

	prediction "code.google.com/p/google-api-go-client/prediction/v1.6"
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

	token := createToken()

	if transport, err := jwt.NewTransport(token); err == nil {
		service, err := prediction.New(transport.Client())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		models := service.Trainedmodels.List("premium-ember-835")
		fmt.Printf("%#v\n", models)

	}

	go tweet(api, frases)
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	fmt.Println("Arreglandome el copete")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(os.Getenv("PORT"), nil)
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
			fmt.Printf("Tweedted: %s\n", frases[randi])
		}
	}

}

func randint(length int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(length)
}

func createToken() *jwt.Token {
	// Craft the ClaimSet and JWT token.
	pemKeyBytes, err := ioutil.ReadFile("pkey.pem")
	if err != nil {
		panic(err)
	}

	iss := "621740100651-7lkhc3notki0lce8ps60tf2ddk8je1p7@developer.gserviceaccount.com"
	scope := "https://www.googleapis.com/auth/prediction"
	t := jwt.NewToken(iss, scope, pemKeyBytes)

	return t
}
