package com.game_engine.payment.grpc;

import com.game_engine.payment.model.Payment;
import com.game_engine.payment.service.PaymentService;
import com.game_engine.payment.v1.*;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;

import java.math.BigDecimal;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.List;
import java.util.UUID;

@GrpcService
@RequiredArgsConstructor
@Slf4j
public class GrpcPaymentService extends PaymentServiceGrpc.PaymentServiceImplBase {

    private final PaymentService paymentService;

    @Override
    public void createDeposit(CreateDepositRequest request, StreamObserver<CreateDepositResponse> responseObserver) {
        try {
            Payment.PaymentMethod method = Payment.PaymentMethod.valueOf(request.getMethod());
            Payment payment = paymentService.createDeposit(
                    request.getUserId(), BigDecimal.valueOf(request.getAmount()),
                    method, request.getCurrency(), request.getDescription());

            responseObserver.onNext(CreateDepositResponse.newBuilder()
                    .setPayment(toProtoPayment(payment)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error creating deposit", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void processDeposit(ProcessDepositRequest request, StreamObserver<ProcessDepositResponse> responseObserver) {
        try {
            UUID paymentId = UUID.fromString(request.getPaymentId());
            Payment payment = paymentService.processDeposit(paymentId);

            responseObserver.onNext(ProcessDepositResponse.newBuilder()
                    .setPayment(toProtoPayment(payment)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error processing deposit", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void createWithdrawal(CreateWithdrawalRequest request, StreamObserver<CreateWithdrawalResponse> responseObserver) {
        try {
            Payment.PaymentMethod method = Payment.PaymentMethod.valueOf(request.getMethod());
            Payment payment = paymentService.createWithdrawal(
                    request.getUserId(), BigDecimal.valueOf(request.getAmount()),
                    method, request.getCurrency(), request.getAccountDetails(), request.getDescription());

            responseObserver.onNext(CreateWithdrawalResponse.newBuilder()
                    .setPayment(toProtoPayment(payment)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error creating withdrawal", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void processWithdrawal(ProcessWithdrawalRequest request, StreamObserver<ProcessWithdrawalResponse> responseObserver) {
        try {
            UUID paymentId = UUID.fromString(request.getPaymentId());
            Payment payment = paymentService.processWithdrawal(paymentId);

            responseObserver.onNext(ProcessWithdrawalResponse.newBuilder()
                    .setPayment(toProtoPayment(payment)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error processing withdrawal", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void refundPayment(RefundPaymentRequest request, StreamObserver<RefundPaymentResponse> responseObserver) {
        try {
            UUID paymentId = UUID.fromString(request.getPaymentId());
            Payment payment = paymentService.refundPayment(paymentId, BigDecimal.valueOf(request.getAmount()));

            responseObserver.onNext(RefundPaymentResponse.newBuilder()
                    .setPayment(toProtoPayment(payment)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error refunding payment", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getPayment(GetPaymentRequest request, StreamObserver<GetPaymentResponse> responseObserver) {
        try {
            UUID paymentId = UUID.fromString(request.getPaymentId());
            var payment = paymentService.getPayment(paymentId);

            GetPaymentResponse.Builder response = GetPaymentResponse.newBuilder().setFound(payment.isPresent());
            payment.ifPresent(p -> response.setPayment(toProtoPayment(p)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting payment", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getPaymentByExternalId(GetPaymentByExternalIdRequest request, StreamObserver<GetPaymentByExternalIdResponse> responseObserver) {
        try {
            var payment = paymentService.getPaymentByExternalId(request.getExternalId());

            GetPaymentByExternalIdResponse.Builder response = GetPaymentByExternalIdResponse.newBuilder().setFound(payment.isPresent());
            payment.ifPresent(p -> response.setPayment(toProtoPayment(p)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting payment by external ID", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getUserPayments(GetUserPaymentsRequest request, StreamObserver<GetUserPaymentsResponse> responseObserver) {
        try {
            Sort sort = request.getSortDir().equalsIgnoreCase("asc")
                    ? Sort.by(request.getSortBy()).ascending()
                    : Sort.by(request.getSortBy()).descending();

            PageRequest pageRequest = PageRequest.of(request.getPage(), request.getSize(), sort);
            Page<Payment> payments = paymentService.getUserPayments(request.getUserId(), pageRequest);

            GetUserPaymentsResponse.Builder response = GetUserPaymentsResponse.newBuilder()
                    .setTotalElements((int) payments.getTotalElements())
                    .setTotalPages(payments.getTotalPages())
                    .setCurrentPage(payments.getNumber());
            payments.getContent().forEach(p -> response.addPayments(toProtoPayment(p)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting user payments", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getUserPaymentSummary(GetUserPaymentSummaryRequest request, StreamObserver<GetUserPaymentSummaryResponse> responseObserver) {
        try {
            Instant startDate = Instant.now().minus(request.getDays(), ChronoUnit.DAYS);
            Instant endDate = Instant.now();

            BigDecimal totalDeposits = paymentService.getUserTotalDeposits(request.getUserId(), startDate, endDate);
            BigDecimal totalWithdrawals = paymentService.getUserTotalWithdrawals(request.getUserId(), startDate, endDate);

            GetUserPaymentSummaryResponse response = GetUserPaymentSummaryResponse.newBuilder()
                    .setTotalDeposits(totalDeposits.doubleValue())
                    .setTotalWithdrawals(totalWithdrawals.doubleValue())
                    .setNetAmount(totalDeposits.subtract(totalWithdrawals).doubleValue())
                    .build();

            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting user payment summary", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getSupportedMethods(GetSupportedMethodsRequest request, StreamObserver<GetSupportedMethodsResponse> responseObserver) {
        try {
            List<Payment.PaymentMethod> depositMethods = paymentService.getSupportedMethods(Payment.PaymentType.DEPOSIT);
            List<Payment.PaymentMethod> withdrawalMethods = paymentService.getSupportedMethods(Payment.PaymentType.WITHDRAWAL);

            GetSupportedMethodsResponse response = GetSupportedMethodsResponse.newBuilder()
                    .addAllDepositMethods(depositMethods.stream().map(Enum::name).toList())
                    .addAllWithdrawalMethods(withdrawalMethods.stream().map(Enum::name).toList())
                    .build();

            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting supported methods", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getSupportedCurrencies(GetSupportedCurrenciesRequest request, StreamObserver<GetSupportedCurrenciesResponse> responseObserver) {
        try {
            List<String> currencies = paymentService.getSupportedCurrencies();
            GetSupportedCurrenciesResponse response = GetSupportedCurrenciesResponse.newBuilder()
                    .addAllCurrencies(currencies).build();

            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting supported currencies", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void cancelPayment(CancelPaymentRequest request, StreamObserver<CancelPaymentResponse> responseObserver) {
        try {
            UUID paymentId = UUID.fromString(request.getPaymentId());
            var paymentOpt = paymentService.getPayment(paymentId);

            if (paymentOpt.isEmpty()) {
                responseObserver.onError(io.grpc.Status.NOT_FOUND.withDescription("Payment not found").asRuntimeException());
                return;
            }

            Payment payment = paymentOpt.get();
            if (payment.getStatus() == Payment.PaymentStatus.PENDING) {
                payment.setStatus(Payment.PaymentStatus.CANCELLED);
                responseObserver.onNext(CancelPaymentResponse.newBuilder()
                        .setPayment(toProtoPayment(payment))
                        .setSuccess(true).build());
            } else {
                responseObserver.onNext(CancelPaymentResponse.newBuilder()
                        .setPayment(toProtoPayment(payment))
                        .setSuccess(false).build());
            }
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error cancelling payment", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    private PaymentServiceProto.Payment toProtoPayment(Payment p) {
        PaymentServiceProto.Payment.Builder builder = PaymentServiceProto.Payment.newBuilder()
                .setId(p.getId().toString())
                .setUserId(p.getUserId() != null ? p.getUserId() : "")
                .setExternalId(p.getExternalId() != null ? p.getExternalId() : "")
                .setType(p.getType() != null ? p.getType().name() : "")
                .setMethod(p.getMethod() != null ? p.getMethod().name() : "")
                .setStatus(p.getStatus() != null ? p.getStatus().name() : "");

        if (p.getAmount() != null) builder.setAmount(p.getAmount().doubleValue());
        if (p.getCurrency() != null) builder.setCurrency(p.getCurrency());
        if (p.getDescription() != null) builder.setDescription(p.getDescription());
        if (p.getCreatedAt() != null) builder.setCreatedAt(p.getCreatedAt().toEpochMilli());
        if (p.getProcessedAt() != null) builder.setProcessedAt(p.getProcessedAt().toEpochMilli());
        if (p.getCompletedAt() != null) builder.setCompletedAt(p.getCompletedAt().toEpochMilli());
        if (p.getFailureReason() != null) builder.setFailureReason(p.getFailureReason());

        return builder.build();
    }
}
