# Chamaconekt Visa API

## Overview

The innovation race is on in the financial services industry globally and it has manifested itself through the proliferation of Open API 
portals.Networks, processors and banks are launching developer portals and inviting all developers to explore their Open APIs to build 
their own rendition of the next killer application.

As Chamaconekt, we are leveraging on such open APIs with payment systems like [Visa Inc.]() to help us build the next generation of products 
and services for the financial services industry because of the tremendous opportunity it presents us as a community.The expectations of 
a modern consumer have changed dramatically over the last few years, and consumers are asking for services that are more personalized, 
faster, cheaper, and more easily accessible. The queues at the retail bank branches are getting smaller, and the younger consumers 
wouldn’t even comprehend a bank that doesn’t have an app. As such, it is no longer enough to offer an app with basic account-based 
services (look up my balance, view my transactions, deposit my checks, etc.), the app competition has moved on to more 
innovative features.

Therefore , this repository integrates Chamaconekt with Visa’s payment system via APIs and brings with it a versatile and secure mobile 
payment solution powered by Visa. Chamaconekt can now have an end-to-end platform that can steadily realize its dream for financial 
inclusion as we embrace innovation, at its best, to deliver superior services in the best way possible to our communities. Kenyans already 
understand the benefits of mobile payments, and this integration offers them a better way to pay and be paid, with a service which is not 
limited by the mobile network they have or the handset they use.

Visa is the world's largest global payments company. In 2015, Visa processed $4.9 trillion worth of payment volume on Visa branded cards 
globally.Visa has 2.4 billion cards in circulation and connects 36 million merchants locations with 14,000 plus financial institutions 
done approximately 150 million times a day on the globe. Last year, Visa launched the [Visa developer platform]().This was a transformational 
development for Visa. It is a significant change in the way Visa will design, build and deliver products going forward. It's the first 
time in Visa's 50 year history that they have made their network, data , risk services , value added services available to partners , 
clients and others as open APIs.

Visa has built one of the world’s most advanced processing networks. It's capable of handling more than 24,000 transactions per second, 
with reliability, convenience and security, including fraud protection for consumers and guaranteed payment for merchants.Integrating 
with Visa brings the advantages of Visa’s global network - security, reliability and global acceptance - and allows consumers 
to make payments both domestically and internationally.

Visa is giving more people in more places access to electronic payments. From the world’s major cities to remote areas without banks,
people are increasingly relying on digital currency along with mobile technology to use their money any time, make purchases online , 
transfer funds across borders and access basic financial services. All of which makes their lives easier and grows economies.


## Latest trends in the Global Financial Industry 

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

- Generating the bootstrap code using goagen  ```goagen bootstrap -d ChamaconektVisa/design```

- Building the code  ```go build```

- Running the code  ```./ChamaconektVisa```

- Accessing the API via the command line tool.

    - Change into the directory with the CLI

    ```cd ChamaconektVisa/tool/mvisa-cli ```

    - Build the code

    ``` go build ``` 

    - Call the code

    ``` ./mvisa-cli ```

- Available commands on the CLI
 create
 show


- The content of the app package:

    This is where the bindings happen between the Go http server and the code.

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



## Resources

[Visa Developer Summit at MWC 2016 ](https://www.youtube.com/watch?v=G2cTfiERHpE&t=2211s)

[Building Microservice Architectures with Go](https://www.youtube.com/watch?v=dVnMLtdJzn4&t=1186s)

[Principles Of Microservices by Sam Newman](https://www.youtube.com/watch?v=PFQnNFe27kU)

[GOTO 2014 • Microservices • Martin Fowler](https://www.youtube.com/watch?v=wgdBVIX9ifA)



 








