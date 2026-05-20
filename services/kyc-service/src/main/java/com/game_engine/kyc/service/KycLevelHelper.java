package com.game_engine.kyc.service;

import com.game_engine.kyc.model.KycVerification;
import com.game_engine.kyc.model.KycVerification.KycStatus;
import com.game_engine.kyc.model.KycVerification.VerificationLevel;

public class KycLevelHelper {

    public static VerificationLevel getRequiredLevelForAction(
            double totalDeposits, double withdrawalAmount,
            double highDepositTrigger, double depositCumulativeTrigger) {
        if (withdrawalAmount > 0) {
            return VerificationLevel.LEVEL_2;
        }
        if (totalDeposits >= highDepositTrigger) {
            return VerificationLevel.LEVEL_3;
        }
        if (totalDeposits >= depositCumulativeTrigger) {
            return VerificationLevel.LEVEL_2;
        }
        return VerificationLevel.LEVEL_0;
    }

    public static KycVerification createNewVerification(java.util.UUID userId) {
        return KycVerification.builder()
                .userId(userId)
                .level(VerificationLevel.LEVEL_0)
                .status(KycStatus.PENDING)
                .build();
    }

    public static void updateLevelIfComplete(KycVerification verification) {
        switch (verification.getLevel()) {
            case LEVEL_0:
                if (verification.getEmailVerified() != null && verification.getPhoneVerified() != null) {
                    if (verification.getEmailVerified() && verification.getPhoneVerified()) {
                        verification.setLevel(VerificationLevel.LEVEL_1);
                        verification.setStatus(KycStatus.VERIFIED);
                    }
                }
                break;
            case LEVEL_1:
                if (Boolean.TRUE.equals(verification.getDocumentVerified())) {
                    verification.setLevel(VerificationLevel.LEVEL_2);
                }
                break;
            case LEVEL_2:
                if (Boolean.TRUE.equals(verification.getAddressVerified()) &&
                        Boolean.TRUE.equals(verification.getSourceOfFundsVerified())) {
                    verification.setLevel(VerificationLevel.LEVEL_3);
                }
                break;
        }
    }
}
