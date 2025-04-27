# UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteUser**](UserAPI.md#deleteuser) | **DELETE** /users/all | Delete a user
[**getUsers**](UserAPI.md#getusers) | **GET** /users | Get user info
[**login**](UserAPI.md#login) | **POST** /login | User login using email and password
[**registerUser**](UserAPI.md#registeruser) | **POST** /users | Register a new user
[**updateUser**](UserAPI.md#updateuser) | **PUT** /users/all | Update user details


# **deleteUser**
```swift
    open class func deleteUser(id: String, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Delete a user

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let id = "id_example" // String | 

// Delete a user
UserAPI.deleteUser(id: id) { (response, error) in
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

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getUsers**
```swift
    open class func getUsers(completion: @escaping (_ data: User?, _ error: Error?) -> Void)
```

Get user info

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient


// Get user info
UserAPI.getUsers() { (response, error) in
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

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **login**
```swift
    open class func login(loginRequest: LoginRequest, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

User login using email and password

Authenticate a user using their email and password.

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let loginRequest = login_request(email: "email_example", password: "password_example") // LoginRequest | 

// User login using email and password
UserAPI.login(loginRequest: loginRequest) { (response, error) in
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
 **loginRequest** | [**LoginRequest**](LoginRequest.md) |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **registerUser**
```swift
    open class func registerUser(userMutation: UserMutation, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Register a new user

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let userMutation = UserMutation(username: "username_example", email: "email_example", password: "password_example") // UserMutation | 

// Register a new user
UserAPI.registerUser(userMutation: userMutation) { (response, error) in
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
 **userMutation** | [**UserMutation**](UserMutation.md) |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateUser**
```swift
    open class func updateUser(userMutation: UserMutation, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Update user details

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let userMutation = UserMutation(username: "username_example", email: "email_example", password: "password_example") // UserMutation | 

// Update user details
UserAPI.updateUser(userMutation: userMutation) { (response, error) in
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
 **userMutation** | [**UserMutation**](UserMutation.md) |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

