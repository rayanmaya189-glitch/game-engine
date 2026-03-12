package com.game_engine.bonus.controller;

import com.game_engine.bonus.model.Bonus;
import com.game_engine.bonus.service.BonusService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@RestController
@RequestMapping("/api/v1/bonuses")
@RequiredArgsConstructor
public class BonusController {
    private final BonusService bonusService;

    @GetMapping
    public ResponseEntity<List<Bonus>> getActiveBonuses() {
        return ResponseEntity.ok(bonusService.getActiveBonuses());
    }

    @GetMapping("/{id}")
    public ResponseEntity<Bonus> getBonus(@PathVariable UUID id) {
        return ResponseEntity.ok(bonusService.getBonusById(id));
    }

    @PostMapping
    public ResponseEntity<Bonus> createBonus(@RequestBody Bonus bonus) {
        return ResponseEntity.ok(bonusService.createBonus(bonus));
    }

    @PostMapping("/{id}/claim")
    public ResponseEntity<Map<String, Object>> claimBonus(@PathVariable UUID id, @RequestParam UUID userId) {
        Map<String, Object> result = bonusService.claimBonus(id, userId);
        return ResponseEntity.ok(result);
    }

    @GetMapping("/eligibility/{userId}")
    public ResponseEntity<Map<String, Object>> checkEligibility(@PathVariable UUID userId) {
        Map<String, Object> eligibility = bonusService.checkEligibility(userId);
        return ResponseEntity.ok(eligibility);
    }

    @GetMapping("/history/{userId}")
    public ResponseEntity<List<Map<String, Object>>> getBonusHistory(@PathVariable UUID userId) {
        List<Map<String, Object>> history = bonusService.getBonusHistory(userId);
        return ResponseEntity.ok(history);
    }

    @GetMapping("/active/{userId}")
    public ResponseEntity<List<Map<String, Object>>> getActiveBonusClaims(@PathVariable UUID userId) {
        List<Map<String, Object>> activeClaims = bonusService.getActiveBonusClaims(userId);
        return ResponseEntity.ok(activeClaims);
    }

    @PostMapping("/wagering/contribute")
    public ResponseEntity<Map<String, Object>> processWageringContribution(
            @RequestParam UUID userId,
            @RequestParam UUID bonusId,
            @RequestParam BigDecimal betAmount,
            @RequestParam String gameType) {
        Map<String, Object> result = bonusService.processWageringContribution(userId, bonusId, betAmount, gameType);
        return ResponseEntity.ok(result);
    }

    @PostMapping("/complete")
    public ResponseEntity<Map<String, Object>> completeBonus(
            @RequestParam UUID userId,
            @RequestParam UUID bonusId,
            @RequestParam BigDecimal winnings) {
        Map<String, Object> result = bonusService.completeBonus(userId, bonusId, winnings);
        return ResponseEntity.ok(result);
    }

    @PostMapping("/cancel")
    public ResponseEntity<Map<String, Object>> cancelBonus(
            @RequestParam UUID userId,
            @RequestParam UUID bonusId,
            @RequestParam String reason) {
        Map<String, Object> result = bonusService.cancelBonus(userId, bonusId, reason);
        return ResponseEntity.ok(result);
    }

    @GetMapping("/stats/{userId}")
    public ResponseEntity<Map<String, Object>> getBonusStats(@PathVariable UUID userId) {
        Map<String, Object> stats = bonusService.getBonusStats(userId);
        return ResponseEntity.ok(stats);
    }
}
