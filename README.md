# AdoptionOS

**AdoptionOS** is an open-source ERP (Enterprise Resource Planning) system designed for animal shelters. It manages the entire lifecycle of animal rescue—from intake and medical records to foster management and public adoption listings.

The platform is designed to run on low-cost hardware (**Raspberry Pi 5**) for on-premise deployments but is cloud-ready for SaaS scaling.

---

## 🏗 Architecture: Modular Monolith

This repository is a **Monorepo** containing both the API and the User Interface.

### 🎨 Frontend (The Face)

The frontend is a **Vue 3** application built with **TypeScript**. It provides a responsive and interactive user experience for adopters, volunteers, and staff.

- **Tech Stack:** Vue 3, TypeScript, TailwindCSS.
- **Key Features:**
  - **Public Portal:** Home, About, Adopt, and Donate pages.
  - **Adoption Process:** Interactive steps for potential adopters.
  - **Volunteer Management:** Digital applications with waiver signing and parental consent logic.
  - **Surrender Forms:** Smart forms for intake management.

### 🧠 Backend (The Brain)

The backend is a high-performance API built with **Go 1.24** (Standard Library) and **PostgreSQL 16**.

- **Tech Stack:** Go, PostgreSQL.
- **Key Features:**
  - **REST API:** Handles all data operations for the frontend.
  - **Data Storage:** PostgreSQL with JSONB for flexible schema requirements (e.g., medical records).
  - **File Management:** Local storage handling for pet images and documents.
  - **Deployment:** Optimized for Raspberry Pi 5 (Production) and macOS (Development).

### Folder Structure

```text
adoption-os/
├── Makefile            # Master commands (Run both FE & BE)
├── backend/            # The "Brain" (Go API)
│   ├── cmd/
│   │   ├── api/        # REST API Handlers, main entry point
│   │   └── seeder/     # Data seeding scripts
│   ├── internal/
│   │   ├── data/       # Database models and access logic
│   │   ├── jsonlog/    # Structured logging
│   │   ├── middleware/ # API middleware (auth, rate limiting)
│   │   ├── models/     # Core data structures
│   │   └── services/   # Business logic for external services (AI, Stripe, etc.)
│   ├── migrations/     # SQL Schema Versioning
│   └── uploads/        # Local file storage for pet images/documents
└── frontend/           # The "Face" (Vue 3 App)
    ├── src/
    │   ├── assets/     # Fonts and static assets
    │   ├── components/ # Reusable Vue components (UI building blocks)
    │   │   ├── about/
    │   │   ├── adopt/
    │   │   ├── common/
    │   │   └── home/
    │   ├── models/     # TypeScript interfaces for data structures
    │   ├── pages/      # Top-level view components for each route
    │   ├── router/     # Vue Router configuration
    │   ├── stores/     # Pinia state management stores
    │   └── styles/     # Global CSS and styling
    ├── public/         # Static assets (favicon, images)
    └── package.json    # UI Dependencies
```

---

## 🚀 Getting Started

Follow these instructions to get the project running on your local machine.

### Prerequisites

- **Go:** Version 1.24 or higher
- **Node.js:** Version 20.x or higher
- **Docker:** For running the PostgreSQL database
- **Make:** To use the simplified commands in the `Makefile`

### Running the Application

1.  **Start the database:**
    The backend requires a running PostgreSQL database. A `docker-compose.yml` file is provided for convenience (though not listed in the initial file structure, it is a common convention). If it exists, you can start the database with:
    ```sh
    docker-compose up -d
    ```
    If not, you will need to manually run a PostgreSQL container. The `backend/Makefile` references a container named `pet_shelter_db`.

2.  **Install Dependencies:**
    This command will install the necessary Go modules for the backend and npm packages for the frontend.
    ```sh
    make setup
    ```

3.  **Run the Backend:**
    This will start the Go API server.
    ```sh
    make run-be
    ```

4.  **Run the Frontend:**
    In a separate terminal, this command will start the Vue development server.
    ```sh
    make run-fe
    ```

Once these steps are complete, the frontend should be accessible at `http://localhost:5173` (or another port specified by Vite) and the backend at `http://localhost:4000`.

---

## 🗺️ Roadmap: Pet Adoption Platform Master Execution Plan

### Year 1: Foundation, Revenue, and Core Operations

