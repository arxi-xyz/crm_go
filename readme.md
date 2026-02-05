# CRM – Functional Specification (MVP)

## 1. Authentication & Roles

### Roles

- SUPER_ADMIN
- ADMIN

### Permissions

- SUPER_ADMIN
    - Create / update / deactivate admins
- ADMIN
    - Login / Logout
    - View own profile
    - Manage users, profiles, and missions

### Authentication

- Admins authenticate via credentials (phone/email + password)
- Authentication is token-based (e.g. JWT)
- Logout invalidates the active session/token

---

## 2. Admin

### Admin Data

- id
- phone or email (unique)
- password_hash
- role (SUPER_ADMIN | ADMIN)
- created_at

### Capabilities

- Login / Logout
- View own profile
- (SUPER_ADMIN only) create new admins

---

## 3. User

### User Data

- id
- phone (unique)
- first_name
- last_name
- profile_image_url

### Notes

- Users are managed by admins
- A user can have one profile

---

## 4. Profile

### Relationship

- One-to-one relationship with User

### Profile Data

- id
- user_id (unique)
- location
- how_known_us
- ranking
- job_title
- updated_at

### Capabilities

- Admins can create, update, and soft delete profiles

---

## 5. Mission

### Mission Data

- id
- title
- user_id
- admin_id
- status
- scheduled_for
- expire_at
- answer_text
- created_at
- answered_at
- deleted_at

### Mission Status

- DRAFT
- SCHEDULED
- SENT
- ANSWERED
- EXPIRED
- CANCELED

### Capabilities

- Admins can create, update, schedule, answer, change status, and soft delete missions

---

## 6. General Rules

- All delete operations are soft delete
- Pagination required for all entities
- Phone must be unique for users
- Profile images stored as URLs