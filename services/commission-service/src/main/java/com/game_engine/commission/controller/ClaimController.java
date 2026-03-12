package com.game_engine.commission.controller;

import com.game_engine.commission.model.*;
import com.game_engine.commission.service.ClaimService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/v1/claims")
public class ClaimController {
    
    @Autowired
    private ClaimService claimService;
    
    // Commission Claims
    @PostMapping("/commission")
    public ResponseEntity<CommissionClaim> submitCommissionClaim(@RequestBody Map<String, Object> request) {
        Long userId = Long.valueOf(request.get("userId").toString());
        Long affiliateId = Long.valueOf(request.get("affiliateId").toString());
        Long commissionId = Long.valueOf(request.get("commissionId").toString());
        BigDecimal amount = new BigDecimal(request.get("amount").toString());
        String claimReason = request.get("claimReason").toString();
        
        CommissionClaim claim = claimService.submitCommissionClaim(userId, affiliateId, commissionId, amount, claimReason);
        return ResponseEntity.status(HttpStatus.CREATED).body(claim);
    }
    
    @GetMapping("/commission/user/{userId}")
    public ResponseEntity<List<CommissionClaim>> getUserCommissionClaims(@PathVariable Long userId) {
        List<CommissionClaim> claims = claimService.getUserCommissionClaims(userId);
        return ResponseEntity.ok(claims);
    }
    
    @GetMapping("/commission/status/{status}")
    public ResponseEntity<List<CommissionClaim>> getCommissionClaimsByStatus(@PathVariable String status) {
        List<CommissionClaim> claims = claimService.getCommissionClaimsByStatus(status);
        return ResponseEntity.ok(claims);
    }
    
