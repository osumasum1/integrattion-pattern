# Wall Street

A wall street firm handles trading of many different commodities in national and international levels, all the transactions are managed internally providing APIs for external customer to fetch reports and data about their stocks, transactions and trade forecasting. They are running on an average of 2000 transactions per second on weekdays and 100 transactions per second on weekends.

Two months ago the firm has been merged with two other firms, this means that their customer base has doubled overnight. Because of it they have had latency issues and outages, increasing the exchange rate from 200 concurrent requests per second to 1000 concurrent requests in their main servers. The increased load is causing delays in the internal transactions and overall disatisfaction in all traders.

What they can do to improve/scale-up their service quality without affecting transactions and external customers?

## Problems detected
* Amount of requests
* Time to get a response

## Solution Proposed
**Messaging - pattern reply-correlation** 
Each client requests is stored in the queue before forwarding the message to the server. When responses are received from the server, correlation information is used to retrieve and restore the correct reply address of the original requester, and send the response to the correct client.
![Pattern](https://github.com/osumasum1/integrattion-pattern/blob/master/images/pattern.png)

**C4 Model** 
![Level 1: Context Diagram](https://github.com/osumasum1/integrattion-pattern/blob/master/images/Level%201.png)
![Level 2: Container Diagram](https://github.com/osumasum1/integrattion-pattern/blob/master/images/Level%202.png)
![Level 3: Component Diagram](https://github.com/osumasum1/integrattion-pattern/blob/master/images/Level%203.png)

## Prerequisites
* [Golang](https://golang.org/dl/)
* [Dep](https://github.com/golang/dep)
