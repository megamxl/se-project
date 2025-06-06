//
// UserMutation.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation
#if canImport(AnyCodable)
import AnyCodable
#endif

@available(*, deprecated, renamed: "OpenAPIClientAPI.UserMutation")
public typealias UserMutation = OpenAPIClientAPI.UserMutation

extension OpenAPIClientAPI {

/** The escape-room instance to join */
public struct UserMutation: Codable, JSONEncodable, Hashable {

    public var username: String?
    public var email: String?
    public var password: String?

    public init(username: String? = nil, email: String? = nil, password: String? = nil) {
        self.username = username
        self.email = email
        self.password = password
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case username
        case email
        case password
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(username, forKey: .username)
        try container.encodeIfPresent(email, forKey: .email)
        try container.encodeIfPresent(password, forKey: .password)
    }
}

}
