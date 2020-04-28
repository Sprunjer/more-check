  package create
  
 
  import (
    "log"
    "net/http"
    "io/ioutil"
    "fmt"
    "github.com/stellar/go/keypair"
    "github.com/stellar/go/clients/horizon"
    "io"
    "github.com/stellar/go/build"
    "github.com/stellar/go/clients/horizon"
    "fmt"
)
  
  func Create() {
  pair, err := keypair.Random()
    if err != nil {
        log.Fatal(err)
    }


    address := pair.Address()
    resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
    if err != nil {
        log.Fatal(err)
    }

    account, err := horizon.DefaultTestNetClient.LoadAccount(pair.Address())
    if err != nil {
        log.Fatal(err)
    }


    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))

    log.Println(pair.Seed())

    log.Println(pair.Address())


    fmt.Println("Balances for account:", pair.Address())

    for _, balance := range account.Balances {
        log.Println(balance)
    }
    
    }
