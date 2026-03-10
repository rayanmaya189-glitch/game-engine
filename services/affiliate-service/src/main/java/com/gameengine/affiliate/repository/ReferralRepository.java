package com.gameengine.affiliate.repository;

import com.gameengine.affiliate.model.Referral;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Repository
public interface ReferralRepository extends JpaRepository<Referral, Long> {
    
    Optional<Referral> findByReferralCode(String referralCode);
    
    List<Referral> findByAffiliateId(Long affiliateId);
    
    @Query("SELECT r FROM Referral r WHERE r.affiliate.affiliateCode = :affiliateCode ORDER BY r.clickedAt DESC")
    List<Referral> findByAffiliateCode(@Param("affiliateCode") String affiliateCode);
    
    @Query("SELECT r FROM Referral r WHERE r.status = :status AND r.affiliate.id = :affiliateId")
    List<Referral> findByStatusAndAffiliateId(@Param("status") String status, @Param("affiliateId") Long affiliateId);
    
    @Query("SELECT COUNT(r) FROM Referral r WHERE r.affiliate.id = :affiliateId AND r.status = 'REGISTERED'")
    Long countRegistrations(@Param("affiliateId") Long affiliateId);
    
    @Query("SELECT COUNT(r) FROM Referral r WHERE r.affiliate.id = :affiliateId AND r.status = 'DEPOSITED'")
    Long countDepositors(@Param("affiliateId") Long affiliateId);
    
    @Query("SELECT r FROM Referral r WHERE r.clickedAt >= :startDate AND r.clickedAt <= :endDate")
    List<Referral> findByClickDateRange(@Param("startDate") LocalDateTime startDate, @Param("endDate") LocalDateTime endDate);
    
    List<Referral> findByCampaignId(String campaignId);
}
