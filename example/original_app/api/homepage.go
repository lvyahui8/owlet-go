package api

import (
    "owlet-go/example/original_app/service"
    "sync"
)

type HomepageFacade struct {
}

var cardProviders = []service.HomepageCardProvider{
    service.BalanceService{},
    service.BannerService{},
    service.TransactionService{},
}

var userService = service.UserService{}

func (hp HomepageFacade) GetData() interface{} {
    user := userService.GetLoggedUser()
    resChannel := make(chan interface{}, len(cardProviders))
    waitGroup := &sync.WaitGroup{}
    waitGroup.Add(len(cardProviders))
    for _, provider := range cardProviders {
        go func(cardProvider service.HomepageCardProvider) {
            defer waitGroup.Done()
            defer func() {
                err := recover()
                if err != nil {
                    if cardProvider.StrongDepend() {
                        panic(err)
                    }
                }
            }()
            resChannel <- map[string]interface{}{
                "cardId": cardProvider.Id(),
                "body":   cardProvider.Provide(user.Id),
            }
        }(provider)
    }
    waitGroup.Wait()
    res := make([]interface{}, 0)
    for item := range resChannel {
        res = append(res, item)
    }
    return res
}
