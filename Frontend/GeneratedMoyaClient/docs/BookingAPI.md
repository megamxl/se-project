# BookingAPI

All URIs are relative to *http://localhost:8098*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bookCar**](BookingAPI.md#bookcar) | **POST** /booking | Book a car
[**deleteBooking**](BookingAPI.md#deletebooking) | **DELETE** /booking | Cancel a booking
[**getBookingById**](BookingAPI.md#getbookingbyid) | **GET** /booking/{id} | Get a specific booking
[**getBookings**](BookingAPI.md#getbookings) | **GET** /booking | Get all bookings by a user
[**updateBooking**](BookingAPI.md#updatebooking) | **PUT** /booking | Update a booking


# **bookCar**
```swift
    open class func bookCar(bookCarRequest: BookCarRequest, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Book a car

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let bookCarRequest = bookCar_request(VIN: "VIN_example", currency: Currency(), startTime: Date(), endTime: Date()) // BookCarRequest | 

// Book a car
BookingAPI.bookCar(bookCarRequest: bookCarRequest) { (response, error) in
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
 **bookCarRequest** | [**BookCarRequest**](BookCarRequest.md) |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteBooking**
```swift
    open class func deleteBooking(bookingId: String, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Cancel a booking

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let bookingId = "bookingId_example" // String | 

// Cancel a booking
BookingAPI.deleteBooking(bookingId: bookingId) { (response, error) in
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
 **bookingId** | **String** |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getBookingById**
```swift
    open class func getBookingById(id: String, completion: @escaping (_ data: [Booking]?, _ error: Error?) -> Void)
```

Get a specific booking

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let id = "id_example" // String | 

// Get a specific booking
BookingAPI.getBookingById(id: id) { (response, error) in
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
 **id** | **String** |  | 

### Return type

[**[Booking]**](Booking.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getBookings**
```swift
    open class func getBookings(completion: @escaping (_ data: [Booking]?, _ error: Error?) -> Void)
```

Get all bookings by a user

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient


// Get all bookings by a user
BookingAPI.getBookings() { (response, error) in
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
This endpoint does not need any parameter.

### Return type

[**[Booking]**](Booking.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateBooking**
```swift
    open class func updateBooking(updateBookingRequest: UpdateBookingRequest, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Update a booking

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let updateBookingRequest = updateBooking_request(bookingId: "bookingId_example", status: "status_example") // UpdateBookingRequest | 

// Update a booking
BookingAPI.updateBooking(updateBookingRequest: updateBookingRequest) { (response, error) in
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
 **updateBookingRequest** | [**UpdateBookingRequest**](UpdateBookingRequest.md) |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

