package com.game-engine.bonus.controller;

import com.game-engine.bonus.model.Bonus;
import com.game-engine.bonus.service.BonusService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
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
}
