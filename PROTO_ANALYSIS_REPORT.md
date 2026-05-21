# Proto File Analysis Report

## Executive Summary

This report identifies **duplicate proto files** across the codebase and provides recommendations for consolidation. All proto files should be maintained in `/workspace/services/common-service/proto/` only.

---

## 1. Current Proto File Distribution

### 1.1 Master Proto Files (CORRECT LOCATION)
Location: `/workspace/services/common-service/proto/`

**20 proto files found:**
- `common/v1/` (5 files): enums.proto, error.proto, money.proto, pagination.proto, timestamp.proto
- `affiliate/v1/`: affiliate_service.proto
- `agent/v1/`: agent_service.proto
- `auth/v1/`: auth_service.proto
- `bonus/v1/`: bonus_service.proto
- `commission/v1/`: commission_service.proto
- `game/v1/`: game_registry.proto
- `jackpot/v1/`: jackpot_service.proto
- `leaderboard/v1/`: leaderboard_service.proto
- `merchant/v1/`: merchant_service.proto
- `payment/v1/`: payment_service.proto
- `risk/v1/`: risk_service.proto
- `tournament/v1/`: tournament_service.proto
- `user/v1/`: user_service.proto
- `wallet/v1/`: wallet_service.proto
- `winners/v1/`: winners_service.proto

### 1.2 Duplicate Proto Files (NEEDS REMOVAL/MIGRATION)

#### A. Legacy Root-Level Protos (`/workspace/proto/`)
- ❌ `/workspace/proto/auth/auth.proto` - DUPLICATE of `common-service/proto/auth/v1/auth_service.proto`
- ❌ `/workspace/proto/wallet/wallet.proto` - DUPLICATE of `common-service/proto/wallet/v1/wallet_service.proto`
- ❌ `/workspace/proto/common/common.proto` - DUPLICATE of common protos in `common-service/proto/common/v1/`

#### B. Service-Specific Duplicates

**Affiliate Service:**
- ❌ `/workspace/services/affiliate-service/src/main/proto/affiliate_service.proto`
  - DUPLICATE of `common-service/proto/affiliate/v1/affiliate_service.proto`
  - Status: Older version, missing enhanced features

**Bonus Service:**
- ❌ `/workspace/services/bonus-service/src/main/proto/bonus_service.proto`
  - DUPLICATE of `common-service/proto/bonus/v1/bonus_service.proto`
  - Status: Different structure, needs comparison

**Commission Service:**
- ❌ `/workspace/services/commission-service/src/main/proto/commission_service.proto`
  - DUPLICATE of `common-service/proto/commission/v1/commission_service.proto`
  - Status: Similar but common-service version is more complete

**Payment Service:**
- ❌ `/workspace/services/payment-service/src/main/proto/payment_service.proto`
  - DUPLICATE of `common-service/proto/payment/v1/payment_service.proto`
  - Status: Simpler version

**Payment Service - Common Types:**
- ❌ `/workspace/services/payment-service/src/main/proto/game_engine/common/v1/enums.proto`
- ❌ `/workspace/services/payment-service/src/main/proto/game_engine/common/v1/money.proto`
- ❌ `/workspace/services/payment-service/src/main/proto/game_engine/common/v1/pagination.proto`
  - All DUPLICATES of corresponding files in `common-service/proto/common/v1/`

**Payment Service - Referenced Services:**
- ❌ `/workspace/services/payment-service/src/main/proto/game_engine/risk/v1/risk_service.proto`
  - DUPLICATE of `common-service/proto/risk/v1/risk_service.proto`
- ❌ `/workspace/services/payment-service/src/main/proto/game_engine/wallet/v1/wallet_service.proto`
  - DUPLICATE of `common-service/proto/wallet/v1/wallet_service.proto`
- ❌ `/workspace/services/payment-service/src/main/proto/game_engine/bonus/v1/bonus_service.proto`
  - DUPLICATE of `common-service/proto/bonus/v1/bonus_service.proto`

**Payment Service - Google Protobuf:**
- ⚠️ `/workspace/services/payment-service/src/main/proto/google/protobuf/empty.proto`
- ⚠️ `/workspace/services/payment-service/src/main/proto/google/protobuf/timestamp.proto`
  - These are standard protobuf types, should use official imports

---

## 2. Detailed Comparison Analysis

### 2.1 Auth Service Comparison

| Aspect | `/workspace/proto/auth/auth.proto` | `common-service/proto/auth/v1/auth_service.proto` |
|--------|-----------------------------------|--------------------------------------------------|
| Package | `auth` | `game_engine.auth.v1` |
| Go Package | `github.com/royalplatform/proto/auth` | `github.com/game_engine/common-service/proto/gen/go/auth/v1` |
| Validation | Uses buf-validate | No validation rules |
| Features | More detailed admin functions | Better structured with comments |
| Timestamps | Uses google.protobuf.Timestamp | Uses google.protobuf.Timestamp |
| **Status** | ❌ OUTDATED | ✅ MASTER |