#### Phase 1: The Foundation (MVP)
*(Status: Code Complete / In Progress)*
- **Public Website:** Responsive Home, About, Adopt, and Donate pages.
- **Digital Intake:** Smart SurrenderPet forms with risk logic.
- **Volunteer Portal:** Application with "Teen vs. Adult" logic.
- **Backend Core:** Go server & PostgreSQL setup.
- **Admin Shell:** Basic secure login and navigation.

#### Phase 2: Revenue & Reach
- **Stripe Foundation:** Process Adoption Fees & One-Time Donations.
- **Pet Record Management:** Admin CRUD for pet profiles.
- **Manual Donation & Cash Handling:** Log offline donations (cash/check).
- **Digital Kennel Cards (E-Ink V1):** Raspberry Pi integration.
- **Smart QR Codes:** Auto-generate unique pet profile links.

#### Phase 3: Onboarding & Communication
- **Social Media Share Generator:** Create "Adopt Me" graphics.
- **Automated Adopter Onboarding:** Triggered emails with care guides.
- **Foster-to-Adopt Fast-Track:** One-click adoption for fosters.
- **SMS Status Integration:** Application status text alerts.
- **Post-Adoption Surveys:** Automated 3-3-3 check-in emails.

#### Phase 4: Engagement Campaigns
- **External API Sync:** Nightly export to Petfinder/RescueGroups.
- **"Name the Litter":** Paid voting module for naming pets.
- **Digital Raffle Manager:** Sell tickets & pick winners via Stripe.
- **Volunteer Leaderboard:** Gamified tracking of volunteer impact.
- **"Happy Tails":** Public success story submission form.

#### Phase 5: Mobile & Field Operations (PWA)
- **Offline "Field Mode":** Service worker caching for intake.
- **Push Notifications:** Browser alerts for urgent needs.
- **"Add to Home Screen":** Native-app feel configuration.
- **Mobile Image Compression:** Client-side resizing for faster uploads.
- **Geolocation Intake:** GPS tagging for stray intakes.

#### Phase 6: Financial Stability
- **Monthly Subscriptions:** Recurring donor engine.
- **Pet Sponsorships:** Link payments to specific animals.
- **"Gotcha Day" Engine:** Anniversary donation emails.
- **Donor Portal:** Self-service receipt & plan management.
- **Tribute Giving:** "In Memory Of" donation flow.
- **Legacy Giving Portal:** Planned giving & bequests section.

#### Phase 7: Advanced Hardware
- **Intake Kiosk:** Tablet-optimized intake UI.
- **Vet Dashboard:** "Airport Screen" for medical schedules.
- **Device Health Monitor:** Admin view for IoT uptime.
- **NFC Support:** Tap-to-view medical records on collars.
- **Smart Supply Buttons:** IoT buttons for inventory alerts.

#### Phase 8: Data & Inventory
- **Weight Visualization:** Growth charts for neonates.
- **Inventory System:** "Supply Closet" tracking & alerts.
- **Medication Logs:** Strict controlled substance auditing.
- **Volunteer Hours:** Clock-in/out tracking for grants.
- **Incident Reporting:** Liability forms for bites/injuries.

#### Phase 9: Matchmaking & Discovery
- **"Find My Match" Quiz:** Lifestyle-based pet filtering.
- **"Pet Alerts":** Saved search notifications for users.
- **Lifestyle Badges:** Auto-tagging (e.g., "Apartment Friendly").
- **Compare View:** Side-by-side pet comparison tool.
- **Video Embeds:** YouTube/TikTok support in profiles.

#### Phase 10: Community Resources
- **Lost & Found Generator:** Public flyer creation tool.
- **Community Adoption Events Calendar:** Manage off-site adoption events & volunteer shifts.
- **Community Calendar:** Public view of clinics/fundraisers.
- **Partner Directory:** Vets & trainers listing.
- **Newsletter Builder:** Drag-and-drop email tool.

#### Phase 11: Specialized Care & Crisis
- **"SOS" Broadcast:** Emergency SMS to skilled volunteers.
- **Fospice Tracker:** Quality of Life logs for hospice pets.
- **Smart Visiting Hours:** Real-time "Open/Closed" widget.
- **Isolation Management:** Contagion tracking workflows.
- **Disaster Mode:** Crisis homepage takeover switch.

#### Phase 12: Compliance, Security & Standards
- **SOC2 Readiness:** Immutable Audit Logs & Centralized Security Monitoring.
- **GDPR/CCPA Privacy Suite:** "Right to be Forgotten" engine, Data Export (JSON), & Cookie Consent.
- **Strict RBAC Enforcement:** Granular permissioning for Staff vs. Admins (prep for SaaS).
- **Accessibility Deep-Dive:** WCAG 2.1 AA Audit & Screen Reader optimization.
- **Penetration Testing:** Third-party security audit & vulnerability patching.
- **Data Retention:** Automated PII deletion policies.

