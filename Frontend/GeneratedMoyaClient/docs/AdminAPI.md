# AdminAPI

All URIs are relative to *http://localhost:8098*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAllBookingsByUser**](AdminAPI.md#getallbookingsbyuser) | **GET** /bookings/all/ | Get All Bookings
[**getAllUsers**](AdminAPI.md#getallusers) | **GET** /users/all | Get all users


# **getAllBookingsByUser**
```swift
    open class func getAllBookingsByUser(completion: @escaping (_ data: [Booking]?, _ error: Error?) -> Void)
```

Get All Bookings

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient


// Get All Bookings
AdminAPI.getAllBookingsByUser() { (response, error) in
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

# **getAllUsers**
```swift
    open class func getAllUsers(completion: @escaping (_ data: [User]?, _ error: Error?) -> Void)
```

Get all users

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient


// Get all users
AdminAPI.getAllUsers() { (response, error) in
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

[**[User]**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