### 2.2 Wallet Service Comparison

| Aspect | `/workspace/proto/wallet/wallet.proto` | `common-service/proto/wallet/v1/wallet_service.proto` |
|--------|---------------------------------------|------------------------------------------------------|
| Package | `wallet` | `game_engine.wallet.v1` |
| Go Package | `github.com/royalplatform/proto/wallet` | `github.com/game_engine/common-service/proto/gen/go/wallet/v1` |
| Money Type | Custom int64 amount | Uses common.v1.Money |
| Features | More detailed withdrawal/deposit flow | Includes betting operations |
| **Status** | ❌ OUTDATED | ✅ MASTER |

### 2.3 Affiliate Service Comparison

| Aspect | Affiliate-service copy | Common-service master |
|--------|----------------------|----------------------|
| Lines | 208 lines | 341 lines |
| Additional RPCs | Basic set | + Performance reports, link management |
| Timestamps | int64 | google.protobuf.Timestamp |
| **Status** | ❌ INCOMPLETE | ✅ COMPLETE |

### 2.4 Payment Service Comparison

| Aspect | Payment-service copy | Common-service master |
|--------|---------------------|----------------------|
| Lines | ~150 lines | ~180 lines |
| Structure | Simpler | Better organized with comments |
| Timestamps | int64 | google.protobuf.Timestamp |
| **Status** | ❌ SIMPLER | ✅ PREFERRED |

### 2.5 Commission Service Comparison

| Aspect | Commission-service copy | Common-service master |
|--------|------------------------|----------------------|
| Lines | ~200+ lines | ~400+ lines |
| Services | 3 services | 3 services (enhanced) |
| Claims Support | Basic | Comprehensive (rebet, insurance, settlement) |
| **Status** | ❌ OUTDATED | ✅ ENHANCED |

### 2.6 Bonus Service Comparison

| Aspect | Bonus-service copy | Common-service master |
|--------|-------------------|----------------------|
| Structure | Flat messages | Organized with sections |
| Rebet Claims | No | Yes |
| Insurance Claims | No | Yes |
| **Status** | ❌ BASIC | ✅ FEATURE-RICH |

### 2.7 Risk Service Comparison

| Aspect | Payment-service copy | Common-service master |
|--------|---------------------|----------------------|
| Functions | Limited | Comprehensive risk profiling |
| Limits | Basic | Detailed limit calculations |
| **Status** | ❌ LIMITED | ✅ COMPREHENSIVE |

---

## 3. Issues Identified

### 3.1 Critical Issues

1. **Multiple Versions**: Same service defined in 2-3 different locations
2. **Inconsistent Packages**: Different package names (`auth` vs `game_engine.auth.v1`)
3. **Inconsistent Go Packages**: Different go_package paths
4. **Feature Divergence**: Master versions have more features than copies
5. **Timestamp Inconsistency**: Some use `int64`, others use `google.protobuf.Timestamp`
6. **Import Path Issues**: Relative imports pointing to wrong locations

### 3.2 Build & Maintenance Issues

1. **Code Duplication**: 15 duplicate proto files
2. **Sync Problems**: Changes in one location not reflected in others
3. **Import Confusion**: Services importing from different sources
4. **Generation Conflicts**: Generated code may conflict

---

## 4. Recommended Actions

### Phase 1: Immediate (Priority: CRITICAL)

#### 4.1.1 Delete Legacy Root Protos
```bash
# Remove legacy root-level protos
rm -rf /workspace/proto/
```

#### 4.1.2 Update Import Paths in Payment Service
The payment service has copies of common types. Update imports to reference common-service:

**Files to update:**
- `/workspace/services/payment-service/src/main/proto/payment_service.proto`

**Change imports from:**
```proto
import "game_engine/common/v1/enums.proto";
import "game_engine/common/v1/money.proto";
import "game_engine/common/v1/pagination.proto";
import "game_engine/risk/v1/risk_service.proto";
import "game_engine/wallet/v1/wallet_service.proto";
import "game_engine/bonus/v1/bonus_service.proto";
```

**To:**
```proto
// These will be resolved via buf.yaml dependencies or protoc include paths
import "common/v1/enums.proto";
import "common/v1/money.proto";
import "common/v1/pagination.proto";
// Risk, wallet, bonus should reference the common-service definitions
```

#### 4.1.3 Remove Duplicate Service Definitions
```bash
# Remove duplicate proto files from individual services
rm /workspace/services/affiliate-service/src/main/proto/affiliate_service.proto
rm /workspace/services/bonus-service/src/main/proto/bonus_service.proto
rm /workspace/services/commission-service/src/main/proto/commission_service.proto
rm /workspace/services/payment-service/src/main/proto/payment_service.proto

# Remove duplicate common types
rm -rf /workspace/services/payment-service/src/main/proto/game_engine/

# Remove standard protobuf copies (use official imports)
rm -rf /workspace/services/payment-service/src/main/proto/google/
```

