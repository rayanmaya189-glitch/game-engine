package com.gameengine.bonus.controller;

import com.gameengine.bonus.model.Bonus;
import com.gameengine.bonus.service.BonusService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.util.List;
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
}
