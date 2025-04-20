#  TODOs

## 🔐 Login / Register

- [X] `LoginView`
- [X] `RegisterView`

## 🙍‍♂️ Profile

- [X] `ProfileView`
    - [X] UserDetails
    - [1/2] edit UserDetails?
    - [X] Logout

## 🚘 Car

- [X] `CarListView`
- [X] `CarDetailView`

## 📅 Booking

- [ ] `MyBookingsView`
- [ ] `BookingDetailView` (use one view for 2 use cases)
    - [ ] to use before making the booking (add button to accept and pay)
    - [ ] to use after booking to look at the details
- [ ] `BookingEditView` (or also in `BookingDetailView` integrated?)

## 🧑‍💼 Admin

- [ ] `AdminUserListView`
    - [ ] list users
    - [ ] search users
    - [ ] delete users? (.swipeAction & .contextMenu)
    - [ ] update users? (.swipeAction & .contextMenu)
- [ ] `AdminCarManagementView`
    - [ ] list cars
    - [ ] search cars
    - [ ] delete cars (.swipeAction & .contextMenu)
    - [ ] edit cars (.swipeAction & .contextMenu)
    - [ ] add cars (ToolbarItem)
- [ ] `AdminBookingManagementView`
    - [ ] list bookings
    - [ ] search bookings
    - [ ] delete bookings (.swipeAction & .contextMenu)
    - [ ] edit bookings (.swipeAction & .contextMenu)

---

## 🛠️ Backend Integration

- [X] Generate `OpenAPIClient`
- [X] Integrate `OpenAPIClient` into the SwiftUI app
- [X] Test API calls
- [X] Add login functionality with `/login`
- [X] Implement token management (store & attach token to requests)
