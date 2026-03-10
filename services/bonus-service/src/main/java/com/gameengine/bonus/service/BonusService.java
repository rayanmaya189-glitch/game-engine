package com.gameengine.bonus.service;

import com.gameengine.bonus.model.Bonus;
import com.gameengine.bonus.repository.BonusRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import java.math.BigDecimal;
import java.util.List;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class BonusService {
    private final BonusRepository bonusRepository;

    public Bonus createBonus(Bonus bonus) {
        return bonusRepository.save(bonus);
    }

    public List<Bonus> getActiveBonuses() {
        return bonusRepository.findByStatus(Bonus.BonusStatus.ACTIVE);
    }

    public Bonus getBonusById(UUID id) {
        return bonusRepository.findById(id).orElse(null);
    }

    public BigDecimal calculateWageringContribution(BigDecimal betAmount, String gameType) {
        // Slot games contribute 100%, table games 10%
        return "SLOT".equalsIgnoreCase(gameType) ? betAmount : betAmount.multiply(BigDecimal.valueOf(0.1));
    }

    public boolean isWageringComplete(UUID bonusId, BigDecimal totalWagered, BigDecimal wageringRequirement) {
        return totalWagered.compareTo(wageringRequirement) >= 0;
    }
}
