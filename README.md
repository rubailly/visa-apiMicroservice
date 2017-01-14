# Chamaconekt Visa API

This repository contains internal APIs that Chamaconekt Microservices use to deliver value using Visa's rich API sets.

## Overview

The innovation race is on in the financial services industry globally and it has manifested itself through the proliferation of Open API 
portals.Networks, processors and banks are launching developer portals and inviting all developers to explore their Open APIs to build 
their own rendition of the next killer application.

As Chamaconekt, we are leveraging on such open APIs with payment systems like [Visa Inc.](https://developer.visa.com/) to help us build 
the next generation of products and services for the financial services industry because of the tremendous opportunity it presents us as 
a community.The expectations of a modern consumer have changed dramatically over the last few years, and consumers are asking for services
hat are more personalized, faster, cheaper, and more easily accessible. The queues at the retail bank branches are getting smaller, and 
the younger consumers wouldn’t even comprehend a bank that doesn’t have an app. As such, it is no longer enough to offer an app with basic 
account-based services (look up my balance, view my transactions, deposit my checks, etc.), the app competition has moved on to more 
innovative features.

Therefore , this repository integrates Chamaconekt with Visa’s payment system via APIs and brings with it a versatile and secure mobile 
payment solution powered by Visa. Chamaconekt can now have an end-to-end platform that can steadily realize its dream for financial 
inclusion as we embrace innovation, at its best, to deliver superior services in the best way possible to our communities. Kenyans already 
understand the benefits of mobile payments, and this integration offers them a better way to pay and be paid, with a service which is not 
limited by the mobile network they have or the handset they use.

Visa is the world's largest global payments company. In 2015, Visa processed $4.9 trillion worth of payment volume on Visa branded cards 
globally.Visa has 2.4 billion cards in circulation and connects 36 million merchants locations with 14,000 plus financial institutions 
done approximately 150 million times a day on the globe. Last year, Visa launched the [Visa developer platform](https://developer.visa.com/).
This was a transformational development for Visa. It is a significant change in the way Visa will design, build and deliver products going
forward. It's the first time in Visa's 50 year history that they have made their network, data , risk services , value added services 
available to partners , clients and others as open APIs.

Visa has built one of the world’s most advanced processing networks. It's capable of handling more than 24,000 transactions per second, 
with reliability, convenience and security, including fraud protection for consumers and guaranteed payment for merchants.Integrating 
with Visa brings the advantages of Visa’s global network - security, reliability and global acceptance - and allows consumers 
to make payments both domestically and internationally.

Visa is giving more people in more places access to electronic payments. From the world’s major cities to remote areas without banks,
people are increasingly relying on digital currency along with mobile technology to use their money any time, make purchases online , 
transfer funds across borders and access basic financial services. All of which makes their lives easier and grows economies.


## Trends in the Global Financial Industry 

Global investment in fintech ventures tripled to $12.21 billion in 2014, clearly signifying that the digital revolution has arrived 
in the financial services sector. It is still unclear whether this presents more of a challenge or an opportunity for the incumbents 
in the industry. But established financial services players are starting to take bold steps to engage with emerging innovations.

It is clear that the digital revolution in financial services is under way, but the impact on current banking players is not as well 
defined. Digital disruption has the potential to shrink the role and relevance of today’s banks, and simultaneously help them create 
better, faster, cheaper services that make them an even more essential part of everyday life for institutions and individuals. To make 
the impact positive, banks are acknowledging that they need to shake themselves out of institutional complacency and recognize that 
merely navigating waves of regulation and waiting for interest rates to rise won’t protect them from obsolescence. 

The world of innovation is incredibly fast paced, and it is no longer possible to keep up, if we plan to ‘own’ every layer that makes 
up our products and services. We need to rely on services from other companies, our partners, or even competitors. While this may sound 
counterintuitive, there is no other alternative, as missing out on the innovation trends is equivalent to becoming irrelevant in the 
marketplace. 

Openness, Collaboration and Investment are the critical themes that emerge for existing banking players if they are to benefit from 
growth driven by new services and productivity. Embracing these themes and creating the right foundations creates challenges to the 
rate of change and approach to risk that are hard- wired into the way banks currently adapt to innovation. This hands an advantage to 
challengers who only hit regulators’ radar once their new business models have found ways to cherry-pick services and customers. 

The concepts of collaboration – or “co-innovation” – are becoming more important within the financial  services and technology industries.
Traditionally, financial services incumbents have been comfortable partnering with others in their own industry - especially where there 
is an opportunity to share processes or services that are considered “non-core”, and which help all collaborators either reduce their 
costs or create a new market opportunity.

Today, the concern is that established financial services players are not doing enough to keep up to speed with this surge in new 
innovation investment. Legacy technology and the difficulty of deploying new technology fast is a big part of this issue. More worrying 
is the speed at which these banks implement new technology. 


## API design


The API service has been described using the goa design language under the directory called design.It has the following files;

goagen generates  a tool by compiling the command specific code generation algorithm together with the design package.

- The DSL executes which creates in-memory data structures that describe the design.

- The code generation algorithm uses the data structures as input to the final output.

## Implementation

- Generate the bootstrap code using goagen  ```goagen bootstrap -d ChamaconektVisa/design```

- Build the code generated from the design  ```go build```

- Run the code  ```./ChamaconektVisa```

### How to access the APIs via the CLI(command line) tool.

- Change into the directory with the CLI

    ```cd ChamaconektVisa/tool/chamaconektvisa-cli ```

- Build the code

    ``` go build ``` 

- Call the code

    ``` ./chamaconektvisa-cli ```

- Available commands on the CLI
    create
    show

### How to access the Swagger UI for API documentation.

- Add the following code on ```resources.go```

```js
var _ = Resource("swagger", func() {
	Description("The API Swagger specification")

	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})
```

- Clone the swagger-ui from [GitHub](https://github.com/swagger-api/swagger-ui) in your home directory

- Copy the Swagger UI folder that you cloned above into the project repository    

- Rebuild the code

    ``` go build ``` 

- Relaunch the code

    ``` ./chamaconektvisa ```

- The output in the terminal looks like this.

```bash
william@william-Compaq-610:~/chamaconekt/gocode/src/ChamaconektVisa$ ./ChamaconektVisa
2017/01/13 16:18:39 [INFO] mount ctrl=Deposit action=Create route=POST /deposit
2017/01/13 16:18:39 [INFO] mount ctrl=Deposit action=Show route=GET /deposit/:id
2017/01/13 16:18:39 [INFO] mount ctrl=Payment action=Create route=POST /payment
2017/01/13 16:18:39 [INFO] mount ctrl=Payment action=Show route=GET /payment/:id
2017/01/13 16:18:39 [INFO] mount ctrl=Swagger files=swagger-ui/ route=GET /swagger-ui/*filepath
2017/01/13 16:18:39 [INFO] mount ctrl=Swagger files=swagger/swagger.json route=GET /swagger.json
2017/01/13 16:18:39 [INFO] mount ctrl=Swagger files=swagger-ui/index.html route=GET /swagger-ui/
2017/01/13 16:18:39 [INFO] mount ctrl=Withdrawal action=Create route=POST /withdrawal
2017/01/13 16:18:39 [INFO] mount ctrl=Withdrawal action=Show route=GET /withdrawal/:id
2017/01/13 16:18:39 [INFO] listen transport=http addr=:8080
```
Notice there are two new endpoints that our API is exposing; The ```swagger-ui``` and the 
```swagger.json``` 





- The content of the app package:

    In the app packages is where the bindings happen between the Go http server and the code.

    - ```controllers.go ```  Contains the controller interface type definitions. There is one such interface per resource defined in the design language. The file 
    also contains the code that “mounts” implementations of these controller interfaces onto the service. The exact meaning of “mounting” 
    a controller is discussed further below.

    - ```contexts.go```    Contains the context data structure definitions. Contexts play a similar role to Martini’s ```martini.Context```, goji’s ```web.C``` 
    or echo’s ```echo.Context``` to take a few arbitrary examples: they are given as first argument to all controller actions and provide 
    helper methods to access the request state and write the response.

    - ```hrefs.go```    Provide global functions for building resource hrefs. Resource hrefs make it possible for responses to link to related resources. goa 
    knows how to build these hrefs by looking at the request path for the resource “canonical” action (by default the ```show``` action). 

    - ```media_types.go```    Contains the media type data structures used by resource actions to build the responses. There is one such data structure generated per 
    view defined in the design.

    - ```user_types.go```    Contains the data structures defined via the Type design language function. Such types may be used to define request payloads and 
    response media types.

    - ```test/```    contains test helpers that make it convenient to test the controller code by making it possible to call the action implementations with 
    controller input and validate the resulting media types.




## Running locally

Assuming a working Go setup:

```bash
go install github.com/chamaconekt/ChamaconektVisa
```

```bash
goa-ChamaconektVisa
```

Once running `goa-ChamaconektVisa` listens on port 8080. 

## APIs

### Deposit API
THe Deposit API is an internal API that interacts with Visa's [CashInPushPayments POST API]()

### Withdrawal API
The Withdrawal API is an internal API that interacts with Visa's [CashOutPushPayments POST]()

### Payment API 
The Payment API is an internal API that interacts with Visa's [MerchantPushPayments POST]() that 

This API is used for payment to small financial institutions for goods or services purchased, either face-to-face or remote.This is 
leveraged using a mobile phone which clients use to  authenticate themselves and also provide payment instructions to the relevant 
institution. .Each and every institution has a unique ID and can be captured during payment via QR codes , key entry , NFC or other means

Upon receiving the payment instructions, the clients issuer (Bank that provided you your Visa card) sends payment instructions to the
instructions bank account.

The institutions acquirer(a bank or financial institution that processes credit or debit card payments on behalf of a merchant) processes 
the Visa message, creates a record of merchant payment and reverts back with a response message containing the MerchantPushPayments Response Attributes.

An acquiring bank is a bank or financial institution that processes credit or debit card payments on behalf of a merchant.


###  Checkout API 

The Checkout API is an internal API that interacts with Visa's [Get Payment Data API](https://developer.visa.com/products/visa_checkout/reference#visa_checkout__get_payment_data_api) 
that obtains clients payment information associated with a payment request from a Visa Checkout transaction.ChamaconektVisa processes
the data as a convenience.This API retrieves the client's payment information for a particular order. 

### Validate-checkout API
The Validate-checkout API is an internal API that interacts with Visa's [Update Payment Information API]()  to provide other Chamaconekt 
microservices with the status of the transaction and final payment amounts  a client is making in the Visa Checkout .This API confirms, and if needed 
modify, the amounts the client specified in the Visa Checkout for a transaction.

# Contributing 

## How to contribute
We're striving to keep master's history with minimal merge bubbles. To achieve this, we're asking pull requests to be submitted 
on top of master.

## Finding things to work on
The first place to start is always looking over the current github issues for the project you are interested in contributing to.Issues 
marked with __help wanted__ are usually pretty self contained and a good place to get started. Chamaconekt also uses these same github 
issues to keep track of what we are working on.If you see any issues that are assigned to a particular person that means someone is 
currently working on that issue.Of course feel free to make your own issues if you think something needs to be added or fixed.


## Basic quality checks
Please ensure that all tests pass before submitting changes.Try to separate logically distinct changes into separate commits and 
thematically distinct commits into separate pull requests.

## Submitting changes
Please sign the [Contributor License Agreement](http://bit.ly/2gtLB0J). All content, comments, and pull requests must follow the
[Chamaconekt Community Guidelines]()

Submit a pull request on top of master
- Include a descriptive commit message
- Changes contributed via pull requests should focus on a sigle issues at a time.

At this point you're waiting on us.We like to at least comment on pull requests within one week (and typically, three business days).
We may suggest some changes or improvements or alternatives.

# Versioning
We use [SemVer](http://semver.org/) for versioning.

# Authors
- [William Ondenge](https://github.com/wondenge)

# License
This project is licensed under the [Apache 2.0 License](https://www.apache.org/licenses/) 
- see the [LICENSE]() file for details


## Resources

[Visa Developer Summit at MWC 2016 ](https://www.youtube.com/watch?v=G2cTfiERHpE&t=2211s)

[Building Microservice Architectures with Go](https://www.youtube.com/watch?v=dVnMLtdJzn4&t=1186s)

[Principles Of Microservices by Sam Newman](https://www.youtube.com/watch?v=PFQnNFe27kU)

[GOTO 2014 • Microservices • Martin Fowler](https://www.youtube.com/watch?v=wgdBVIX9ifA)



 








