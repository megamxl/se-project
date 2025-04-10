//
//  Font+Extension.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 08.04.25.
//

import SwiftUI

extension View {
    func faDuotoneLight(size: CGFloat) -> some View {
        self.font(.custom("FontAwesome6Duotone-Light", size: size))
    }
    func faDuotoneRegular(size: CGFloat) -> some View {
        self.font(.custom("FontAwesome6Duotone-Regular", size: size))
    }
    func faDuotoneSolid(size: CGFloat) -> some View {
        self.font(.custom("FontAwesome6Duotone-Solid", size: size))
    }
    func faDuotoneThin(size: CGFloat) -> some View {
        self.font(.custom("FontAwesome6Duotone-Thin", size: size))
    }
}
