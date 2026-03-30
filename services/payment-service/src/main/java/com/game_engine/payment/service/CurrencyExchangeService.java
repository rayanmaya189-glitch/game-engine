package com.game_engine.payment.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

@Service
@Slf4j
public class CurrencyExchangeService {

    private final WebClient webClient;

    @Value("${payment.exchange.api-url:}")
    private String exchangeApiUrl;

    @Value("${payment.exchange.api-key:}")
    private String exchangeApiKey;

    @Value("${payment.exchange.base-currency:USD}")
    private String baseCurrency;

    @Value("${payment.exchange.cache-ttl-ms:60000}")
    private long cacheTtlMs;

    private final Map<String, CachedRate> rateCache = new ConcurrentHashMap<>();

    public CurrencyExchangeService(WebClient.Builder webClientBuilder) {
        this.webClient = webClientBuilder.build();
    }

    public BigDecimal convert(BigDecimal amount, String fromCurrency, String toCurrency) {
        if (fromCurrency.equalsIgnoreCase(toCurrency)) {
            return amount;
        }

        BigDecimal fromRate = getRate(fromCurrency);
        BigDecimal toRate = getRate(toCurrency);

        return amount.multiply(toRate).divide(fromRate, 4, RoundingMode.HALF_UP);
    }

    public BigDecimal getRate(String currency) {
        if (currency.equalsIgnoreCase(baseCurrency)) {
            return BigDecimal.ONE;
        }

        CachedRate cached = rateCache.get(currency.toUpperCase());
        if (cached != null && !cached.isExpired()) {
            return cached.rate;
        }

        BigDecimal rate = fetchRateFromApi(currency);
        rateCache.put(currency.toUpperCase(), new CachedRate(rate, System.currentTimeMillis()));
        return rate;
    }

    private BigDecimal fetchRateFromApi(String currency) {
        if (exchangeApiUrl == null || exchangeApiUrl.isEmpty()) {
            log.warn("Exchange API not configured, using fallback rates for {}", currency);
            return getFallbackRate(currency);
        }

        try {
            Map<String, Object> response = webClient.get()
                    .uri(exchangeApiUrl + "/latest?base={base}&symbols={symbol}", baseCurrency, currency)
                    .header("Authorization", "Bearer " + exchangeApiKey)
                    .retrieve()
                    .bodyToMono(Map.class)
                    .block();

            if (response != null && response.containsKey("rates")) {
                Map<String, Number> rates = (Map<String, Number>) response.get("rates");
                Number rate = rates.get(currency.toUpperCase());
                if (rate != null) {
                    return BigDecimal.valueOf(rate.doubleValue());
                }
            }
        } catch (Exception e) {
            log.error("Failed to fetch exchange rate for {}: {}", currency, e.getMessage());
        }

        return getFallbackRate(currency);
    }

    private BigDecimal getFallbackRate(String currency) {
        return switch (currency.toUpperCase()) {
            case "USD" -> BigDecimal.ONE;
            case "EUR" -> BigDecimal.valueOf(0.92);
            case "GBP" -> BigDecimal.valueOf(0.79);
            case "JPY" -> BigDecimal.valueOf(149.50);
            case "CNY" -> BigDecimal.valueOf(7.24);
            case "KRW" -> BigDecimal.valueOf(1320.00);
            case "THB" -> BigDecimal.valueOf(35.50);
            case "VND" -> BigDecimal.valueOf(24500.00);
            case "IDR" -> BigDecimal.valueOf(15600.00);
            case "MYR" -> BigDecimal.valueOf(4.65);
            case "SGD" -> BigDecimal.valueOf(1.34);
            case "AUD" -> BigDecimal.valueOf(1.53);
            case "CAD" -> BigDecimal.valueOf(1.36);
            case "CHF" -> BigDecimal.valueOf(0.88);
            case "INR" -> BigDecimal.valueOf(83.10);
            case "BRL" -> BigDecimal.valueOf(4.97);
            case "MXN" -> BigDecimal.valueOf(17.15);
            case "ZAR" -> BigDecimal.valueOf(18.90);
            case "RUB" -> BigDecimal.valueOf(91.50);
            case "AED" -> BigDecimal.valueOf(3.67);
            case "BTC" -> BigDecimal.valueOf(0.000015);
            case "ETH" -> BigDecimal.valueOf(0.00035);
            case "USDT" -> BigDecimal.ONE;
            default -> {
                log.warn("Unknown currency {}, defaulting to 1:1 rate", currency);
                yield BigDecimal.ONE;
            }
        };
    }

    private static class CachedRate {
        final BigDecimal rate;
        final long timestamp;

        CachedRate(BigDecimal rate, long timestamp) {
            this.rate = rate;
            this.timestamp = timestamp;
        }

        boolean isExpired() {
            return System.currentTimeMillis() - timestamp > 60000;
        }
    }
}
