# 🛣️ Luna Development Roadmap

A structured roadmap to guide the development of Luna — a comprehensive, auditable farm management system.

---

## 🔧 Technologies

- **Language**: Go (Golang)
- **Database**: PostgreSQL (via `sqlc`)
- **UI**: HTML/CSS/JS (Vanilla JS initially, future migration to Svelte + Tailwind)
- **Auth**: Simple session/token-based authentication
- **File Storage**: Local filesystem (can be abstracted for future object store)

---

## 📆 Milestones & Timeline

### 🧱 Milestone 1: System Planning & Architecture (Week 1)
- [x] Define system scope and functional requirements
- [x] Write Markdown spec and core models
- [X] Draw ERD and identify table relationships
- [X] Design Go struct layout reflecting DB schema

**Deliverables:**
- ERD (image + .mwb/.dbml/.mermaid)
- Initial Go struct definitions
- Luna system spec (done)

---

### 🗃️ Milestone 2: Database Schema & Codegen (Week 2–3)
- [X] Create PostgreSQL schema
- [X] Set test DB up 
- [X] Implement Basic CRUD for models
  - [X] farms
  - [X] farmers
  - [X] livestock
  - [X] workers
  - [X] attachments
  - [X] locations
- [ ] Seed scripts for dev/testing

**Deliverables:**
- SQL schema files
- Sample seed data

---

### 🧠 Milestone 3: Core Backend API (Week 4–6)
- [ ] Build REST API (Fiber or Chi preferred)
- [ ] CRUD for:
  - Farms
  - Farmers
  - Ownerships
  - Locations
  - Livestock
  - Demographics
  - Attachments
- [ ] Cultivated area tracker
- [ ] File/image upload endpoints

**Deliverables:**
- Functional backend with DB integration
- Unit tests for core API
- Swagger/OpenAPI docs (optional)

---

### 📝 Milestone 4: Validation & Business Rules (Week 6–7)
- [ ] Enforce ownership rules (100%, one manager)
- [ ] Validate farmer full name format (4 parts min)
- [ ] Validate cultivated vs total area
- [ ] File type/size validation

**Deliverables:**
- Validation service/middleware
- Error handling responses

---

### 🔐 Milestone 5: Authentication & Audit Logging (Week 7–8)
- [ ] Implement user model + login (hardcoded or minimal DB)
- [ ] Session or JWT-based auth
- [ ] Build audit logging middleware
  - Who did what, when, on what
  - Log changes on create/update/delete

**Deliverables:**
- Auth middleware
- Audit logs DB and service
- Secure routes for backend

---

### 🧭 Milestone 6: Admin UI v1 (Week 9–10)
- [ ] Basic HTML/CSS/JS UI:
  - Farm list, create/edit
  - Farmer profiles
  - Ownership assigner
  - Image/document upload
  - Audit log viewer
- [ ] Responsive layout (Bootstrap or Tailwind CDN)

**Deliverables:**
- Static admin interface
- JS calling REST APIs
- Authenticated session support

---

### 📊 Milestone 7: Reports & Stats (Week 11–12)
- [ ] Farm summary reports
- [ ] Livestock & demographics summaries
- [ ] Export to CSV
- [ ] Cultivation vs total area graphs

**Deliverables:**
- JSON API + optional charting
- Download/export endpoint

---

### 📦 Milestone 8: Deployment & Packaging (Week 13)
- [ ] Dockerfile and docker-compose (PostgreSQL + app)
- [ ] File storage abstraction
- [ ] Backup scripts (DB & file assets)

**Deliverables:**
- Production-ready build
- README with setup instructions
- Local dev env setup (`make dev` or similar)

---

## 🌿 Future Enhancements (Post v1.0)

- [ ] GIS farm location map
- [ ] Role-based permissions (admin, engineer, auditor)
- [ ] Mobile-first frontend
- [ ] PDF export (farm certificate, summaries)
- [ ] Svelte + Tailwind UI migration
- [ ] Object storage (S3-compatible)
- [ ] Offline-first data collection app

---

## 📌 Weekly Checkpoints

| Week | Goal                                  | Status     |
|------|----------------------------------------|------------|
| 1    | Planning, ERD, Spec                    | ✅ In progress |
| 2–3  | DB Schema, SQLc Integration            | ⏳          |
| 4–6  | Backend CRUD + File Upload             | ⏳          |
| 6–7  | Validation + Rule Engine               | ⏳          |
| 7–8  | Auth + Audit Logging                   | ⏳          |
| 9–10 | Admin UI (Vanilla JS)                  | ⏳          |
| 11–12| Reports + Export                       | ⏳          |
| 13   | Docker + Deployment                    | ⏳          |

---

## 🧰 Tools & Utilities



- **Database**: PostgreSQL
- **Migrations**: `golang-migrate`
- **Codegen**: `sqlc`
- **Framework**: `Chi`, `Fiber`, or standard `net/http`
- **Frontend**: Vanilla JS (HTML templates)
- **Auth**: Sessions or JWT (simple user model)
- **Testing**: `Testify`, `httptest`, `sqlmock`