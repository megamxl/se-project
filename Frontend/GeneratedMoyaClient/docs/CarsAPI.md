# CarsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCar**](CarsAPI.md#addcar) | **POST** /cars | Add a new car
[**deleteCar**](CarsAPI.md#deletecar) | **DELETE** /cars | Delete a car
[**getCarByVin**](CarsAPI.md#getcarbyvin) | **GET** /carByVin | Get a car by VIN
[**listCars**](CarsAPI.md#listcars) | **GET** /cars | List available cars
[**updateCar**](CarsAPI.md#updatecar) | **PUT** /cars | Update car details


# **addCar**
```swift
    open class func addCar(car: Car, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Add a new car

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let car = Car(VIN: "VIN_example", model: "model_example", brand: "brand_example", imageURL: "imageURL_example", pricePerDay: 123) // Car | 

// Add a new car
CarsAPI.addCar(car: car) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **car** | [**Car**](Car.md) |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteCar**
```swift
    open class func deleteCar(VIN: String, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Delete a car

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let VIN = "VIN_example" // String | 

// Delete a car
CarsAPI.deleteCar(VIN: VIN) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **VIN** | **String** |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCarByVin**
```swift
    open class func getCarByVin(VIN: String, completion: @escaping (_ data: Car?, _ error: Error?) -> Void)
```

Get a car by VIN

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let VIN = "VIN_example" // String | 

// Get a car by VIN
CarsAPI.getCarByVin(VIN: VIN) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **VIN** | **String** |  | 

### Return type

[**Car**](Car.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listCars**
```swift
    open class func listCars(currency: Currency, startTime: Date? = nil, endTime: Date? = nil, completion: @escaping (_ data: [CarListInner]?, _ error: Error?) -> Void)
```

List available cars

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let currency = Currency() // Currency | The currency The user want to pay in
let startTime = Date() // Date | Start time for filtering cars based on availability (optional)
let endTime = Date() // Date | End time for filtering cars based on availability (optional)

// List available cars
CarsAPI.listCars(currency: currency, startTime: startTime, endTime: endTime) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](.md) | The currency The user want to pay in | 
 **startTime** | **Date** | Start time for filtering cars based on availability | [optional] 
 **endTime** | **Date** | End time for filtering cars based on availability | [optional] 

### Return type

[**[CarListInner]**](CarListInner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateCar**
```swift
    open class func updateCar(car: Car, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Update car details

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let car = Car(VIN: "VIN_example", model: "model_example", brand: "brand_example", imageURL: "imageURL_example", pricePerDay: 123) // Car | 

// Update car details
CarsAPI.updateCar(car: car) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **car** | [**Car**](Car.md) |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

