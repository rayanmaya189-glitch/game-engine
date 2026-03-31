import os

AML_GRPC_HOST = os.environ.get("AML_GRPC_HOST", "localhost")
AML_GRPC_PORT = int(os.environ.get("AML_GRPC_PORT", "9114"))
FRAUD_GRPC_HOST = os.environ.get("FRAUD_GRPC_HOST", "localhost")
FRAUD_GRPC_PORT = int(os.environ.get("FRAUD_GRPC_PORT", "9115"))
GRPC_PORT = int(os.environ.get("RISK_GRPC_PORT", "9116"))
PORT = int(os.environ.get("RISK_PORT", "9016"))

# Risk scoring thresholds
DEPOSIT_LIMITS = {
    'low': int(os.environ.get("RISK_DEPOSIT_LIMIT_LOW", "50000")),
    'medium': int(os.environ.get("RISK_DEPOSIT_LIMIT_MEDIUM", "10000")),
    'high': int(os.environ.get("RISK_DEPOSIT_LIMIT_HIGH", "1000")),
    'critical': int(os.environ.get("RISK_DEPOSIT_LIMIT_CRITICAL", "0")),
}
WITHDRAWAL_LIMITS = {
    'low': int(os.environ.get("RISK_WITHDRAWAL_LIMIT_LOW", "50000")),
    'medium': int(os.environ.get("RISK_WITHDRAWAL_LIMIT_MEDIUM", "5000")),
    'high': int(os.environ.get("RISK_WITHDRAWAL_LIMIT_HIGH", "500")),
    'critical': int(os.environ.get("RISK_WITHDRAWAL_LIMIT_CRITICAL", "0")),
}
