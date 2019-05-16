# cs686_BlockChain_P2_starter

## What is the application?
A Book review application. Allows users to post new reviews and ratings of previously read books.

## The Problem
Research reveals that 93 percent of shoppers say the testimonials they read affect their purchasing decisions. 

Not all review sites, nor the reviews left on them are genuine. The trust placed in these platforms means there’s a lot of manipulation going on review sites. According to the same survey, four in five of us have read a fake review in the past year, and 84 percent of us have admitted we struggle to decipher which posts are authentic.

One of the reasons why writing reviews are so easy is that review writers don’t have to prove that they have purchased a product – and, in some cases, they don’t even need to provide their name or use an account to submit their review. 

Research done by Podium. 
https://www.google.com/search?q=podium&oq=podium+&aqs=chrome..69i57j69i60j0l4.1034j1j7&sourceid=chrome&ie=UTF-8 

## Why do we need blockchain technology for this application?
Blockchain technology allows us to satisfy both authenticity and integrity of book reviews: 

- Authenticity: Reviews posted can be verified that it came from the corresponding reviewer. 
- Integrity: Reviews posted onto the site cannot be tampered with by a third party. 

We can achieve these two features using public-private key cryptology on the blockchain platform, coupled with blockchain technology we can ensure immutability via our proof of work mechanism.

## Overview 
The application separates miner nodes from client nodes and each have overlapping and distinct functionalities of their own. 
#### Client Nodes
Client Nodes have user interfaces which allow clients to post new book reviews, check the book reviews on the platform, and check all the books that are in the client library. Client nodes receive input from users and sends the reviews onto Miners to be added into the blockchain. 
#### Miner Nodes
Miner Nodes do not have interfaces, however, we are able to check the number of transactions that a miner has processed and the total transaction fees accumulated over each block produced. Miners receive new reviews (transactions) from Client Nodes and stores them into a local transaction pool. When transaction pools are not empty, they add new transactions into their block and attempt to solve the block with ProofOfWork. 



## Workflow 
#### Lifecycle of a new book review 
1. New users use their address as their public key, and usernames as their private key.
2. Client Node Workflow:
 - User submits a new book review 
 - Client Node adds a bookId to the new book review by checking a local book database - this is shared with all other Client nodes via HeartBeats.
 - Client Node creates a stamp using the private key (username) of the user, and adds it into the new ReviewData.
 - Client Node then hashes the new ReviewData - Hash(ReviewData) to obtain a Signature.
 - Client sends a) Signature b) ReviewDataJson, to all Miners in its Peerlist.
4. Miners workflow:
 - Miner Node receive the new ReviewDataJson and Signature.
 - Miner Node verifies that Hash(ReviewData) == Signature, since the ReviewData was created with the Sender's private key
 - Miner Node adds the ReviewData into their transaction pool queue. 
 - Concurrently, if the transaction pool is not empty, Miners poll new ReviewData from the transaction pool to be added into their block.
5. Miners attempt to solve the ProofOfWork required to earn the transaction fees. 
6. If they solve ProofOfWork, they broadcast the solved block to the peerlist of peer nodes. Else, they chase the next block.
7. Miner Node adds the transaction fees into its wallet. 

## What are the list of functionalities?
### Client Node
#### Allows a client to retrieve all book reviews that exist on the Book Review Application
```
GET /getAllBookReviews
Response: text/html
```

#### Gets the book review page
```
GET /newBookReview
Response: text/html
```

#### Gets all the books in the local book database
```
GET /getAllBooks
Response: text/html
```

#### Allows a client to post a new book review onto the Book Review Application
```
POST /newBookReview
Content-type: application/json

Struct bookReviewWithId 
{"title": string,
"reviewText": string,
"reviewRating": int,
"txFee": float,
"pubKey": string,
"priKey": string,
}
```

- MPT: key = bookTitle; value = bookReview JSON
 
- title: User generated content
- reviewText: User generated content
- reviewRating: User generated content <Value from 0-5>
- transactionFee: User generated content
- pubKey: Sender Address
- priKey: Sender Username

### Miner Node
#### Shows the all transaction fees obtained and total accumulated fees
```
GET /showWallet
Response: text/html
```

#### Miner Node receives the incoming ReviewData from sent from the Client Nodes
```
POST /reviewObjectReceive
Response: text/html
```

## Define the success of this product?
 - [x] Users will be able to post new reviews to the application 
 - [x] Users will be able to view all reviewed books on the platform

## Midpoint Task List
 - [x] Set up new data structures 
 - [x] POST newBookReview endpoint
 - [x] GET allBookReviews endpoint
 - [x] GET createBookReview endpoint

## Final Deadline Task List & Milestones
 - [x] NewBookReview FrontEnd 
 - [x] AllBookReviews FrontEnd
 - [x] Workflow logic for POST /newBookReview
 - [x] Workflow logic for GET /allBookReviews
 - [x] Workflow logic for GET /newBookReview
 - [x] Workflow for Miner transactions 
 - [x] Workflow for Book Database
 - [x] Local Transaction pool 
 
## Future work and improvements 
- [ ] Implement RSA public & private key pair for each Node, ensures Integrity of data


## Disclaimer & Assumptions
This project is an academic level abstract of an actual decentralized application on the public blockchain. 
The implementation of the underlying blockchain is simplified to fit the nature of this project, and several assumptions have been made about transactions to allow the project to work seamlessly. 

1. Client that posts a new book review owns sufficient transaction fees and/or application currency
2. New book review is assumed to be a valid book
3. Client Nodes have all been signed up to the platform and have obtained a unique private username