### Phase 2: Configuration Updates (Priority: HIGH)

#### 4.2.1 Update buf.yaml in Each Service
Ensure all services reference common-service protos:

```yaml
version: v1
deps:
  - buf.build/googleapis/googleapis
  - buf.build/bufbuild/protovalidate
build:
  roots:
    - src/main/proto
    - ../common-service/proto  # Add common-service as dependency
```

#### 4.2.2 Update Build Scripts
Update all Makefiles and build scripts to:
1. Generate code from common-service first
2. Use common-service protos as dependencies
3. Avoid regenerating common types in each service

### Phase 3: Service Code Updates (Priority: HIGH)

#### 4.3.1 Update Java Services
Update import statements in Java code:
```java
// OLD
import com.game_engine.affiliate.v1.AffiliateService;

// NEW (if package changed)
import com.game_engine.affiliate.v1.AffiliateService;
```

#### 4.3.2 Update Go Services
Update import paths in Go code:
```go
// OLD
import "gen/go/auth/v1"

// NEW
import "github.com/game_engine/common-service/proto/gen/go/auth/v1"
```

#### 4.3.3 Update Python Services
Update import paths in Python code similarly.

### Phase 4: Verification (Priority: MEDIUM)

#### 4.4.1 Compile All Protos
```bash
cd /workspace/services/common-service/proto
buf generate

# Verify no errors in dependent services
cd /workspace/services/payment-service
buf generate
```

#### 4.4.2 Run Tests
Ensure all services compile and tests pass after proto consolidation.

#### 4.4.3 Integration Testing
Test inter-service communication to ensure proto compatibility.

---

## 5. File Inventory

### 5.1 Files to KEEP (Master Copies)
All files in `/workspace/services/common-service/proto/`:
```
✅ common/v1/enums.proto
✅ common/v1/error.proto
✅ common/v1/money.proto
✅ common/v1/pagination.proto
✅ common/v1/timestamp.proto
✅ affiliate/v1/affiliate_service.proto
✅ agent/v1/agent_service.proto
✅ auth/v1/auth_service.proto
✅ bonus/v1/bonus_service.proto
✅ commission/v1/commission_service.proto
✅ game/v1/game_registry.proto
✅ jackpot/v1/jackpot_service.proto
✅ leaderboard/v1/leaderboard_service.proto
✅ merchant/v1/merchant_service.proto
✅ payment/v1/payment_service.proto
✅ risk/v1/risk_service.proto
✅ tournament/v1/tournament_service.proto
✅ user/v1/user_service.proto
✅ wallet/v1/wallet_service.proto
✅ winners/v1/winners_service.proto
```

### 5.2 Files to DELETE
```
❌ /workspace/proto/auth/auth.proto
❌ /workspace/proto/wallet/wallet.proto
❌ /workspace/proto/common/common.proto
❌ /workspace/services/affiliate-service/src/main/proto/affiliate_service.proto
❌ /workspace/services/bonus-service/src/main/proto/bonus_service.proto
❌ /workspace/services/commission-service/src/main/proto/commission_service.proto
❌ /workspace/services/payment-service/src/main/proto/payment_service.proto
❌ /workspace/services/payment-service/src/main/proto/game_engine/ (entire directory)
❌ /workspace/services/payment-service/src/main/proto/google/ (entire directory)
```

### 5.3 Files to UPDATE
```
⚠️ All service build configurations (buf.yaml, Makefile, pom.xml, etc.)
⚠️ All service source code imports
⚠️ CI/CD pipelines
⚠️ Documentation
```

---

## 6. Migration Checklist

- [ ] Backup current state
- [ ] Delete legacy `/workspace/proto/` directory
- [ ] Delete duplicate protos in service directories
- [ ] Update buf.yaml in all services
- [ ] Update build scripts (Makefiles, Gradle, etc.)
- [ ] Update import statements in Go code
- [ ] Update import statements in Java code
- [ ] Update import statements in Python code
- [ ] Regenerate all proto code
- [ ] Compile all services
- [ ] Run unit tests
- [ ] Run integration tests
- [ ] Update documentation
- [ ] Commit changes

---

## 7. Conclusion

**Current State:** 35 proto files scattered across 6 locations
**Target State:** 20 proto files in 1 location (`/workspace/services/common-service/proto/`)

**Benefits of Consolidation:**
1. Single source of truth
2. Easier maintenance
3. Consistent versioning
4. Reduced build complexity
5. Clearer dependencies
6. Better code reuse

**Estimated Effort:** 4-8 hours for complete migration
**Risk Level:** Medium (requires careful testing)
**Recommended Timeline:** Complete in next sprint

---

*Report generated: $(date)*
*Total proto files analyzed: 35*
*Duplicates identified: 15*
*Master files: 20*
