# Luna Farm Management System

Luna is a system designed to manage agricultural farms, their ownership, cultivation, farmer details, livestock, and related documents, with full audit logging.

---

## 🗂️ Core Entities

### 🏡 Farm
- `Code`: Unique (e.g. `EG/1`, `EG/2`, ...) - auto-incremented using DB-safe mechanism and serves as primary key
- `ArabicName`
- `EnglishName`
- `Engineer`: Responsible supervisor (normally the user who created it)
- `Manager`: One of the owning farmers 
- `Farmers`: List of owning farmers with ownership share (e.g. 4200M²)
- `TotalArea`: Float - in square meter (could be entered as feddan, hectare, etc.)
- `CultivatedArea`: Float
- `YearOfReclamation`: Int (e.g. 2017)
- `OwnershipDocument`: File upload (one or more docs)
- `Locations`: List of parts, each with coordinates , name and description
- `Livestock`: List of animals per type (type, count) - (TODO: future enhancement for livestock) 
- `Workers`: Set of age/gender groups with count
- `Attachments`: General documents (name, file)

---

### 👨‍🌾 Farmer
- `IDNumber`: National ID (14-digit) (validated) (manually entered) - serves as primary key
- `FullName`: At least 4 parts (validated)
- `PersonalImage`: Photo
- `NationalIDImageFront`: Photo of ID front
- `NationalIDImageBack`: Photo of ID back
- `PhoneNumber`: E.164 format recommended

---

### 🤝 Ownership
- `FarmID`
- `FarmerID`
- `Share`: Float (e.g. 25.5%) (validated)
- `IsManager`: Boolean (validated)

---

## 🌱 Cultivation
- `FarmID`
- `CultivatedArea`
- `LastUpdated`

---

## 🐄 Livestock
- `FarmID`
- `Type`: e.g. Cow, Sheep, Chicken
- `Count`: Int

---

## 👥 Workers
- `FarmID`
- `AgeGroup`: e.g. "<18", "18–60", "60+"
- `Gender`: Male / Female
- `Count`

---

## 📎 Attachments
- `FarmID`
- `Name`: Description or title
- `File`: Upload

---

## 📍 Location
- `FarmID`
- `Coordinates`: GeoJSON `Point` or `Polygon`: Textual description or GPS polygon

---

## 🧾 Audit Logs
Tracks all create, update, delete actions.

- `Action`: "CREATE", "UPDATE", "DELETE"
- `Table`: Entity name
- `User`: Who did it
- `Timestamp`
- `OldValue`: JSON (for updates/deletes)
- `NewValue`: JSON (for creates/updates)

---

## ✅ Validation Rules
- Farm code must be unique.
- Farmer full name must have 4+ parts.
- Farmer ID must be unique and valid 14-digit Egyptian number.
- One and only one manager per farm.
- Ownership shares must sum to 100%.
- Cultivated area must not exceed total area.
- Images must be JPG/PNG/PDF (max size configurable).
- Only valid document types are allowed.

---

## 🧱 System Modules

### 1. **Core Data Models**
- Define all structs and DB schema

### 2. **API Services**
- RESTful endpoints for each entity
- Authenticated actions for audit log

### 3. **Audit Logger**
- Middleware logs all changes

### 4. **File Handling**
- Store farmer images, ID documents, ownership docs, and attachments

### 5. **Admin UI**
- Web interface to manage farms, farmers, ownership, and view logs

### 6. **Report Generator**
- Farm reports, demographic summaries, livestock stats

---

## 🔐 Authentication
- Admin login required
- All actions tied to user for audit trail

---

## 🛠️ Development Tasks

1. Design ERD & define Go structs
2. Implement DB schema and migration
3. Build backend API (CRUD, uploads, validation)
4. Integrate audit logger
5. Build admin web UI
6. Add reporting & export functions
7. Seed system with test data
8. Package for deployment (Docker or other)

---

## 🧾 Example Farm Code Generation
- Format: `EG/1`, `EG/2`, ...
- Auto-incremented using DB-safe mechanism

---

## 🔄 Future Enhancements
- GIS mapping integration
- Mobile data collection app
- Role-based permissions
- PDF farm summary generation

---

## 📁 Storage Recommendations
- Database: PostgreSQL
- Files: Local FS or S3-compatible object store

---

## 🔧 Technologies
- Language: Go (Golang)
- Database: PostgreSQL (SQLc)
- UI: HTML/CSS/JS (Vanilla) - for now - later will be Svelte + Tailwind or something similar
- Auth: Simple session/token
