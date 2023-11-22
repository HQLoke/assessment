INSTRUCTIONS:

Build a microservice application in Golang that:
1. Receive price history of BTCUSDT from Binance (use their test net at https://testnet.binance.vision/)
2. From the price history received, calculate the Average True Range (https://www.investopedia.com/terms/a/atr.asp) of BTC/USDT pair
3. The Average True Range should be 14 days and use the daily interval chart.

USAGE:
/api/indicator/atr?symbol=BTCUSDT&period=14