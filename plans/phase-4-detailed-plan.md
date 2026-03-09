# Phase 4: Financial & Compliance - Detailed Implementation Plan

## Overview

Phase 4 implements the financial backbone and regulatory compliance layer: Payment processing, KYC verification, AML detection, Fraud prevention, Risk scoring, and the Bonus/Promotion system.

**Prerequisites**: Phase 3 complete (Multiplayer, Poker, Tournaments, Notifications).

---

## 1. Payment Service (Java Spring Boot)

### 1.1 Project Setup
- Spring Boot with Spring Web, Spring Data JPA, Spring Security
- gRPC server (port 9012), PostgreSQL (casino_payments), Redis, NATS
- PCI DSS considerations: No raw card data stored, tokenization only

### 1.2 Database Migrations (casino_payments)

- `payment_methods` - Player saved payment methods (tokenized)
- `deposits` - Deposit requests and status tracking
- `withdrawals` - Withdrawal requests with approval workflow
- `gateway_transactions` - Raw gateway response logs
- `reconciliation_records` - Daily reconciliation entries
- `chargeback_records` - Chargeback tracking

### 1.3 Payment Gateway Integrations

| Gateway | Type | Implementation |
|---------|------|---------------|
| Stripe | Cards | Stripe SDK, Payment Intents API |
| Adyen | Cards + Local | Adyen Drop-in, multi-method |
| Skrill | E-wallet | REST API integration |
| Neteller | E-wallet | REST API integration |
| Coinbase Commerce | Crypto | Webhook-based confirmation |
| BitPay | Crypto | Invoice API |
| Paysafecard | Prepaid | REST API |
| Local Bank Transfer | Bank | Country-specific APIs |

- **Adapter pattern**: Each gateway implements `PaymentGatewayAdapter` interface
- **Gateway selection**: Configurable per currency, country, and payment method
- **Failover**: If primary gateway fails, route to secondary

### 1.4 Deposit Flow
1. Player selects payment method and amount
2. Validate: min/max limits, KYC level requirements, daily limits
3. Create deposit record (status: initiated)
4. Redirect to gateway / process tokenized payment
5. Receive webhook/callback from gateway
6. Verify webhook signature
7. Update deposit status (completed/failed)
8. Call Wallet Service `Credit` gRPC
9. Publish `financial.events.deposit.completed` to NATS
10. Trigger notification

### 1.5 Withdrawal Flow
1. Player requests withdrawal (amount, method)
2. Validate: min/max limits, KYC verified, wagering requirements met
3. Check AML: Call Risk Scoring Service for risk assessment
4. Create withdrawal record (status: pending_review)
5. **Auto-approve** if: amount < threshold AND risk_score < medium AND KYC level >= 2
6. **Manual review** if: amount > threshold OR risk_score >= medium OR first withdrawal
7. Admin approves/rejects in admin panel
8. On approval: Call Wallet Service `Debit`, initiate gateway payout
9. Track payout status from gateway
10. Publish `financial.events.withdrawal.completed` to NATS

### 1.6 Reconciliation
- Daily automated job:
  - Fetch all gateway transactions for the day
  - Compare with internal deposit/withdrawal records
  - Flag discrepancies for manual review
  - Generate reconciliation report
- Monthly settlement reports per gateway

### 1.7 Chargeback Handling
- Webhook from gateway on chargeback
- Auto-suspend player account
- Create chargeback record
- Debit disputed amount from wallet (if balance available)
- Alert compliance team
- Track dispute resolution

---

## 2. KYC Service (Java Spring Boot)

### 2.1 Project Setup
- Spring Boot, gRPC (port 9013), PostgreSQL (casino_users), S3 (documents)

### 2.2 Third-Party Integration
- **Primary**: Sumsub (Sum&Substance) - Full KYC/AML provider
- **Fallback**: Onfido
- **Adapter pattern**: `KYCProviderAdapter` interface

### 2.3 Verification Flows

#### Level 1: Basic Verification
- Email verification (handled by Auth Service)
- Phone verification (SMS OTP)
- Basic personal info (name, DOB, country)
- **Auto-verified** via email + phone

#### Level 2: Identity Verification
- Government ID upload (passport, national ID, driver license)
- OCR extraction of document data
- Document authenticity check (via Sumsub)
- Face match: Selfie vs document photo
- **Processing**: Auto-verified by provider, manual review for edge cases

#### Level 3: Enhanced Due Diligence
- Proof of address (utility bill, bank statement < 3 months)
- Source of funds declaration
- Source of wealth documentation (for high-value players)
- **Processing**: Always manual review by compliance team

### 2.4 KYC Triggers
- First withdrawal attempt → Require Level 2
- Cumulative deposits > $2,000 → Require Level 2
- Cumulative deposits > $10,000 → Require Level 3
- Risk score elevated → Require Level 3
- Admin manual request

