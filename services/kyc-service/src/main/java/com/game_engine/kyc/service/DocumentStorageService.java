package com.game_engine.kyc.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import software.amazon.awssdk.core.sync.RequestBody;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.*;
import software.amazon.awssdk.services.s3.presigner.S3Presigner;
import software.amazon.awssdk.services.s3.presigner.model.GetObjectPresignRequest;

import java.time.Duration;
import java.util.UUID;

@Service
@Slf4j
public class DocumentStorageService {

    private final S3Client s3Client;
    private final S3Presigner presigner;

    private final String bucketName;
    private final Region region;

    public DocumentStorageService(
            @Value("${kyc.document.s3.bucket:}") String bucketName,
            @Value("${kyc.document.s3.region:us-east-1}") String regionName) {

        this.bucketName = bucketName != null ? bucketName : "";
        this.region = Region.of(regionName);

        if (!this.bucketName.isEmpty()) {

            this.s3Client = S3Client.builder()
                    .region(this.region)
                    .build();

            this.presigner = S3Presigner.builder()
                    .region(this.region)
                    .build();

        } else {

            // Development mode
            this.s3Client = null;
            this.presigner = null;
        }
    }

    /**
     * Store document in S3
     */
    public String storeDocument(UUID userId,
            String documentType,
            byte[] data) {

        String key = String.format(
                "users/%s/%s/%s-%d.pdf",
                userId,
                documentType,
                documentType,
                System.currentTimeMillis());

        if (s3Client == null) {
            log.debug("Development mode - mock document stored: {}", key);
            return "s3://mock-bucket/" + key;
        }

        try {

            PutObjectRequest request = PutObjectRequest.builder()
                    .bucket(bucketName)
                    .key(key)
                    .contentType("application/pdf")
                    .serverSideEncryption(ServerSideEncryption.AES256)
                    .build();

            s3Client.putObject(
                    request,
                    RequestBody.fromBytes(data));

            log.info("Document stored successfully: {}", key);

            return "s3://" + bucketName + "/" + key;

        } catch (S3Exception e) {

            log.error(
                    "Failed to store document: {}",
                    e.awsErrorDetails().errorMessage());

            throw new RuntimeException(
                    "Failed to store document",
                    e);
        }
    }

    /**
     * Generate secure pre-signed view URL
     */
    public String generateViewUrl(String s3Key,
            Duration expiration) {

        if (presigner == null) {

            String devBaseUrl = System.getenv()
                    .getOrDefault("KYC_DOCUMENT_BASE_URL", "");

            if (devBaseUrl.isEmpty()) {
                throw new RuntimeException(
                        "S3 not configured and KYC_DOCUMENT_BASE_URL not set");
            }

            return devBaseUrl + "/" + s3Key;
        }

        try {

            GetObjectRequest getObjectRequest = GetObjectRequest.builder()
                    .bucket(bucketName)
                    .key(s3Key)
                    .build();

            GetObjectPresignRequest presignRequest = GetObjectPresignRequest.builder()
                    .signatureDuration(expiration)
                    .getObjectRequest(getObjectRequest)
                    .build();

            String url = presigner
                    .presignGetObject(presignRequest)
                    .url()
                    .toString();

            log.debug("Generated pre-signed URL for {}", s3Key);

            return url;

        } catch (S3Exception e) {

            log.error(
                    "Failed to generate view URL: {}",
                    e.awsErrorDetails().errorMessage());

            throw new RuntimeException(
                    "Failed to generate view URL",
                    e);
        }
    }

    /**
     * Delete document from S3
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

            log.error(
                    "Failed to delete document: {}",
                    e.awsErrorDetails().errorMessage());
        }
    }

    /**
     * Extract key from S3 URL
     */
    public String extractKey(String s3Url) {

        if (s3Url == null || s3Url.isBlank()) {
            return null;
        }

        if (s3Url.startsWith("s3://")) {

            String withoutProtocol = s3Url.substring(5);

            int firstSlash = withoutProtocol.indexOf('/');

            if (firstSlash != -1) {
                return withoutProtocol.substring(firstSlash + 1);
            }
        }

        return s3Url;
    }
}