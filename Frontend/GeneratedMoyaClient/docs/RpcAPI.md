# RpcAPI

All URIs are relative to *http://localhost:8098*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listBookingsInRange**](RpcAPI.md#listbookingsinrange) | **GET** /bookings/rpc/in_range | List Bookings in time frame cars


# **listBookingsInRange**
```swift
    open class func listBookingsInRange(startTime: Date? = nil, endTime: Date? = nil, completion: @escaping (_ data: [Booking]?, _ error: Error?) -> Void)
```

List Bookings in time frame cars

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let startTime = Date() // Date | Start time for filtering (optional)
let endTime = Date() // Date | End time for filtering (optional)

// List Bookings in time frame cars
RpcAPI.listBookingsInRange(startTime: startTime, endTime: endTime) { (response, error) in
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
 **startTime** | **Date** | Start time for filtering | [optional] 
 **endTime** | **Date** | End time for filtering | [optional] 

### Return type

[**[Booking]**](Booking.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

