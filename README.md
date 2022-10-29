# V Pay Onboard Program

This is a Go Application in Acquiring Core Team 's Onboard Program.

Has features such as:

+ Create order by Merchant
+ Show order detail which user can confirm payment
+ Admin Panel can show detail order information

[Detail](https://confluence.zalopay.vn/display/ZTM/%5BPMT-AC%5D+Onboarding+Program)

## Installation

Using terminal and run below command
```bash
# 1. Set up environment
make up-env

# 2. Run order service
make run-order

# 3. Run pay service
make run-pay

# 4. Run admin panel
make run-admin
```

## Usage

```python
# API create-order for Merchant

curl --location --request POST 'localhost:8099/api/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "MerchantID": "Beo oi 2",
    "AppID": 2007,
    "Amount": 800,
    "ProductCode": "QR",
    "Description": "Thanh toán Beo oi",
    "Mac" : "65d39bf75a5814c4eca7963db445e6b3930fdd949ab1ad69c50b1e41854dca0a"
}'

# Run web admin for view order created with link
http://localhost:3001

```