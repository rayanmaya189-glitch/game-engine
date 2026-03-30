package com.game_engine.payment.service;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Deposit.*;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
@RequiredArgsConstructor
@Slf4j
public class GatewayResolver {

    private final List<PaymentGatewayAdapter> gatewayAdapters;

    public PaymentGatewayAdapter getGatewayAdapter(PaymentGateway gateway) {
        return gatewayAdapters.stream()
                .filter(a -> a.getGatewayType() == gateway)
                .findFirst()
                .orElseThrow(() -> new IllegalArgumentException("Gateway not supported: " + gateway));
    }

    public List<PaymentGatewayAdapter> getAvailableAdapters() {
        return gatewayAdapters;
    }

    public DepositStatus mapResponseStatusToDepositStatus(String responseStatus) {
        if (responseStatus == null) return DepositStatus.FAILED;

        return switch (responseStatus) {
            case "COMPLETED" -> DepositStatus.COMPLETED;
            case "PROCESSING" -> DepositStatus.PROCESSING;
            case "PENDING_VERIFICATION" -> DepositStatus.PENDING_VERIFICATION;
            case "FAILED" -> DepositStatus.FAILED;
            case "CANCELLED" -> DepositStatus.CANCELLED;
            default -> DepositStatus.PROCESSING;
        };
    }
}
