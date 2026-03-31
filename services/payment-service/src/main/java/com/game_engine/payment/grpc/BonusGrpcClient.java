package com.game_engine.payment.grpc;

import com.game_engine.bonus.v1.BonusServiceGrpc;
import com.game_engine.bonus.v1.CheckWageringRequirementsRequest;
import com.game_engine.bonus.v1.CheckWageringRequirementsResponse;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.util.UUID;
import java.util.concurrent.TimeUnit;

@Component
@Slf4j
public class BonusGrpcClient {

    @Value("${bonus.service.host:bonus-service}")
    private String bonusServiceHost;

    @Value("${bonus.service.grpc.port:9013}")
    private int bonusServicePort;

    @Value("${bonus.service.timeout-seconds:10}")
    private long timeoutSeconds;

    private ManagedChannel channel;
    private BonusServiceGrpc.BonusServiceBlockingStub blockingStub;

    @PostConstruct
    public void init() {
        log.info("Connecting to Bonus Service at {}:{}", bonusServiceHost, bonusServicePort);
        channel = ManagedChannelBuilder
                .forAddress(bonusServiceHost, bonusServicePort)
                .usePlaintext()
                .build();
        blockingStub = BonusServiceGrpc.newBlockingStub(channel);
    }

    @PreDestroy
    public void shutdown() {
        if (channel != null && !channel.isShutdown()) {
            try {
                channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
            } catch (InterruptedException e) {
                channel.shutdownNow();
                Thread.currentThread().interrupt();
            }
        }
    }

    public boolean checkWageringRequirements(UUID userId, double amount) {
        log.debug("Checking wagering requirements for user {} amount {} via gRPC", userId, amount);

        try {
            CheckWageringRequirementsRequest request = CheckWageringRequirementsRequest.newBuilder()
                    .setUserId(userId.toString())
                    .setAmount(amount)
                    .build();

            CheckWageringRequirementsResponse response = blockingStub
                    .withDeadlineAfter(timeoutSeconds, TimeUnit.SECONDS)
                    .checkWageringRequirements(request);

            log.debug("Wagering requirements for user {}: met={}, wagered={}, required={}",
                    userId, response.getRequirementsMet(), response.getWageredAmount(), response.getRequiredAmount());
            return response.getRequirementsMet();
        } catch (StatusRuntimeException e) {
            log.error("gRPC error checking wagering requirements for user {}: {}", userId, e.getStatus());
            throw new BonusServiceException("Bonus service unavailable: " + e.getStatus().getDescription(), e);
        }
    }

    public static class BonusServiceException extends RuntimeException {
        public BonusServiceException(String message, Throwable cause) {
            super(message, cause);
        }
    }
}
