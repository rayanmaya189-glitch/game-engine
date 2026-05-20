package com.game_engine.payment.grpc;

import com.game_engine.risk.v1.GetRiskScoreRequest;
import com.game_engine.risk.v1.GetRiskScoreResponse;
import com.game_engine.risk.v1.RiskServiceGrpc;
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
public class RiskGrpcClient {

    @Value("${risk.service.host:risk-service}")
    private String riskServiceHost;

    @Value("${risk.service.grpc.port:9016}")
    private int riskServicePort;

    @Value("${risk.service.timeout-seconds:10}")
    private long timeoutSeconds;

    private ManagedChannel channel;
    private RiskServiceGrpc.RiskServiceBlockingStub blockingStub;

    @PostConstruct
    public void init() {
        log.info("Connecting to Risk Service at {}:{}", riskServiceHost, riskServicePort);
        channel = ManagedChannelBuilder
                .forAddress(riskServiceHost, riskServicePort)
                .usePlaintext()
                .build();
        blockingStub = RiskServiceGrpc.newBlockingStub(channel);
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

    public int getRiskScore(UUID userId) {
        log.debug("Getting risk score for user {} via gRPC", userId);

        try {
            GetRiskScoreRequest request = GetRiskScoreRequest.newBuilder()
                    .setUserId(userId.toString())
                    .build();

            GetRiskScoreResponse response = blockingStub
                    .withDeadlineAfter(timeoutSeconds, TimeUnit.SECONDS)
                    .getRiskScore(request);

            log.debug("Risk score for user {}: score={}, level={}", userId, response.getScore(),
                    response.getRiskLevel());
            return response.getScore();
        } catch (StatusRuntimeException e) {
            log.error("gRPC error getting risk score for user {}: {}", userId, e.getStatus());
            throw new RiskServiceException("Risk service unavailable: " + e.getStatus().getDescription(), e);
        }
    }

    public static class RiskServiceException extends RuntimeException {
        public RiskServiceException(String message, Throwable cause) {
            super(message, cause);
        }
    }
}