### 2.5 Document Management
- Upload to S3 with encryption (AES-256)
- Pre-signed URLs for admin viewing (time-limited)
- Retention policy: 5 years after account closure
- GDPR: Right to erasure (anonymize, retain for compliance period)

---

## 3. AML Detection Service (Python FastAPI)

### 3.1 Project Setup
- FastAPI with async support
- gRPC server (port 9014)
- PostgreSQL (casino_compliance), pgvector, TimescaleDB, Redis, NATS
- ML libraries: scikit-learn, XGBoost, pandas

### 3.2 Rules Engine

#### 3.2.1 Transaction Rules
| Rule | Description | Threshold | Action |
|------|-------------|-----------|--------|
| Structuring | Multiple deposits just below reporting threshold | 3+ deposits within 24h totaling > $10K | Alert |
| Rapid Deposit-Withdraw | Deposit then withdraw with minimal play | < 3x wagering of deposit | Alert |
| Large Transaction | Single large transaction | > $10,000 | CTR Report |
| Velocity | Unusual transaction frequency | > 10 transactions/hour | Alert |
| Round-tripping | Deposit → minimal play → withdraw to different method | Pattern match | Alert + Block |
| Third-party | Deposit from non-matching name | Name mismatch detection | Block + Alert |

#### 3.2.2 Behavioral Rules
| Rule | Description | Action |
|------|-------------|--------|
| Chip Dumping | Intentional losing in poker to transfer funds | Alert |
| Minimal Play | Very low wagering relative to deposits | Alert |
| Pattern Change | Sudden change in betting patterns | Enhanced monitoring |
| Multi-account | Same device/IP with multiple accounts | Alert + Investigate |
| Geographic Anomaly | Transactions from high-risk jurisdictions | Enhanced monitoring |

### 3.3 ML Risk Model

- **Feature engineering**:
  - Transaction velocity (deposits/withdrawals per period)
  - Average transaction amount
  - Wagering-to-deposit ratio
  - Game diversity score
  - Session duration patterns
  - Payment method diversity
  - Geographic consistency
- **Model**: XGBoost ensemble classifier
  - Training data: Historical flagged/cleared cases
  - Output: Risk score 0-100
  - Categories: Low (0-25), Medium (26-50), High (51-75), Critical (76-100)
- **Player behavior embeddings** (pgvector):
  - 768-dimensional vectors representing player behavior patterns
  - Updated daily from TimescaleDB event data
  - Similarity search to find players matching known fraud patterns

### 3.4 Alert Management
- Alert lifecycle: Generated → Assigned → Investigating → Resolved/Escalated
- Auto-assign to compliance team based on alert type
- SLA tracking: High alerts must be reviewed within 24 hours
- Case management: Group related alerts into cases

### 3.5 Regulatory Reporting
- **CTR (Currency Transaction Report)**: Auto-generate for transactions > $10K
- **SAR (Suspicious Activity Report)**: Generate from investigated alerts
- **Export formats**: FinCEN BSA format, local regulatory formats
- **Audit trail**: All report submissions logged

---

## 4. Fraud Detection Service (Python FastAPI)

### 4.1 Project Setup
- FastAPI, gRPC (port 9015), PostgreSQL, pgvector, Redis, NATS

### 4.2 Multi-Account Detection
- **Device fingerprinting**: Canvas, WebGL, audio context, fonts, screen resolution
- **IP correlation**: Track IP addresses per account, flag shared IPs
- **Email pattern**: Detect email variations (john+1@, j.o.h.n@)
- **Behavioral clustering**: pgvector similarity search on behavior embeddings
- **Graph analysis**: Build player relationship graph, detect clusters

### 4.3 Bot Detection
- **Behavioral signals**:
  - Consistent action timing (humans have variance)
  - No mouse movement patterns (mobile: no touch patterns)
  - Perfect play patterns (always optimal strategy)
  - 24/7 activity without breaks
- **CAPTCHA triggers**: On suspicious behavior score
- **Device attestation**: Google Play Integrity / Apple App Attest

### 4.4 Collusion Detection (Poker)
- **Signals**:
  - Players frequently at same tables
  - Asymmetric chip flow between specific players
  - Coordinated fold/raise patterns
  - Shared IP/device between table players
- **Analysis**: Run after each poker session, batch analysis nightly
- **Action**: Flag for manual review, auto-separate if confidence high

### 4.5 Real-Time Fraud Scoring
- **Latency target**: < 100ms per transaction
- **Scoring pipeline**:
  1. Extract features from transaction + player history (Redis cache)
  2. Run through rules engine (fast path)
  3. Run through ML model (if rules inconclusive)
  4. Return composite score
- **Actions based on score**:
  - 0-25: Allow
  - 26-50: Allow + enhanced monitoring
  - 51-75: Allow + require additional verification
  - 76-100: Block + alert

---

## 5. Risk Scoring Service (Python FastAPI)

### 5.1 Project Setup
- FastAPI, gRPC (port 9016), PostgreSQL, Redis

