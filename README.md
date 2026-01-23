# AdoptionOS

**AdoptionOS** is an open-source ERP (Enterprise Resource Planning) system designed for animal shelters. It manages the entire lifecycle of animal rescue—from intake and medical records to foster management and public adoption listings.

The platform is designed to run on **low-cost, energy-efficient hardware** for on-premise deployments but is cloud-ready for SaaS scaling.

---

## 🏗 Architecture: Modular Monolith

This repository is a **Monorepo** containing both the API and the User Interface.

### 🎨 Frontend

The frontend is a **Vue 3** application built with **TypeScript**. It provides a responsive and interactive user experience for adopters, volunteers, and staff.

- **Tech Stack:** Vue 3, TypeScript.
- **Key Features:**
  - **Public Portal:** Home, About, Adopt, and Donate pages.
  - **Adoption Process:** Interactive steps for potential adopters.
  - **Volunteer Management:** Digital applications with waiver signing and parental consent logic.
  - **Surrender Forms:** Smart forms for intake management.

### 🧠 Backend

The backend is a high-performance API built with **Go (Latest Stable)** (Standard Library) and **PostgreSQL (Latest Stable)**.

- **Tech Stack:** Go, PostgreSQL.
- **Key Features:**
  - **REST API:** Handles all data operations for the frontend.
  - **Data Storage:** PostgreSQL with JSONB for flexible schema requirements (e.g., medical records).
  - **Asset Management:** Secure handling for pet images and documents.
  - **Deployment:** Optimized for energy-efficient hardware (Production) and macOS (Development).

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
└── firmware/           # TinyGo Firmware for Pico 2 W E-ink Displays
```

---

## 🚀 Getting Started

Follow these instructions to get the project running on your local machine.

### Prerequisites

- **Go:** Latest Stable Version
- **Node.js:** Latest Stable Version
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

> [!CAUTION]
> **Security Warning**: Ensure you create a `.env` file for local configuration. **NEVER** commit secrets or API keys to the repository.

---

## �️ Security

If you discover a security vulnerability, please do not open a public issue. Email security@adoption-os.org (placeholder) instead.

---

## �🗺️ Roadmap: Adoption OS Master Execution Plan

### Year 1: Foundation, Revenue, and Core Operations

#### Phase 1: Core Infrastructure & MVP Launch

_(Status: In Progress)_

- [x] **Public Website:** Build the responsive marketing front-end (Home, About, Adopt, Donate) using Vue 3 and Tailwind, ensuring fast load times on mobile.
- [x] **Digital Intake:** Create a multi-step "Surrender Pet" wizard that flags high-risk intakes (e.g., "Bite History" or "Medical Emergency") for immediate staff review.
- [x] **Volunteer Portal:** Develop a gated portal for volunteers to apply, utilizing conditional logic to serve different waiver forms for minors (Teen) vs. Adults.
- [x] **Adoption Application:** Implement the primary long-form adoption questionnaire with real-time validation to prevent incomplete submissions.
- [ ] **Backend Core:** Initialize the Go server, set up the PostgreSQL schema with migrations, and configure environment-based configuration (Dev vs. Prod).

#### Phase 2: Agency Administration & Workflow

- [ ] **Admin Shell:** Build the secure internal layout including secure token-based authentication, role-based route guards, and a responsive sidebar navigation for staff.
- [ ] **Admin Dashboard:** Create a "Morning Overview" widget showing pending applications, urgent medical tasks, and recent donations at a glance.
- [ ] **Pet Record Management:** Develop full CRUD (Create, Read, Update, Delete) views for animal profiles, including status toggles (Available/Pending/Adopted).
- [ ] **Mobile Image Compression:** Implement client-side image resizing (Canvas API) to auto-convert 10MB phone photos to optimized WebP thumbnails _before_ upload.
- [ ] **External API Sync:** Build a scheduled job to securely sync pet data with Petfinder and RescueGroups API standards.
- [ ] **Volunteer Hours:** Create a digital "Time Clock" for volunteers to log start/stop times, generating reports needed for grant compliance.
- [ ] **Intake Kiosk:** Design a simplified, high-contrast "Kiosk Mode" view for tablets used at the front desk for rapid animal intake.

#### Phase 3: Transactions, Communication & Automation

- [ ] **Stripe Foundation:** Integrate Stripe PaymentIntents for handling secure credit card processing for adoption fees and one-off donations.
- [ ] **Post-Adoption Surveys:** Set up a background worker to schedule automated "Check-in" emails at 3 days, 3 weeks, and 3 months post-adoption.
- [ ] **Automated Adopter Onboarding:** Create a trigger system that sends specific PDF care guides (e.g., "Kitten 101" vs. "Senior Dog Care") based on the adopted animal's attributes.
- [ ] **Vet Dashboard:** Build a "Read-Only" display mode optimized for TV screens to show the daily surgery and medication schedule in the medical room.
- [ ] **Manual Donation & Cash Handling:** Add an administrative form to manually log cash/check donations so the financial reporting remains accurate offline.
- [ ] **Foster-to-Adopt Fast-Track:** Build a simplified workflow that allows current fosters to finalize an adoption without re-entering their personal data.
- [ ] **SMS Status Integration:** Connect Brevo/Twilio to send automated text alerts (e.g., "Application Approved") to reduce follow-up phone calls.
- [x] **Smart QR Codes:** Implement a generator that creates unique QR codes linking directly to a pet’s profile, printable for kennel cards.
- [x] **Digital Kennel Cards (E-Ink V2):** Develop the Go script for IoT devices to fetch pet details via API and render them to an e-ink display.

#### Phase 4: PWA & Mobile Capabilities

- [ ] **Offline "Field Mode":** Configure Service Workers and IndexedDB to allow staff to fill out intake forms in areas with zero cell service, syncing when connection is restored.
- [ ] **Push Notifications:** Implement the Web Push API to send browser-based alerts to volunteers for urgent needs (e.g., "Transporter needed ASAP").
- [ ] **"Add to Home Screen":** Configure the `manifest.json` to allow users to install the website as a standalone app on iOS/Android.
- [ ] **Geolocation Intake:** Use the browser Geolocation API to automatically tag the GPS coordinates of stray animal intakes for mapping data.

#### Phase 5: Gamification & Fundraising Campaigns

- [ ] **"Name the Litter":** Build a paid voting module where donors pay $5 to submit or vote on names for a new litter of kittens.
- [ ] **Digital Raffle Manager:** Create a compliant raffle system that sells virtual tickets via Stripe and uses a cryptographically secure random number generator to pick winners.
- [ ] **Volunteer Leaderboard:** Gamify volunteer retention by displaying "Top Hours This Month" and awarding digital badges for milestones.
- [ ] **"Happy Tails":** Build a public submission form allowing past adopters to upload photos and stories, which enter a moderation queue before publishing.

#### Phase 6: Donor Stewardship & Recurring Revenue

- [ ] **Monthly Subscriptions:** Implement Stripe Subscriptions to handle recurring "Guardian Angel" monthly giving plans.
- [ ] **Pet Sponsorships:** Create a logic flow that links a specific donation to a specific Pet ID, showing a "Sponsored by [Name]" badge on the pet's profile.
- [ ] **"Gotcha Day" Engine:** A nightly job that checks adoption dates and sends "Happy 1 Year Anniversary" emails to past adopters with a donation prompt.
- [ ] **Donor Portal:** A self-service dashboard where donors can download year-end tax receipts and update their credit card information.
- [ ] **Tribute Giving:** Add a dedicated flow for "In Memory Of" / "In Honor Of" donations that sends a notification card to the honoree.
- [ ] **Legacy Giving Portal:** Create informational pages and a lead-generation form for users interested in including the shelter in their will/estate.

#### Phase 7: Enterprise Security & Governance

- [ ] **Smart Visiting Hours:** A widget that auto-updates the website header based on current time and "Holiday Mode" overrides set by admins.
- [ ] **SOC2 Readiness:** Implement immutable audit logging (who clicked what, when) and centralized security monitoring for all admin actions.
- [ ] **GDPR/CCPA Privacy Suite:** Build tools for users to request "Export My Data" (JSON dump) or "Delete My Data" to comply with privacy laws.
- [ ] **Strict RBAC Enforcement:** Refactor permission logic to support granular roles (e.g., "Volunteer Leader" can edit shifts but not pet medical records).
- [ ] **Accessibility Deep-Dive:** Conduct a full audit to ensure all forms and navigation meet WCAG 2.1 AA standards for screen readers and keyboard navigation.
- [ ] **Penetration Testing:** Coordinate a third-party security audit to stress-test the Go API against SQL injection and XSS attacks.
- [ ] **Data Retention:** Create automated cron jobs to anonymize or delete sensitive user data (like old applications) after a set retention period (e.g., 2 years).

#### Phase 8: Clinical Health & Inventory Operations

- [ ] **Medication Logs:** Build a strict digital ledger for veterinary pharmaceuticals, requiring specific user permissions to log dispensing.
- [ ] **Weight Visualization:** Integrate Chart.js to visualize kitten weight gain over time, flagging animals that are losing weight.
- [ ] **Incident Reporting:** Create a legal-grade form for documenting bites or accidents, generating PDFs for insurance purposes.
- [ ] **Inventory System:** A "Supply Closet" tracker that alerts admins when critical items (like vaccines or kitten formula) drop below par levels.

#### Phase 9: Adopter Matchmaking & Marketing

- [ ] **"Find My Match" Quiz:** A logic-based quiz that scores available pets against user lifestyle inputs (activity level, kids, other pets).
- [ ] **"Pet Alerts":** Allow users to subscribe to saved search criteria (e.g., "Female Golden Retriever") and receive email alerts when a match is added.
- [ ] **Lifestyle Badges:** Auto-tag pets with badges like "Apartment Friendly" or "Good with Cats" based on their behavioral assessment data.
- [ ] **Compare View:** A side-by-side interface allowing users to pin 2-3 pets and compare their stats (age, size, energy) on one screen.
- [ ] **Social Media Share Generator:** A tool that generates a dynamic Open Graph image (with pet photo and stats) for beautiful link previews on Facebook/Twitter.
- [ ] **Video Embeds:** Update the media gallery to support embedding YouTube or TikTok URLs alongside static images.

#### Phase 10: Community Outreach & Event Management

- [ ] **Lost & Found Generator:** A tool where the public can input details of a lost pet and instantly generate a printable PDF flyer.
- [ ] **Community Adoption Events Calendar:** An internal scheduling tool to assign specific animals and volunteers to off-site events (e.g., "PetSmart Saturday").
- [ ] **Community Calendar:** A public-facing calendar component displaying vaccination clinics, fundraisers, and orientation dates.
- [ ] **Partner Directory:** A searchable directory of recommended local vets, trainers, and pet sitters managed by the shelter.
- [ ] **Newsletter Builder:** A drag-and-drop email composition tool that pulls live pet data directly from the database into the newsletter layout.

#### Phase 11: IoT & Connected Hardware

- [ ] **Device Health Monitor:** A dashboard view showing the "Heartbeat" status, CPU temp, and connection quality of all deployed IoT devices.
- [ ] **NFC Support:** Update the mobile app to read NFC tags on kennel cards or collars to instantly pull up the medical record on a staff phone.
- [ ] **Smart Supply Buttons:** Integration with physical IoT buttons (like AWS IoT buttons) to trigger "Restock Needed" alerts for laundry or food.

#### Phase 12: Crisis Response & Specialized Protocols

- [ ] **"SOS" Broadcast:** A "Break Glass" feature that instantly sends SMS/Email blasts to a pre-vetted list of emergency fosters during overcrowding.
- [ ] **Fospice Tracker:** Specialized quality-of-life tracking logs for hospice pets, prioritizing comfort metrics over medical cure.
- [ ] **Isolation Management:** A workflow to digitally "Lock" a pet profile if they are suspected of having a contagious disease (like Parvo), preventing accidental movement.
- [ ] **Disaster Mode:** A master switch that replaces the homepage with critical emergency info (e.g., "Evacuation in Progress") and disables non-essential forms.

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

- **Volunteer Matching Logic:** Rank applicants by skill/availability.
- **Foster Matching Logic:** Rate potential fosters.
- **Engagement Metrics:** Track "Show-up Rate."
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

- **Task Automation:** Event-driven workflows for shelter tasks.
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
- **Retention Analytics:** Identify at-risk volunteers.
- **Breed Identification:** Computer vision breed estimates.
- **Sentiment Analysis:** Analyze feedback for issues.

#### Phase 24: Enterprise Expansion

- **Multi-Location Support:** Manage multiple sites.
- **Franchise Model:** Parent/Child org structures.
- **API Access:** Open API for enterprise tools.
- **White Labeling:** Branding removal.
- **Global Localization:** Multi-language/currency support.
