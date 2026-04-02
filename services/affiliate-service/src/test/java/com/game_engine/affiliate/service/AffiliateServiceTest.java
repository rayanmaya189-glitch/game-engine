package com.game_engine.affiliate.service;

import com.game_engine.affiliate.entity.Affiliate;
import com.game_engine.affiliate.entity.Referral;
import com.game_engine.affiliate.repository.AffiliateRepository;
import com.game_engine.affiliate.repository.ReferralRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class AffiliateServiceTest {

    @Mock
    private AffiliateRepository affiliateRepository;

    @Mock
    private ReferralRepository referralRepository;

    @InjectMocks
    private AffiliateService affiliateService;

    private Affiliate testAffiliate;
    private Referral testReferral;

    @BeforeEach
    void setUp() {
        testAffiliate = new Affiliate();
        testAffiliate.setId(1L);
        testAffiliate.setName("Test Affiliate");
        testAffiliate.setEmail("test@example.com");
        testAffiliate.setAffiliateCode("AFF-ABC12345");
        testAffiliate.setStatus("ACTIVE");
        testAffiliate.setTier("BRONZE");
        testAffiliate.setCommissionRate(new BigDecimal("0.2000"));
        testAffiliate.setTotalReferrals(0);
        testAffiliate.setTotalRevenue(BigDecimal.ZERO);

        testReferral = new Referral();
        testReferral.setId(1L);
        testReferral.setAffiliateCode("AFF-ABC12345");
        testReferral.setPlayerId(100L);
        testReferral.setSource("LINK");
        testReferral.setStatus("ACTIVE");
        testReferral.setTotalDeposits(new BigDecimal("500.00"));
    }

    @Test
    void registerAffiliate_shouldCreateNewAffiliate() {
        when(affiliateRepository.findByEmail("new@example.com")).thenReturn(Optional.empty());
        when(affiliateRepository.save(any(Affiliate.class))).thenAnswer(inv -> inv.getArgument(0));

        Affiliate result = affiliateService.registerAffiliate("New User", "new@example.com");

        assertNotNull(result);
        assertEquals("New User", result.getName());
        assertEquals("new@example.com", result.getEmail());
        assertTrue(result.getAffiliateCode().startsWith("AFF-"));
        verify(affiliateRepository).save(any(Affiliate.class));
    }

    @Test
    void registerAffiliate_shouldThrowWhenEmailExists() {
        when(affiliateRepository.findByEmail("test@example.com")).thenReturn(Optional.of(testAffiliate));

        assertThrows(RuntimeException.class,
                () -> affiliateService.registerAffiliate("Test", "test@example.com"));
        verify(affiliateRepository, never()).save(any());
    }

    @Test
    void trackReferral_shouldCreateReferralAndIncrementCount() {
        when(affiliateRepository.findByAffiliateCode("AFF-ABC12345")).thenReturn(Optional.of(testAffiliate));
        when(affiliateRepository.save(any(Affiliate.class))).thenReturn(testAffiliate);
        when(referralRepository.save(any(Referral.class))).thenAnswer(inv -> inv.getArgument(0));

        Referral result = affiliateService.trackReferral("AFF-ABC12345", 100L, "LINK");

        assertNotNull(result);
        assertEquals("AFF-ABC12345", result.getAffiliateCode());
        assertEquals(100L, result.getPlayerId());
        assertEquals("ACTIVE", result.getStatus());
        assertEquals(1, testAffiliate.getTotalReferrals());
        verify(referralRepository).save(any(Referral.class));
    }

    @Test
    void trackReferral_shouldThrowWhenAffiliateNotFound() {
        when(affiliateRepository.findByAffiliateCode("INVALID")).thenReturn(Optional.empty());

        assertThrows(RuntimeException.class,
                () -> affiliateService.trackReferral("INVALID", 100L, "LINK"));
    }

    @Test
    void calculateCommissions_shouldReturnCorrectCommission() {
        Referral r1 = new Referral();
        r1.setTotalDeposits(new BigDecimal("1000.00"));
        Referral r2 = new Referral();
        r2.setTotalDeposits(new BigDecimal("500.00"));

        when(affiliateRepository.findById(1L)).thenReturn(Optional.of(testAffiliate));
        when(referralRepository.findByAffiliateCode("AFF-ABC12345")).thenReturn(Arrays.asList(r1, r2));
        when(affiliateRepository.save(any(Affiliate.class))).thenReturn(testAffiliate);

        BigDecimal commission = affiliateService.calculateCommissions(1L);

        assertEquals(0, new BigDecimal("300.00").compareTo(commission));
        assertEquals(0, new BigDecimal("1500.00").compareTo(testAffiliate.getTotalRevenue()));
    }

    @Test
    void calculateCommissions_shouldThrowWhenAffiliateNotFound() {
        when(affiliateRepository.findById(999L)).thenReturn(Optional.empty());

        assertThrows(RuntimeException.class, () -> affiliateService.calculateCommissions(999L));
    }

    @Test
    void getAffiliateStats_shouldReturnCorrectStats() {
        Referral activeRef = new Referral();
        activeRef.setStatus("ACTIVE");
        activeRef.setTotalDeposits(new BigDecimal("1000.00"));

        Referral inactiveRef = new Referral();
        inactiveRef.setStatus("INACTIVE");
        inactiveRef.setTotalDeposits(new BigDecimal("200.00"));

        Referral nullDepositsRef = new Referral();
        nullDepositsRef.setStatus("ACTIVE");
        nullDepositsRef.setTotalDeposits(null);

        when(affiliateRepository.findById(1L)).thenReturn(Optional.of(testAffiliate));
        when(referralRepository.findByAffiliateCode("AFF-ABC12345"))
                .thenReturn(Arrays.asList(activeRef, inactiveRef, nullDepositsRef));

        AffiliateService.AffiliateStats stats = affiliateService.getAffiliateStats(1L);

        assertEquals(3, stats.getTotalReferrals());
        assertEquals(2, stats.getActiveReferrals());
        assertEquals(0, new BigDecimal("1200.00").compareTo(stats.getTotalDeposits()));
        assertEquals(0, new BigDecimal("240.00").compareTo(stats.getTotalCommission()));
    }

    @Test
    void updateAffiliateTier_shouldUpdateTierAndRate() {
        when(affiliateRepository.findById(1L)).thenReturn(Optional.of(testAffiliate));
        when(affiliateRepository.save(any(Affiliate.class))).thenAnswer(inv -> inv.getArgument(0));

        Affiliate result = affiliateService.updateAffiliateTier(1L, "GOLD");

        assertEquals("GOLD", result.getTier());
        assertEquals(0, new BigDecimal("0.3000").compareTo(result.getCommissionRate()));
    }

    @Test
    void updateAffiliateStatus_shouldUpdateStatus() {
        when(affiliateRepository.findById(1L)).thenReturn(Optional.of(testAffiliate));
        when(affiliateRepository.save(any(Affiliate.class))).thenAnswer(inv -> inv.getArgument(0));

        Affiliate result = affiliateService.updateAffiliateStatus(1L, "SUSPENDED");

        assertEquals("SUSPENDED", result.getStatus());
    }

    @Test
    void getAffiliateByCode_shouldReturnAffiliate() {
        when(affiliateRepository.findByAffiliateCode("AFF-ABC12345")).thenReturn(Optional.of(testAffiliate));

        Optional<Affiliate> result = affiliateService.getAffiliateByCode("AFF-ABC12345");

        assertTrue(result.isPresent());
        assertEquals("Test Affiliate", result.get().getName());
    }

    @Test
    void getReferralsByAffiliateCode_shouldReturnReferrals() {
        when(referralRepository.findByAffiliateCode("AFF-ABC12345"))
                .thenReturn(Arrays.asList(testReferral));

        List<Referral> result = affiliateService.getReferralsByAffiliateCode("AFF-ABC12345");

        assertEquals(1, result.size());
        assertEquals(100L, result.get(0).getPlayerId());
    }
}
