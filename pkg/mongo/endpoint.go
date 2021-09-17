package mongo

import (
    "context"
    "github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
    InsertFruit endpoint.Endpoint
    DeleteFruit endpoint.Endpoint
    GetFruit    endpoint.Endpoint
    GetFruits   endpoint.Endpoint
}

func InsertFruit(service Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        fruit := request.(FruitRequest)

        return service.InsertFruit(&fruit), nil
    }
}

func GetFruits(service Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        fruit := request.(FruitRequest)

        return service.GetFruits(&fruit), nil
    }
}

func GetFruit(service Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        fruit := request.(FruitRequest)

        return service.GetFruit(&fruit), nil
    }
}

func DeleteFruits(service Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        fruit := request.(FruitRequest)

        return service.DeleteFruits(&fruit), nil
    }
}