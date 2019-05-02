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

## Workflow 
#### Lifecycle of a new book review 
1. New clients create a new public and private key pair.
2. New clients register with the platform by providing their public keys.
3. Client Workflow:
 - Client checks if the book title already exists and adds a bookId to the new reviewObject
 - Client encrypts the ReviewObject with its private key
 - Client sends a) EncryptedReviewObject b) ReviewObject
 - Client sends a) & b) to each address of its Peerlist
4. Miners workflow:
 - Verify the new ReviewObjects by ensuring the decrypt(EncryptedReviewObject) == ReviewObject
 - Propagate ReviewObject to peerlist
 - Add ReviewObject into their transaction pool queue 
 - Miners pull a ReviewObject from the transaction pool to be added into their block
5. Miners attempt to solve the ProofOfWork required to earn the BlockReward and transaction fees associated with the ReviewObjects. 
6. They broadcast the solved block to the peerlist of peer nodes. 
7. Each miner checks the ReviewObject of the solved block against their existing transaction pool and removes it from their pool
8. Repeat from step 4.

## What are the list of functionalities?

#### Allows a client to retrieve all book reviews that exist on the Book Review Application
```
GET /allBookReviews
Response: text/html
```

#### Allows a client to generate a new public private key pair
```
GET /privatePublicKeyPair

Response: 
{
"public": string,
"private": string,
}
```

#### Takes the client to the page where a new book review can be created 
```
GET /createBookReview
Response: text/html
```

#### Allows a client to post a new book review onto the Book Review Application
##### Implementation details
```
POST /newBookReview
Content-type: application/json

Struct bookReviewWithId 
{"title": string,
"reviewText": string,
"reviewRating": int,
"transactionFee": float,
"publicKey": string,
"signature": string,
"bookId": int,
}
```

- MPT: <key> = bookTitle: <value> = bookReview JSON element 
 
- title: User generated content
- reviewText: User generated content
- reviewRating: User generated content <Value from 0-5>
- transactionFee: User generated content
- publicKey: Sender public key
- signature: privateKey(BookReview)
- bookId: bookId

## Define the success of this product?
 - [ ] Users will be able to post new reviews to the application 
 - [ ] Users will be able to view all reviewed books on the platform

## Midpoint Task List
 - [x] Set up new data structures 
 - [x] POST newBookReview endpoint
 - [x] GET allBookReviews endpoint
 - [x] GET createBookReview endpoint

## Final Deadline Task List & Milestones
 - [ ] NewBookReview FrontEnd 
 - [ ] AllBookReviews FrontEnd
 - [ ] Public Private Key generation for clients
 - [ ] Workflow logic for POST /newBookReview
 - [ ] Workflow logic for GET /allBookReviews
 - [ ] Workflow logic for GET /createBookReview

## Disclaimer & Assumptions
This project is an academic level abstract of an actual decentralized application on the public blockchain. 
The implementation of the underlying blockchain is simplified to fit the nature of this project, and several assumptions have been made about transactions to allow the project to work seamlessly. 

1. Client that posts a new book review owns sufficient transaction fees 
2. New book review is assumed to be a valid book
