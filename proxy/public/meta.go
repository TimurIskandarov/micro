// Package classification Geo API.
//
// Документация Geo API.
//
//	Schemes: http, https
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
//	Security:
//	- Bearer
//
//
//	SecurityDefinitions:
//	  Bearer:
//	    type: apiKey
//	    name: Authorization
//	    in: header
//
// swagger:meta
package public

import "proxy/internal/model"

//go:generate `swagger generate spec -o public/swagger.json --scan-models`

// swagger:route POST /api/register auth RegisterRequest
//		Регистрация нового пользователя
// Security:
// - basic
// Responses:
// 	 200: body:RegisterResponse

// swagger:parameters RegisterRequest
type RegisterRequest struct {
	// Регистрация
	// in:body
	// required:true
	// example: {"name":"tim","password":"123"}
	Registration string
}

// swagger:model RegisterResponse
type RegisterResponse struct {
	// in:body
	// Token содержит информацию о регистрации
	Token string
}

// swagger:route POST /api/login auth LoginRequest
//		Авторизация пользователя
// Security:
// - basic
// Responses:
// 	 200: body:LoginResponse

// swagger:parameters LoginRequest
type LoginRequest struct {
	// Авторизация
	// in:body
	// required:true
	// example: {"name":"tim","password":"123"}
	Authorization string
}

// swagger:model LoginResponse
type LoginResponse struct {
	// in:body
	// Token содержит информацию о токене
	Token string
}

// swagger:route POST /api/address/geocode geo GeocodeRequest
//		Получение адреса по долготе и широте
// Security:
// - Bearer: []
// Responses:
//   200: body:GeocodeResponse

// GeocodeRequest security:
// - Bearer: []

// swagger:parameters GeocodeRequest
type GeocodeRequest struct {
	// Координаты
	// in:body
	// required:true
	// example: {"lat":"59.7221","lng":"30.4554"}
	Coordinates string
}

// swagger:model GeocodeResponse
type GeocodeResponse struct {
	// in:body
	Addresses model.Address
	// Addresses содержит информацию о адресах
}

// swagger:route POST /api/address/search geo AddressRequest
//		Получение координат по адресу
// Security:
// - Bearer: []
// Responses:
//	 200: body:AddressResponse

// AddressRequest security:
// - Bearer: []

// swagger:parameters AddressRequest
type AddressRequest struct {
	// Адрес
	// in:body
	// required:true
	// example: {"query": "москва тверская"}
	Query string
}

// swagger:model AddressResponse
type AddressResponse struct {
	// in:body
	// Addresses содержит информацию о координатах
	Addresses model.Address
}
