package com.game-engine.bonus.repository;

import com.game-engine.bonus.model.Bonus;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.List;
import java.util.UUID;

public interface BonusRepository extends JpaRepository<Bonus, UUID> {
    List<Bonus> findByStatus(Bonus.BonusStatus status);
    List<Bonus> findByType(Bonus.BonusType type);
}