### 5.2 Unified Risk Profile
- Aggregate signals from AML, Fraud, KYC, and transaction history
- **Risk factors** (weighted):
  - KYC level (lower = higher risk)
  - AML alert history
  - Fraud score history
  - Transaction patterns
  - Device/location changes
  - Account age
  - VIP level (higher VIP = lower risk weight)
- **Dynamic limits**: Auto-adjust deposit/withdrawal limits based on risk
- **Risk categories**: Low, Medium, High, Critical
- **Automated actions**:
  - Medium: Enhanced monitoring, lower withdrawal limits
  - High: Require additional KYC, manual withdrawal review
  - Critical: Account suspension, all withdrawals blocked

---

## 6. Bonus/Promotion Service (Java Spring Boot)

### 6.1 Project Setup
- Spring Boot, gRPC (port 9017), PostgreSQL (casino_bonuses), Redis, NATS

### 6.2 Database Migrations (casino_bonuses)

- `bonus_campaigns` - Campaign definitions
- `bonus_rules` - Eligibility and wagering rules
- `player_bonuses` - Awarded bonuses per player
- `wagering_progress` - Wagering requirement tracking
- `free_spin_awards` - Free spin bonus tracking
- `bonus_transactions` - Bonus-related wallet transactions

### 6.3 Bonus Types Implementation

#### Welcome Bonus (First Deposit Match)
- Config: Match percentage (100%), max bonus ($500), min deposit ($20)
- Wagering: 30x bonus amount on eligible games
- Expiry: 30 days from award
- Game weights: Slots 100%, Blackjack 10%, Roulette 20%

#### Reload Bonus
- Config: Match percentage (50%), max bonus ($200)
- Trigger: Specific deposit (2nd, 3rd) or promotional period
- Wagering: 25x bonus amount

#### No-Deposit Bonus
- Config: Fixed amount ($10) or free spins (50)
- Trigger: Registration, promotional code
- Wagering: 40x bonus amount
- Max cashout: 5x bonus amount

#### Free Spins
- Config: Number of spins, specific game, bet value per spin
- Trigger: Deposit, promotion, loyalty reward
- Winnings credited as bonus balance with wagering requirements

#### Cashback
- Config: Percentage of net losses (10%), period (weekly)
- Calculation: Net losses = total bets - total wins
- Credit: Real money (no wagering) or bonus money (with wagering)
- VIP-tier specific rates

#### Referral Bonus
- Config: Bonus for referrer and referee
- Trigger: Referee makes first deposit
- Wagering: Standard wagering requirements

### 6.4 Wagering Requirement Engine
- Track every bet placed against wagering progress
- Game contribution weights (configurable per campaign)
- Progress calculation: `progress = sum(bet_amount * game_weight) / required_amount`
- On completion: Convert bonus balance to real balance
- On expiry: Forfeit remaining bonus balance and winnings

### 6.5 Bonus Abuse Detection
- Multi-account bonus claiming
- Bonus hunting patterns (deposit → claim → minimal play → withdraw)
- Excessive bonus-to-deposit ratio
- Integration with Fraud Detection Service

---

## 7. Admin Panel Extensions

### 7.1 Payment Management
- Deposit list with status filters and search
- Withdrawal approval queue with risk indicators
- Manual adjustment form with approval workflow
- Gateway health monitoring
- Reconciliation report viewer

### 7.2 KYC Review Interface
- Document review queue (pending verifications)
- Side-by-side document viewer (ID + selfie)
- Approve/reject with notes
- Request additional documents
- KYC statistics dashboard

### 7.3 Compliance Dashboard
- AML alerts list with severity indicators
- Alert detail view with transaction timeline
- Case management interface
- SAR filing workflow
- Risk score overview (player distribution by risk level)

### 7.4 Bonus Management
- Campaign creation wizard
- Active campaigns list
- Player bonus tracking (search by player)
- Wagering progress monitoring
- Bonus performance analytics (cost, conversion, ROI)

---

## Phase 4 Completion Criteria

- [ ] Payment Service processing deposits via Stripe and at least one other gateway
- [ ] Withdrawal flow with auto-approve and manual review working
- [ ] Daily reconciliation job running successfully
- [ ] KYC Service verifying documents via Sumsub integration
- [ ] KYC auto-triggers firing at configured thresholds
- [ ] AML rules engine detecting structuring and rapid deposit-withdraw patterns
- [ ] ML risk model scoring players with < 100ms latency
- [ ] Fraud detection identifying multi-account and bot patterns
- [ ] Risk scoring aggregating all signals into unified risk profile
- [ ] Bonus system awarding welcome, reload, free spins, and cashback bonuses
- [ ] Wagering requirement engine tracking progress correctly
- [ ] Admin panel: withdrawal approval, KYC review, compliance dashboard, bonus management
- [ ] All financial transactions have complete audit trail
- [ ] PCI DSS compliance: no raw card data stored anywhere
