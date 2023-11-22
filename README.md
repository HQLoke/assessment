## INSTRUCTIONS

#### Build a microservice application in Golang that:
#### 1. Receive price history of BTCUSDT from Binance (use their test net at https://testnet.binance.vision/)
#### 2. From the price history received, calculate the Average True Range (https://www.investopedia.com/terms/a/atr.asp) of BTC/USDT pair
#### 3. The Average True Range should be 14 days and use the daily interval chart.

## USAGE
### 1. Start localhost server using makefile

```bash
make
```

### 2. Port
#### The server is running on port '3456'.

### 3. API Endpoint Structure
#### The API endpoint structure is as follows:
```
http://localhost:3456/api/indicator/atr?symbol=<symbol>&period=<period>
```
#### Replace `<symbol>` with desired symbol (currently just BTCUSDT) and `<period>` with the ATR period.

### 4. Parameters
#### `symbol` (mandatory)
#### `period` (optional): default value is 14. Cannot exceed the number of data points.
#### Finally, this GET request will calculate the Average True Range (ATR) indicator for the symbol based on your parameters.