### Year 2 & Beyond: Scaling to SaaS & Advanced Intelligence

#### Phase 13: SaaS Multi-Tenancy
- **Tenant Logic:** Database sharding by organization_id.
- **Usage Metering:** Track storage/SMS limits per org.
- **Stripe Billing Portal:** SaaS subscription management.
- **Feature Gating:** Code-level permission checks.
- **Setup Wizard:** Automated onboarding for new shelters.

#### Phase 14: Advanced Scoring Logic (Part 1 - Adopters)
- **Adopter Scoring Engine:** Algorithm to score applications.
- **Risk Flags:** Auto-highlight "Red Flags."
- **Application Sorting:** Default sort by "Score."
- **Auto-Decline:** Logic to auto-reject unqualified apps.
- **Reviewer Assignment:** Auto-assign based on workload.

#### Phase 15: Advanced Scoring Logic (Part 2 - Workforce)
- **Volunteer Scoring System:** Rank applicants by skill/availability.
- **Foster Scoring System:** Rate potential fosters.
- **Reliability Index:** Track "Show-up Rate."
- **Capacity Planning:** Forecast volunteer needs.
- **Skill-Gap Analysis:** Identify missing volunteer skills.

#### Phase 16: Waitlist & Demand Management
- **Waitlist System:** Users sign up for breeds.
- **Auto-Notification:** Email waitlist on intake match.
- **Priority Access:** Early access for Sponsors.
- **Demand Heatmap:** View requested breeds vs. intake.
- **"Coming Soon" Teasers:** Blurred profiles for medical hold.

#### Phase 17: Medical Intelligence
- **Medical Reminder System:** Alerts for boosters/meds.
- **Treatment Plans:** Templated medical workflows.
- **Vet Partner Portal:** External vet upload access.
- **Cost-per-Animal Tracking:** Medical spend per pet.
- **Symptom Checker:** Triage tool for fosters.

#### Phase 18: Operational Automation
- **Task Automation:** IFTTT logic for shelter tasks.
- **Document Generator:** Auto-fill PDF contracts.
- **Shift Reminders:** SMS/Email shift reminders.
- **Lost & Found Board:** Internal matching logic.
- **Key Management:** Digital key checkout log.

#### Phase 19: Social & Marketing Automation
- **Social Auto-Poster:** Auto-draft posts for new pets.
- **Story Highlight Generator:** "Weekly Recap" video creator.
- **Ad Integration:** Boost Facebook posts from admin.
- **Influencer Portal:** Media kits for promoters.
- **UTM Tracking:** Track adoption sources.

#### Phase 20: Foster Experience 2.0
- **Foster Portal Dashboard:** Central hub for fosters.
- **Supply Request Flow:** One-click reorders.
- **Foster-to-Foster Chat:** Secure advice channels.
- **Vacation Coverage:** "I'm Away" toggle.
- **Photo Uploader:** Easy mobile photo submission.

#### Phase 21: Data Analytics & Reporting
- **Adoption Funnel Analytics:** Applicant drop-off analysis.
- **Length-of-Stay Reports:** Trends by breed/age.
- **Donor Lifetime Value (LTV):** High-value supporter ID.
- **Impact Reports:** "Year in Review" generator.
- **Heatmaps:** Geo-analysis of adopters/intakes.

#### Phase 22: Integrations (The "Walled Garden")
- **Microchip Registry Sync:** API integration for registration.
- **Internal Accounting Ledger:** Track income/expenses securely.
- **Internal Email System:** Marketing campaigns without 3rd parties.
- **Internal Chat System:** Secure team communication.
- **Internal Calendar System:** Shifts/Events management.

#### Phase 23: AI & Machine Learning
- **AI Bio Generator:** Creative text from tags.
- **Photo Quality Scorer:** Rate photos for adoption appeal.
- **Churn Prediction:** Identify at-risk volunteers.
- **Breed Identification:** Computer vision breed estimates.
- **Sentiment Analysis:** Analyze feedback for issues.

#### Phase 24: Enterprise Expansion
- **Multi-Location Support:** Manage multiple sites.
- **Franchise Model:** Parent/Child org structures.
- **API Access:** Open API for enterprise tools.
- **White Labeling:** Branding removal.
- **Global Localization:** Multi-language/currency support.