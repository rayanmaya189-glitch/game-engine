package com.game_engine.kyc.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import software.amazon.awssdk.core.sync.RequestBody;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.*;

import java.nio.charset.StandardCharsets;
import java.time.Duration;
import java.util.UUID;

/**
 * Document Storage Service
 * 
 * Handles storage of KYC documents in S3 with encryption.
 * Provides pre-signed URLs for secure document access.
 */
@Service
@Slf4j
public class DocumentStorageService {

    private final S3Client s3Client;
    private final String bucketName;
    private final Region region;

    public DocumentStorageService(
            @Value("${kyc.document.s3.bucket:}") String bucketName,
            @Value("${kyc.document.s3.region:us-east-1}") String region) {
        
        if (bucketName != null && !bucketName.isEmpty()) {
            this.s3Client = S3Client.builder()
                    .region(Region.of(region))
                    .build();
        } else {
            // Use mock for development
            this.s3Client = null;
        }
        this.bucketName = bucketName != null ? bucketName : "kyc-documents-dev";
        this.region = Region.of(region);
    }

    /**
     * Store a document in S3 with encryption
     */
    public String storeDocument(UUID userId, String documentType, byte[] data) {
        String key = String.format("users/%s/%s/%s-%s.pdf", 
                userId, documentType, documentType, System.currentTimeMillis());

        if (s3Client == null) {
            // Development mode - return mock URL
            log.debug("Development mode - document stored at: {}", key);
            return "s3://" + bucketName + "/" + key;
        }

        try {
            PutObjectRequest request = PutObjectRequest.builder()
                    .bucket(bucketName)
                    .key(key)
                    .contentType("application/pdf")
                    .serverSideEncryption(ServerSideEncryption.AES256)
                    .build();

            s3Client.putObject(request, RequestBody.fromBytes(data));
            log.info("Document stored: {}", key);

            return "s3://" + bucketName + "/" + key;

        } catch (S3Exception e) {
            log.error("Failed to store document: {}", e.awsErrorDetails().errorMessage());
            throw new RuntimeException("Failed to store document", e);
        }
    }

    /**
     * Generate pre-signed URL for viewing a document
     */
    public String generateViewUrl(String s3Key, Duration expiration) {
        if (s3Client == null) {
            String devBaseUrl = System.getenv().getOrDefault("KYC_DOCUMENT_BASE_URL", "");
            if (devBaseUrl.isEmpty()) {
                throw new RuntimeException("S3 client not configured and KYC_DOCUMENT_BASE_URL not set");
            }
            return devBaseUrl + "/" + s3Key;
        }

        try {
            GetObjectRequest request = GetObjectRequest.builder()
                    .bucket(bucketName)
                    .key(s3Key)
                    .build();

            return s3Client.presignGetObject(request, expiration);

        } catch (S3Exception e) {
            log.error("Failed to generate view URL: {}", e.awsErrorDetails().errorMessage());
            throw new RuntimeException("Failed to generate view URL", e);
        }
    }

    /**
     * Delete a document
     */
    public void deleteDocument(String s3Key) {
        if (s3Client == null) {
            return;
        }

        try {
            DeleteObjectRequest request = DeleteObjectRequest.builder()
                    .bucket(bucketName)
                    .key(s3Key)
                    .build();

            s3Client.deleteObject(request);
            log.info("Document deleted: {}", s3Key);

        } catch (S3Exception e) {
            log.error("Failed to delete document: {}", e.awsErrorDetails().errorMessage());
        }
    }

    /**
     * Extract S3 key from S3 URL
     */
    public String extractKey(String s3Url) {
        if (s3Url == null) {
            return null;
        }
        if (s3Url.startsWith("s3://")) {
            return s3Url.substring(5).split("/", 2)[1];
        }
        return s3Url;
    }
}