    @GetMapping("/commission/{id}")
    public ResponseEntity<CommissionClaim> getCommissionClaimById(@PathVariable Long id) {
        return claimService.getCommissionClaimById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    @PostMapping("/commission/{id}/approve")
    public ResponseEntity<CommissionClaim> approveCommissionClaim(@PathVariable Long id, @RequestBody Map<String, String> request) {
        String adminNote = request.get("adminNote");
        CommissionClaim claim = claimService.approveCommissionClaim(id, adminNote);
        return ResponseEntity.ok(claim);
    }
    
    @PostMapping("/commission/{id}/reject")
    public ResponseEntity<CommissionClaim> rejectCommissionClaim(@PathVariable Long id, @RequestBody Map<String, String> request) {
        String adminNote = request.get("adminNote");
        CommissionClaim claim = claimService.rejectCommissionClaim(id, adminNote);
        return ResponseEntity.ok(claim);
    }
    
    @PostMapping("/commission/{id}/pay")
    public ResponseEntity<CommissionClaim> payCommissionClaim(@PathVariable Long id) {
        CommissionClaim claim = claimService.payCommissionClaim(id);
        return ResponseEntity.ok(claim);
    }
    
    // Rebet Claims
    @PostMapping("/rebet")
    public ResponseEntity<RebetClaim> createRebetClaim(@RequestBody Map<String, Object> request) {
        Long userId = Long.valueOf(request.get("userId").toString());
        Long bonusId = Long.valueOf(request.get("bonusId").toString());
        String bonusCode = request.get("bonusCode").toString();
        BigDecimal bonusAmount = new BigDecimal(request.get("bonusAmount").toString());
        BigDecimal rebetRequirement = new BigDecimal(request.get("rebetRequirement").toString());
        Long gameId = Long.valueOf(request.get("gameId").toString());
        Long betId = Long.valueOf(request.get("betId").toString());
        
        RebetClaim claim = claimService.createRebetClaim(userId, bonusId, bonusCode, bonusAmount, rebetRequirement, gameId, betId);
        return ResponseEntity.status(HttpStatus.CREATED).body(claim);
    }
    
    @PostMapping("/rebet/{id}/update-progress")
    public ResponseEntity<RebetClaim> updateRebetProgress(@PathVariable Long id, @RequestBody Map<String, Object> request) {
        BigDecimal additionalBetAmount = new BigDecimal(request.get("amount").toString());
        RebetClaim claim = claimService.updateRebetProgress(id, additionalBetAmount);
        return ResponseEntity.ok(claim);
    }
    
    @PostMapping("/rebet/{id}/claim")
    public ResponseEntity<RebetClaim> claimRebet(@PathVariable Long id) {
        RebetClaim claim = claimService.claimRebet(id);
        return ResponseEntity.ok(claim);
    }
    
    @GetMapping("/rebet/user/{userId}")
    public ResponseEntity<List<RebetClaim>> getUserRebetClaims(@PathVariable Long userId) {
        List<RebetClaim> claims = claimService.getUserRebetClaims(userId);
        return ResponseEntity.ok(claims);
    }
    
    @GetMapping("/rebet/user/{userId}/claimable")
    public ResponseEntity<List<RebetClaim>> getClaimableRebets(@PathVariable Long userId) {
        List<RebetClaim> claims = claimService.getClaimableRebets(userId);
        return ResponseEntity.ok(claims);
    }
    
    // Insurance Claims
    @PostMapping("/insurance")
    public ResponseEntity<InsuranceClaim> submitInsuranceClaim(@RequestBody Map<String, Object> request) {
        Long userId = Long.valueOf(request.get("userId").toString());
        Long gameId = Long.valueOf(request.get("gameId").toString());
        Long betId = Long.valueOf(request.get("betId").toString());
        String insurancePolicyId = request.get("insurancePolicyId").toString();
        String claimType = request.get("claimType").toString();
        BigDecimal insuredAmount = new BigDecimal(request.get("insuredAmount").toString());
        BigDecimal lossAmount = new BigDecimal(request.get("lossAmount").toString());
        String claimReason = request.get("claimReason").toString();
        String evidenceDetails = request.get("evidenceDetails") != null ? request.get("evidenceDetails").toString() : null;
        
        InsuranceClaim claim = claimService.submitInsuranceClaim(userId, gameId, betId, insurancePolicyId, 
                claimType, insuredAmount, lossAmount, claimReason, evidenceDetails);
        return ResponseEntity.status(HttpStatus.CREATED).body(claim);
    }
    
    @PostMapping("/insurance/{id}/approve")
    public ResponseEntity<InsuranceClaim> approveInsuranceClaim(@PathVariable Long id, @RequestBody Map<String, Object> request) {
        Long reviewedBy = Long.valueOf(request.get("reviewedBy").toString());
        String adminNote = request.get("adminNote").toString();
        InsuranceClaim claim = claimService.approveInsuranceClaim(id, reviewedBy, adminNote);
        return ResponseEntity.ok(claim);
    }
    
    @PostMapping("/insurance/{id}/reject")
    public ResponseEntity<InsuranceClaim> rejectInsuranceClaim(@PathVariable Long id, @RequestBody Map<String, Object> request) {
        Long reviewedBy = Long.valueOf(request.get("reviewedBy").toString());
        String adminNote = request.get("adminNote").toString();
        InsuranceClaim claim = claimService.rejectInsuranceClaim(id, reviewedBy, adminNote);
        return ResponseEntity.ok(claim);
    }
    
    @PostMapping("/insurance/{id}/pay")
    public ResponseEntity<InsuranceClaim> payInsuranceClaim(@PathVariable Long id) {
        InsuranceClaim claim = claimService.payInsuranceClaim(id);
        return ResponseEntity.ok(claim);
    }
    
    @GetMapping("/insurance/user/{userId}")
    public ResponseEntity<List<InsuranceClaim>> getUserInsuranceClaims(@PathVariable Long userId) {
        List<InsuranceClaim> claims = claimService.getUserInsuranceClaims(userId);
        return ResponseEntity.ok(claims);
    }
    
    @GetMapping("/insurance/status/{status}")
    public ResponseEntity<List<InsuranceClaim>> getInsuranceClaimsByStatus(@PathVariable String status) {
        List<InsuranceClaim> claims = claimService.getInsuranceClaimsByStatus(status);
        return ResponseEntity.ok(claims);
    }
    
    // Settlements
    @GetMapping("/settlement/user/{userId}")
    public ResponseEntity<List<Settlement>> getUserSettlements(@PathVariable Long userId) {
        List<Settlement> settlements = claimService.getUserSettlements(userId);
        return ResponseEntity.ok(settlements);
    }
    
    @GetMapping("/settlement/status/{status}")
    public ResponseEntity<List<Settlement>> getSettlementsByStatus(@PathVariable String status) {
        List<Settlement> settlements = claimService.getSettlementsByStatus(status);
        return ResponseEntity.ok(settlements);
    }
    
    @GetMapping("/settlement/type/{type}")
    public ResponseEntity<List<Settlement>> getSettlementsByType(@PathVariable String type) {
        List<Settlement> settlements = claimService.getSettlementsByType(type);
        return ResponseEntity.ok(settlements);
    }
    
    @GetMapping("/settlement/{id}")
    public ResponseEntity<Settlement> getSettlementById(@PathVariable Long id) {
        return claimService.getSettlementById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    @GetMapping("/user/{userId}/total-pending")
    public ResponseEntity<Map<String, BigDecimal>> getUserTotalPending(@PathVariable Long userId) {
        BigDecimal total = claimService.getUserTotalPendingClaims(userId);
        return ResponseEntity.ok(Map.of("totalPending", total));
    }
    
    @GetMapping("/user/{userId}/total-settled")
    public ResponseEntity<Map<String, BigDecimal>> getUserTotalSettled(@PathVariable Long userId) {
        BigDecimal total = claimService.getUserTotalSettled(userId);
        return ResponseEntity.ok(Map.of("totalSettled", total));
    }
}
